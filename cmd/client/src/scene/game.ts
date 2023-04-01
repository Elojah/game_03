import { GameObjects, Scene } from 'phaser';
import { grpc } from '@improbable-eng/grpc-web';

import { Mutex } from 'async-mutex';

import * as jspb from 'google-protobuf';

import { ulid, parse } from '../lib/ulid'

import { API } from 'cmd/api/grpc/api_pb_service';

import * as Animation from 'pkg/entity/animation_pb';
import * as AnimationDTO from 'pkg/entity/dto/animation_pb';

import * as Entity from 'pkg/entity/entity_pb';
import * as EntityDTO from 'pkg/entity/dto/entity_pb';

import * as PC from 'pkg/entity/pc_pb';
import * as PCDTO from 'pkg/entity/dto/pc_pb';

import * as Cell from 'pkg/room/cell_pb';
import * as CellDTO from 'pkg/room/dto/cell_pb';

import * as World from 'pkg/room/world_pb';
import * as WorldDTO from 'pkg/room/dto/world_pb';

const entitySpriteDepth = 42

type valueof<T> = T[keyof T];
type Orientation = valueof<Cell.OrientationMap>

type BodyEntity = {
	E: Entity.E
	Body: Phaser.Types.Physics.Arcade.SpriteWithDynamicBody
	Animations: Map<string, string>
	Orientation: Orientation
}

type GraphicEntity = {
	E: Entity.E
	Animations: Map<string, string>
	Direction: Phaser.Geom.Point
	Interpolation: number
	Sprite: Phaser.GameObjects.Sprite

	Objects: Map<string, Phaser.Physics.Arcade.StaticGroup>
	Colliders: Map<string, Phaser.Physics.Arcade.Collider>
}

function updateGraphicEntity(ge: GraphicEntity, x: number, y: number) {
	// update direction to last known position
	if (ge.Direction.x != x || ge.Direction.y != y) {
		// TODO update interpolation as speed to catch up latency
		ge.Interpolation = 0
	}

	// TODO Add prediction/reconciliation ?

	ge.Direction.setTo(x, y)
}

function interpolateGraphicEntity(ge: GraphicEntity) {
	if (ge.Interpolation < 1) {
		ge.Interpolation = ge.Interpolation + 0.05
	}
}

type GraphicCell = {
	Cell: Cell.Cell
	Tilemap: Phaser.Tilemaps.Tilemap
	Objects: Map<string, Phaser.Physics.Arcade.StaticGroup>
	Colliders: Map<string, Phaser.Physics.Arcade.Collider>
	Layers: Map<string, Phaser.Tilemaps.TilemapLayer | null>
}

export class Game extends Scene {

	// Self
	PC: PC.PC;
	Entity: BodyEntity;
	EntityID: string;
	Cursors: Phaser.Types.Input.Keyboard.CursorKeys
	// Controls: Controls;

	// Cells & maps
	Cells: Map<Orientation, GraphicCell>
	CellsByID: Map<string, Orientation>
	// Specific loader for cells
	CellLoader: Phaser.Loader.LoaderPlugin
	// Current world loaded
	World: World.World
	// Client cell loading
	Border: Map<Orientation, number>
	// Background image
	Background: Phaser.GameObjects.TileSprite
	// Loading mutex
	Loading: Uint8Array | undefined

	// Entities
	Entities: Map<string, GraphicEntity>
	// Specific loader for entities
	EntityLoader: Phaser.Loader.LoaderPlugin
	// Grpc client to send entity updates to server
	EntityClient: grpc.Client<Entity.E, Entity.E>
	// Entity update map to detect disconnect
	EntityLastTick: Map<string, number>
	cleanEntitiesDelay: number

	SyncTimer: number

	// Spritesheets already loaded
	SpriteSheets: Map<string, integer>

	// Static blank cell for filling
	Blank: GraphicCell

	// Start load pc mutex before map
	LoadMapMutex: Mutex

	constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
		super(config);
	}
	init(pc: PCDTO.PC) {
		this.PC = pc.getPc() as PC.PC;
		const e = pc.getEntity() as Entity.E;
		this.Entity = {
			E: e,
			Body: this.physics.add.sprite(0, 0, ''),
			Animations: new Map(),
			Orientation: Cell.Orientation.DOWN
		}
		this.EntityID = ulid(this.Entity.E.getId_asU8())

		this.Cells = new Map()
		this.CellsByID = new Map()
		this.CellLoader = new Phaser.Loader.LoaderPlugin(this)
		this.Border = new Map()

		this.Entities = new Map()
		this.EntityLoader = new Phaser.Loader.LoaderPlugin(this)
		this.EntityLastTick = new Map()
		this.cleanEntitiesDelay = 2 * 1000 // 2 seconds without tick (async) triggers entity removal

		this.SpriteSheets = new Map()

		this.World = new World.World()

		this.LoadMapMutex = new Mutex()

		const m = this.make.tilemap()
		this.Blank = {
			Cell: new Cell.Cell,
			Tilemap: m,
			Layers: new Map(),
			Colliders: new Map(),
			Objects: new Map(),
		}
	}
	preload() { }
	create() {
		// Connect for entity sync
		this.connect()
			.then(() => {
				console.log('disconnect from server')
			})

		// Connect for entity send
		this.connectUpdate()

		this.Loading = this.Entity.E.getCellid_asU8()

		this.Cursors = this.input.keyboard!.createCursorKeys();

		this.LoadMapMutex.waitForUnlock()
			.then(() => {

				// Load world & cell sync
				this.listWorld((() => {
					const req = new WorldDTO.ListWorldReq()

					req.setIdsList([this.PC.getWorldid_asU8()])
					req.setSize(1)

					return req
				})())
					.then((ws: WorldDTO.ListWorldResp) => {
						if (ws.getWorldsList().length != 1) {
							console.log('failed to load world')
						} else {
							this.World = ws.getWorldsList()[0]
						}
					}).then(() => {
						return this.loadBackground()
					}).then(() => {
						return this.loadMap(Cell.Orientation.NONE)
					})
					.then(() => {
						this.load.start()
						this.load.on('complete', () => {
							console.log('load complete')
						})
					}).then(() => {
						// post first loading launch
						setInterval(() => {
							this.cleanEntities()
						}, this.cleanEntitiesDelay)
					})
					.catch((err) => {
						console.log('setup error', err)
					})
			})
	}

	updateBodyEntity() {
		// Controls + local anim update
		const speed: number = 200
		let animationID = null
		this.Entity.Body.body.setVelocity(0)

		// Move entity
		if (this.Cursors.up.isDown) {
			animationID = this.Entity.Animations.get('walk_up')
			this.Entity.Body.body.setVelocityY(-speed)
			this.Entity.Orientation = Cell.Orientation.UP
		} else if (this.Cursors.down.isDown) {
			animationID = this.Entity.Animations.get('walk_down')
			this.Entity.Body.body.setVelocityY(speed)
			this.Entity.Orientation = Cell.Orientation.DOWN
		}

		if (this.Cursors.right.isDown) {
			animationID = this.Entity.Animations.get('walk_right')
			this.Entity.Body.body.setVelocityX(speed)
			this.Entity.Orientation = Cell.Orientation.RIGHT
		} else if (this.Cursors.left.isDown) {
			animationID = this.Entity.Animations.get('walk_left')
			this.Entity.Body.body.setVelocityX(-speed)
			this.Entity.Orientation = Cell.Orientation.LEFT
		}

		if (!animationID) {
			switch (this.Entity.Orientation) {
				case Cell.Orientation.UP:
					animationID = this.Entity.Animations.get('idle_up')
					break
				case Cell.Orientation.RIGHT:
					animationID = this.Entity.Animations.get('idle_right')
					break
				case Cell.Orientation.DOWN:
					animationID = this.Entity.Animations.get('idle_down')
					break
				case Cell.Orientation.LEFT:
					animationID = this.Entity.Animations.get('idle_left')
					break
			}
		}

		if (animationID) {
			this.Entity.E.setAnimationid(parse(animationID))

			const duplicateID = this.Entity.Animations.get(animationID)
			if (duplicateID) {
				this.Entity?.Body?.play(duplicateID, true)
			}
		}
	}

	// positionServerToClient(ge: GraphicEntity): Phaser.Geom.Point {
	// 	const id = ulid(ge.E.getCellid_asU8())

	// 	this.Cells.get(id)
	// 	// return
	// }

	// positionClientToServer(ge: GraphicEntity): Phaser.Geom.Point {
	// 	const id = ulid(ge.E.getCellid_asU8())

	// 	this.Cells.get(id)
	// 	// return
	// }

	update(time: number, deltaTime: number) {
		this.updateBodyEntity()
		this.updateBackground()

		let o: Orientation = Cell.Orientation.NONE
		const x = this.Entity.Body.body.x
		const y = this.Entity.Body.body.y
		const up = this.Border.get(Cell.Orientation.UP)!
		const right = this.Border.get(Cell.Orientation.RIGHT)!
		const down = this.Border.get(Cell.Orientation.DOWN)!
		const left = this.Border.get(Cell.Orientation.LEFT)!

		if (y < up) {
			if (x < left) {
				o = Cell.Orientation.UPLEFT
			} else if (x < right) {
				o = Cell.Orientation.UP
			} else {
				o = Cell.Orientation.UPRIGHT
			}
		} else if (y < down) {
			if (x < left) {
				o = Cell.Orientation.LEFT
			} else if (x < right) {
				o = Cell.Orientation.NONE
			} else {
				o = Cell.Orientation.RIGHT
			}
		} else {
			if (x < left) {
				o = Cell.Orientation.DOWNLEFT
			} else if (x < right) {
				o = Cell.Orientation.DOWN
			} else {
				o = Cell.Orientation.DOWNRIGHT
			}
		}

		if (o != Cell.Orientation.NONE && this.Loading == undefined) {
			const id = this.Cells.get(o)?.Cell.getId_asU8()
			this.Loading = id


			if (id) {
				this.Entity.E.setCellid(id)

				this.loadMap(o).then(() => {
					console.log('loaded cell')
				})
			} else {
				// don't load cell, out of world
			}
		}

		// Move entity
		this.Entity.E.setX(Math.round(x))
		this.Entity.E.setY(Math.round(y))

		this.Entities.forEach((e: GraphicEntity) => {
			// self reconciliation
			if (ulid(e.E.getId_asU8()) == ulid(this.Entity.E.getId_asU8())) {
				return
			}

			interpolateGraphicEntity(e)

			const x = e.Sprite.x + ((e.Direction.x - e.Sprite.x) * e.Interpolation)
			const y = e.Sprite.y + ((e.Direction.y - e.Sprite.y) * e.Interpolation)
			e?.Sprite.setX(x)
			e?.Sprite.setY(y)

			const animationID = e.Animations.get(ulid(e.E.getAnimationid_asU8()))
			if (animationID) {
				e?.Sprite.play(animationID, true)
			}
		})
	}

	async loadBackground() {
		const bg = 'game_background'
		this.load.image(bg, 'img/' + bg + '.png')
		console.log('add listener on ', 'filecomplete-image-' + bg)
		this.load.on('filecomplete-image-' + bg, () => {
			console.log('image completed ', bg)
			this.Background = this.add.tileSprite(0, 0, 1024, 512, bg, 'img/' + bg + '.png').setOrigin(0).setScrollFactor(0);
			this.Background.setDisplaySize(this.game.scale.width, this.game.scale.height)
		})
	}

	updateBackground() {
		if (this.Background) {
			this.Background.tilePositionX = this.Entity.Body.body.x % 1024
			this.Background.tilePositionY = this.Entity.Body.body.y % 512
		}
	}

	// Orientation o uses preload surrounded cells
	// Cell.Orientation.NONE loads everything
	async loadMap(o: Orientation) {
		// call current pc cell
		return this.listCell((() => {
			const req = new CellDTO.ListCellReq()

			req.setIdsList([this.Entity.E.getCellid_asU8()])

			return req
		})())
			.then((cells: CellDTO.ListCellResp) => {
				// load current pc cell
				if (cells.getCellsList().length != 1) {
					throw new Error('failed to load cell')
				}

				const c = cells.getCellsList()[0]
				const contig = c.getContiguousMap() as jspb.Map<number, Uint8Array>

				let newCellIDs: Uint8Array[] = []
				let deletedCells: GraphicCell[] = []

				switch (o) {
					case Cell.Orientation.NONE:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Cell.Orientation.UP)!,
							contig.get(Cell.Orientation.UPRIGHT)!,
							contig.get(Cell.Orientation.RIGHT)!,
							contig.get(Cell.Orientation.DOWNRIGHT)!,
							contig.get(Cell.Orientation.DOWN)!,
							contig.get(Cell.Orientation.DOWNLEFT)!,
							contig.get(Cell.Orientation.LEFT)!,
							contig.get(Cell.Orientation.UPLEFT)!,
						)

						// assign cells to delete
						// deletedCells.push()
						// create new blank graphic cell from loaded cell
						this.Cells.set(Cell.Orientation.NONE, {
							Cell: c,
							Tilemap: this.Blank.Tilemap,
							Layers: this.Blank.Layers,
							Colliders: this.Blank.Colliders,
							Objects: this.Blank.Objects
						})

						this.Cells.delete(Cell.Orientation.UP)
						this.Cells.delete(Cell.Orientation.UPRIGHT)
						this.Cells.delete(Cell.Orientation.RIGHT)
						this.Cells.delete(Cell.Orientation.DOWNRIGHT)
						this.Cells.delete(Cell.Orientation.DOWN)
						this.Cells.delete(Cell.Orientation.DOWNLEFT)
						this.Cells.delete(Cell.Orientation.LEFT)
						this.Cells.delete(Cell.Orientation.UPLEFT)
						break;
					case Cell.Orientation.UP:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Cell.Orientation.UPLEFT)!,
							contig.get(Cell.Orientation.UP)!,
							contig.get(Cell.Orientation.UPRIGHT)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Cell.Orientation.DOWNLEFT)!,
							this.Cells.get(Cell.Orientation.DOWN)!,
							this.Cells.get(Cell.Orientation.DOWNRIGHT)!,
						)
						// shift preloaded
						this.Cells.set(Cell.Orientation.DOWNLEFT, this.Cells.get(Cell.Orientation.LEFT)!)
						this.Cells.set(Cell.Orientation.DOWN, this.Cells.get(Cell.Orientation.NONE)!)
						this.Cells.set(Cell.Orientation.DOWNRIGHT, this.Cells.get(Cell.Orientation.RIGHT)!)
						this.Cells.set(Cell.Orientation.LEFT, this.Cells.get(Cell.Orientation.UPLEFT)!)
						this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.UP)!)
						this.Cells.set(Cell.Orientation.RIGHT, this.Cells.get(Cell.Orientation.UPRIGHT)!)
						this.Cells.delete(Cell.Orientation.UPLEFT)
						this.Cells.delete(Cell.Orientation.UP)
						this.Cells.delete(Cell.Orientation.UPRIGHT)
						break;
					case Cell.Orientation.UPRIGHT:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Cell.Orientation.UPLEFT)!,
							contig.get(Cell.Orientation.UP)!,
							contig.get(Cell.Orientation.UPRIGHT)!,
							contig.get(Cell.Orientation.RIGHT)!,
							contig.get(Cell.Orientation.DOWNRIGHT)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Cell.Orientation.UPLEFT)!,
							this.Cells.get(Cell.Orientation.LEFT)!,
							this.Cells.get(Cell.Orientation.DOWNLEFT)!,
							this.Cells.get(Cell.Orientation.DOWN)!,
							this.Cells.get(Cell.Orientation.DOWNRIGHT)!,
						)
						// shift preloaded
						this.Cells.set(Cell.Orientation.LEFT, this.Cells.get(Cell.Orientation.UP)!)
						this.Cells.set(Cell.Orientation.DOWNLEFT, this.Cells.get(Cell.Orientation.NONE)!)
						this.Cells.set(Cell.Orientation.DOWN, this.Cells.get(Cell.Orientation.RIGHT)!)
						this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.UPRIGHT)!)
						this.Cells.delete(Cell.Orientation.UPLEFT)
						this.Cells.delete(Cell.Orientation.UP)
						this.Cells.delete(Cell.Orientation.UPRIGHT)
						this.Cells.delete(Cell.Orientation.RIGHT)
						this.Cells.delete(Cell.Orientation.DOWNRIGHT)
						break;
					case Cell.Orientation.RIGHT:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Cell.Orientation.UPRIGHT)!,
							contig.get(Cell.Orientation.RIGHT)!,
							contig.get(Cell.Orientation.DOWNRIGHT)!
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Cell.Orientation.UPLEFT)!,
							this.Cells.get(Cell.Orientation.LEFT)!,
							this.Cells.get(Cell.Orientation.DOWNLEFT)!,
						)
						// shift preloaded
						this.Cells.set(Cell.Orientation.UPLEFT, this.Cells.get(Cell.Orientation.UP)!)
						this.Cells.set(Cell.Orientation.LEFT, this.Cells.get(Cell.Orientation.NONE)!)
						this.Cells.set(Cell.Orientation.DOWNLEFT, this.Cells.get(Cell.Orientation.DOWN)!)
						this.Cells.set(Cell.Orientation.UP, this.Cells.get(Cell.Orientation.UPRIGHT)!)
						this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.RIGHT)!)
						this.Cells.set(Cell.Orientation.DOWN, this.Cells.get(Cell.Orientation.DOWNRIGHT)!)
						this.Cells.delete(Cell.Orientation.UPRIGHT)
						this.Cells.delete(Cell.Orientation.RIGHT)
						this.Cells.delete(Cell.Orientation.DOWNRIGHT)
						break;
					case Cell.Orientation.DOWNRIGHT:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Cell.Orientation.DOWNLEFT)!,
							contig.get(Cell.Orientation.DOWN)!,
							contig.get(Cell.Orientation.DOWNRIGHT)!,
							contig.get(Cell.Orientation.RIGHT)!,
							contig.get(Cell.Orientation.UPRIGHT)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Cell.Orientation.UPLEFT)!,
							this.Cells.get(Cell.Orientation.LEFT)!,
							this.Cells.get(Cell.Orientation.DOWNLEFT)!,
							this.Cells.get(Cell.Orientation.UP)!,
							this.Cells.get(Cell.Orientation.UPRIGHT)!,
						)
						// shift preloaded
						this.Cells.set(Cell.Orientation.LEFT, this.Cells.get(Cell.Orientation.DOWN)!)
						this.Cells.set(Cell.Orientation.UPLEFT, this.Cells.get(Cell.Orientation.NONE)!)
						this.Cells.set(Cell.Orientation.UP, this.Cells.get(Cell.Orientation.RIGHT)!)
						this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.DOWNRIGHT)!)
						this.Cells.delete(Cell.Orientation.DOWNLEFT)
						this.Cells.delete(Cell.Orientation.DOWN)
						this.Cells.delete(Cell.Orientation.DOWNRIGHT)
						this.Cells.delete(Cell.Orientation.RIGHT)
						this.Cells.delete(Cell.Orientation.UPRIGHT)
						break;
					case Cell.Orientation.DOWN:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Cell.Orientation.DOWNLEFT)!,
							contig.get(Cell.Orientation.DOWN)!,
							contig.get(Cell.Orientation.DOWNRIGHT)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Cell.Orientation.UPLEFT)!,
							this.Cells.get(Cell.Orientation.UP)!,
							this.Cells.get(Cell.Orientation.UPRIGHT)!,
						)
						// shift preloaded
						this.Cells.set(Cell.Orientation.UPLEFT, this.Cells.get(Cell.Orientation.LEFT)!)
						this.Cells.set(Cell.Orientation.UP, this.Cells.get(Cell.Orientation.NONE)!)
						this.Cells.set(Cell.Orientation.UPRIGHT, this.Cells.get(Cell.Orientation.RIGHT)!)
						this.Cells.set(Cell.Orientation.LEFT, this.Cells.get(Cell.Orientation.DOWNLEFT)!)
						this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.DOWN)!)
						this.Cells.set(Cell.Orientation.RIGHT, this.Cells.get(Cell.Orientation.DOWNRIGHT)!)
						this.Cells.delete(Cell.Orientation.DOWNLEFT)
						this.Cells.delete(Cell.Orientation.DOWN)
						this.Cells.delete(Cell.Orientation.DOWNRIGHT)
						break;
					case Cell.Orientation.DOWNLEFT:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Cell.Orientation.DOWNRIGHT)!,
							contig.get(Cell.Orientation.DOWN)!,
							contig.get(Cell.Orientation.DOWNLEFT)!,
							contig.get(Cell.Orientation.LEFT)!,
							contig.get(Cell.Orientation.UPLEFT)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Cell.Orientation.UPLEFT)!,
							this.Cells.get(Cell.Orientation.UP)!,
							this.Cells.get(Cell.Orientation.UPRIGHT)!,
							this.Cells.get(Cell.Orientation.RIGHT)!,
							this.Cells.get(Cell.Orientation.DOWNRIGHT)!,
						)
						// shift preloaded
						this.Cells.set(Cell.Orientation.RIGHT, this.Cells.get(Cell.Orientation.DOWN)!)
						this.Cells.set(Cell.Orientation.UPRIGHT, this.Cells.get(Cell.Orientation.NONE)!)
						this.Cells.set(Cell.Orientation.UP, this.Cells.get(Cell.Orientation.LEFT)!)
						this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.DOWNLEFT)!)
						this.Cells.delete(Cell.Orientation.DOWNRIGHT)
						this.Cells.delete(Cell.Orientation.DOWN)
						this.Cells.delete(Cell.Orientation.DOWNLEFT)
						this.Cells.delete(Cell.Orientation.LEFT)
						this.Cells.delete(Cell.Orientation.UPLEFT)
						break;
					case Cell.Orientation.LEFT:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Cell.Orientation.UPLEFT)!,
							contig.get(Cell.Orientation.LEFT)!,
							contig.get(Cell.Orientation.DOWNLEFT)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Cell.Orientation.UPRIGHT)!,
							this.Cells.get(Cell.Orientation.RIGHT)!,
							this.Cells.get(Cell.Orientation.DOWNRIGHT)!,
						)
						// shift preloaded
						this.Cells.set(Cell.Orientation.UPRIGHT, this.Cells.get(Cell.Orientation.UP)!)
						this.Cells.set(Cell.Orientation.RIGHT, this.Cells.get(Cell.Orientation.NONE)!)
						this.Cells.set(Cell.Orientation.DOWNRIGHT, this.Cells.get(Cell.Orientation.DOWN)!)
						this.Cells.set(Cell.Orientation.UP, this.Cells.get(Cell.Orientation.UPLEFT)!)
						this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.LEFT)!)
						this.Cells.set(Cell.Orientation.DOWN, this.Cells.get(Cell.Orientation.DOWNLEFT)!)
						this.Cells.delete(Cell.Orientation.UPLEFT)
						this.Cells.delete(Cell.Orientation.LEFT)
						this.Cells.delete(Cell.Orientation.DOWNLEFT)
						break;
					case Cell.Orientation.UPLEFT:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Cell.Orientation.UPRIGHT)!,
							contig.get(Cell.Orientation.UP)!,
							contig.get(Cell.Orientation.UPLEFT)!,
							contig.get(Cell.Orientation.LEFT)!,
							contig.get(Cell.Orientation.DOWNLEFT)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Cell.Orientation.UPRIGHT)!,
							this.Cells.get(Cell.Orientation.RIGHT)!,
							this.Cells.get(Cell.Orientation.DOWNRIGHT)!,
							this.Cells.get(Cell.Orientation.DOWN)!,
							this.Cells.get(Cell.Orientation.DOWNLEFT)!,
						)
						// shift preloaded
						this.Cells.set(Cell.Orientation.RIGHT, this.Cells.get(Cell.Orientation.UP)!)
						this.Cells.set(Cell.Orientation.DOWNRIGHT, this.Cells.get(Cell.Orientation.NONE)!)
						this.Cells.set(Cell.Orientation.DOWN, this.Cells.get(Cell.Orientation.LEFT)!)
						this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.UPLEFT)!)
						this.Cells.delete(Cell.Orientation.UPRIGHT)
						this.Cells.delete(Cell.Orientation.UP)
						this.Cells.delete(Cell.Orientation.UPLEFT)
						this.Cells.delete(Cell.Orientation.LEFT)
						this.Cells.delete(Cell.Orientation.DOWNLEFT)
						break;
				}

				// update border after none cell is up to date
				const cn = this.Cells.get(Cell.Orientation.NONE)!
				this.Border.set(Cell.Orientation.UP, cn.Cell.getY() * this.World.getCellheight())
				this.Border.set(Cell.Orientation.RIGHT, (cn.Cell.getX() + 1) * this.World.getCellwidth())
				this.Border.set(Cell.Orientation.DOWN, (cn.Cell.getY() + 1) * this.World.getCellheight())
				this.Border.set(Cell.Orientation.LEFT, cn.Cell.getX() * this.World.getCellwidth())

				this.cleanCells(deletedCells).then(() => { console.log('finish to destroy unused tilemaps') })

				// reassign cellsbyid with new cells
				this.CellsByID.clear()
				this.Cells.forEach((v, k) => {
					if (v) {
						this.CellsByID.set(ulid(v.Cell.getId_asU8()), k)
					}
				})

				return this.listCell((() => {
					const req = new CellDTO.ListCellReq()

					req.setIdsList(newCellIDs.filter((v) => (!(v == undefined))))

					return req
				})())
			})
			.then((cells: CellDTO.ListCellResp) => {

				// map new loaded cells
				const cellMap = new Map<string, Cell.Cell>()
				cells.getCellsList().map((v) => {
					cellMap.set(ulid(v.getId_asU8()), v)
				})
				if (o == Cell.Orientation.NONE) {
					const c = this.Cells.get(Cell.Orientation.NONE)!
					cellMap.set(ulid(c?.Cell.getId_asU8()), c.Cell)
				}

				const contig = this.Cells.get(Cell.Orientation.NONE)?.Cell.getContiguousMap() as jspb.Map<number, Uint8Array>
				if (o == Cell.Orientation.NONE) {
					contig.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.NONE)?.Cell.getId_asU8()!)
				}

				// must fit above loaded cells array
				const loadedCells = new Array<Orientation>()
				switch (o) {
					case Cell.Orientation.NONE:
						loadedCells.push(
							Cell.Orientation.NONE,
							Cell.Orientation.UP,
							Cell.Orientation.UPRIGHT,
							Cell.Orientation.RIGHT,
							Cell.Orientation.DOWNRIGHT,
							Cell.Orientation.DOWN,
							Cell.Orientation.DOWNLEFT,
							Cell.Orientation.LEFT,
							Cell.Orientation.UPLEFT
						)
						break;
					case Cell.Orientation.UP:
						loadedCells.push(
							Cell.Orientation.UP,
							Cell.Orientation.UPRIGHT,
							Cell.Orientation.UPLEFT
						)
						break;
					case Cell.Orientation.UPRIGHT:
						loadedCells.push(
							Cell.Orientation.UP,
							Cell.Orientation.UPRIGHT,
							Cell.Orientation.RIGHT,
							Cell.Orientation.DOWNRIGHT,
							Cell.Orientation.UPLEFT
						)
						break;
					case Cell.Orientation.RIGHT:
						loadedCells.push(
							Cell.Orientation.UPRIGHT,
							Cell.Orientation.RIGHT,
							Cell.Orientation.DOWNRIGHT,
						)
						break;
					case Cell.Orientation.DOWNRIGHT:
						loadedCells.push(
							Cell.Orientation.UPRIGHT,
							Cell.Orientation.RIGHT,
							Cell.Orientation.DOWNRIGHT,
							Cell.Orientation.DOWN,
							Cell.Orientation.DOWNLEFT,
						)
						break;
					case Cell.Orientation.DOWN:
						loadedCells.push(
							Cell.Orientation.DOWNRIGHT,
							Cell.Orientation.DOWN,
							Cell.Orientation.DOWNLEFT,
						)
						break;
					case Cell.Orientation.DOWNLEFT:
						loadedCells.push(
							Cell.Orientation.DOWNRIGHT,
							Cell.Orientation.DOWN,
							Cell.Orientation.DOWNLEFT,
							Cell.Orientation.LEFT,
							Cell.Orientation.UPLEFT
						)
						break;
					case Cell.Orientation.LEFT:
						loadedCells.push(
							Cell.Orientation.DOWNLEFT,
							Cell.Orientation.LEFT,
							Cell.Orientation.UPLEFT
						)
						break;
					case Cell.Orientation.UPLEFT:
						loadedCells.push(
							Cell.Orientation.UP,
							Cell.Orientation.UPRIGHT,
							Cell.Orientation.DOWNLEFT,
							Cell.Orientation.LEFT,
							Cell.Orientation.UPLEFT
						)
						break;
				}

				let loaded = 0
				loadedCells.map((o) => {
					const resetLoad = () => {
						loaded++
						if (loaded == loadedCells.length) {
							console.log('loading unlock')
							this.Loading = undefined
						}
					}

					if (!contig.has(o)) {
						// world border
						console.log('contig has no:', o)
						resetLoad()
						return
					}

					const c = cellMap.get(ulid(contig.get(o)!))!
					if (!c) {
						// world border
						// TO INVESTIGATE, shouldn't happen ?
						console.log('cell not found:', o)
						return
					}

					const tm = ulid(c.getTilemap_asU8())
					this.CellLoader.tilemapTiledJSON(tm, 'json/assets/' + tm + '.json')

					this.Cells.set(o, {
						Cell: c,
						Tilemap: this.Blank.Tilemap,
						Layers: new Map(),
						Colliders: new Map(),
						Objects: new Map(),
					})
					this.CellsByID.set(ulid(c.getId_asU8()), o)

					const loadTM = () => {
						console.log('load_tilemap ', o)

						// create new cell
						const map = this.make.tilemap({ key: tm, width: this.World.getCellwidth(), height: this.World.getCellheight() })

						let sets = new Array<Phaser.Tilemaps.Tileset>()

						const loadTS = (tsName: string) => {
							const ts = map.addTilesetImage(tsName)

							if (!ts) {
								console.log('failed to add tileset image ' + tsName)
								return
							}

							sets.push(ts)

							console.log('loaded TS ', sets.length, '/', map.tilesets.length)
							if (sets.length < map.tilesets.length) {
								return
							}

							const cc = this.Cells.get(o)
							if (!cc) {
								// should never happen
								return
							}

							cc.Tilemap = map

							const x = c.getX() * this.World.getCellwidth()
							const y = c.getY() * this.World.getCellheight()

							map.layers.map((l) => {
								const layer = map.createLayer(l.name, sets, x, y)
								if (!layer) {
									console.log('failed to create layer:', l.name, x, y)
									return
								}

								// TODO: required or not ?
								layer.setPipeline('main')

								console.log('created layer:', l.name, x, y)
								cc.Layers.set(l.name, layer)
							})

							// add objects
							map.objects.map((os) => {
								const objects = map.createFromObjects(os.name, {}, false) as Phaser.GameObjects.Sprite[]

								// offset on layer position
								const group = this.physics.add.staticGroup(objects.map((o) => {
									o.x += x
									o.y += y

									return o
								}))
								cc.Objects.set(os.name, group)

								const collider = this.physics.add.collider(this.Entity.Body, group)

								console.log('created object layer:', os.name)
								cc.Colliders.set(os.name, collider)
							})

							// Loading reset
							// Add safe concurrency for next 2 instructions
							resetLoad()
						}

						map.tilesets.map((ts) => {
							if (this.textures.exists(ts.name)) {
								loadTS(ts.name)
							} else {
								this.CellLoader.image(ts.name, 'img/assets/' + ts.name + '.png')
								console.log('add listener on ', 'filecomplete-image-' + ts.name)
								this.CellLoader.on('filecomplete-image-' + ts.name, () => {
									console.log('image completed on ', o)
									loadTS(ts.name)
								})
							}
						})
					}

					// Check if files already been loaded
					if (this.cache.tilemap.exists(tm)) {
						loadTM()
					} else {
						console.log('add listener on ', 'filecomplete-tilemapJSON-' + tm)
						this.CellLoader.on('filecomplete-tilemapJSON-' + tm, () => {
							console.log('json completed on ', o)
							loadTM()
						})
					}
				})

				this.CellLoader.on('complete', () => {
					console.log('reset loader')
					this.CellLoader.removeAllListeners()
					this.CellLoader.reset()
				})

				console.log('start global cell loading')
				this.CellLoader.start()
			})
	}

	async cleanCells(cells: GraphicCell[]) {
		cells.map((c) => {
			if (c) {
				console.log('clean cell', ulid(c.Cell.getId_asU8()))
				c.Colliders.forEach((v, k) => {
					if (v.world) {
						v.destroy()
					}
				})

				c.Objects.forEach((v, k) => {
					if (v.world) {
						v.destroy(true, true)
					}
				})

				c.Tilemap.destroy()
			}
		})
	}

	async cleanEntities() {
		let ids: string[] = new Array()

		this.EntityLastTick.forEach((val, key) => {
			if (val == 0) {
				ids.push(key)
			} else {
				this.EntityLastTick.set(key, 0)
			}
		})

		console.log('cleaning entities', ids)

		ids.forEach((id) => {
			let e = this.Entities.get(id)!

			this.EntityLastTick.delete(id)

			// TODO: ensure it cleans everything correctly ?
			e.Sprite.destroy()

			// remove collisions
			e.Colliders.forEach((v, k) => {
				if (v.world) {
					v.destroy()
				}
			})

			e.Objects.forEach((v, k) => {
				if (v.world) {
					v.destroy(true, true)
				}
			})

			// TODO: clean animations if not used elsewhere ?

			this.Entities.delete(id)
		})
	}

	// Connect
	async connect() {
		const release = await this.LoadMapMutex.acquire()
		// call list entities on current cell
		return this.connectPC((message: EntityDTO.ListEntityResp) => {
			// load all entities
			let entityIDs: Uint8Array[] = []

			message.getEntitiesList().forEach((entry: Entity.E) => {
				const id = ulid(entry.getId_asU8())

				// Set last tick to 0
				this.EntityLastTick.set(id, (this.EntityLastTick.get(id) || 0) + 1)

				if (this.Entities.has(id)) {
					// update state only
					const x = entry.getX() //+ (this.Cells.get(this.CellsByID.get(ulid(entry.getCellid_asU8()))!)?.Cell.getX()! * this.World.getCellwidth())
					const y = entry.getY() //+ (this.Cells.get(this.CellsByID.get(ulid(entry.getCellid_asU8()))!)?.Cell.getY()! * this.World.getCellheight())
					updateGraphicEntity(this.Entities.get(id)!, x, y)

					this.Entities.get(id)!.E = entry
				} else {
					// create associated sprite
					// receive own entity from server once and initialize
					if (id == ulid(this.Entity.E.getId_asU8())) {
						this.Entity.Body.destroy()

						if (entry.getObjectsList().length == 0) {
							console.log('player entity has no collision body. set default')
							this.Entity.Body = this.physics.add.sprite(entry.getX(), entry.getY(), id).setSize(16, 16).setOffset(0, 0)
						} else {
							// pick first dynamic body from list to assign as main collision object
							const obj = entry.getObjectsList().at(0)!
							this.Entity.Body = this.physics.add.sprite(entry.getX(), entry.getY(), id).
								setSize(obj.getWidth(), obj.getHeight()).
								setOffset(obj.getX(), obj.getY())
						}

						console.log('set body from server info')
						this.Entity.E.setX(entry.getX())
						this.Entity.E.setY(entry.getY())

						// local player sprite loaded, start camera follow
						this.cameras.main.startFollow(this.Entity.Body)
						this.Entity.Body.setDepth(entitySpriteDepth)

						this.Entities.set(id, {
							E: entry,
							Animations: new Map(),
							Direction: new Phaser.Geom.Point(entry.getX(), entry.getY()),
							Interpolation: 0.00,
							Sprite: this.Entity.Body,

							Objects: new Map(),
							Colliders: new Map(),
						})

						this.syncEntity()
						release()
					} else {
						const sprite = this.add.sprite(entry.getX(), entry.getY(), id)
						sprite.setDepth(entitySpriteDepth - 1)

						let ge: GraphicEntity = {
							E: entry,
							Animations: new Map(),
							Direction: new Phaser.Geom.Point(entry.getX(), entry.getY()),
							Interpolation: 0.00, // default speed
							Sprite: sprite,

							Objects: new Map(),
							Colliders: new Map(),
						}

						console.log('set entity: ', id, entry.getX(), entry.getY())

						// set collision objects
						// offset on layer position
						const objects = entry.getObjectsList()
						if (objects.length > 0) {
							const group = this.physics.add.staticGroup(objects.map((b) => {
								console.log('set entity collision: ', id, b.getX() + ge.E.getX(), b.getY() + ge.E.getY())
								return this.physics.add.
									staticImage(
										ge.E.getX() + b.getX(),
										ge.E.getY() + b.getY(),
										'').
									setSize(b.getWidth(), b.getHeight()).
									setVisible(false).
									setImmovable(true).
									setOffset(0, 0)
							}))

							const collider = this.physics.add.collider(this.Entity.Body, group)

							ge.Objects.set(id, group)
							ge.Colliders.set(id, collider)
							console.log('created entity collision:', id)
						}

						this.Entities.set(id, ge)
					}

					// load entity animations
					entityIDs.push(entry.getId_asU8())
				}
			})

			if (entityIDs.length == 0) {
				return
			}

			// load all animations
			this.listAnimation((() => {
				const req = new AnimationDTO.ListAnimationReq()

				req.setEntityidsList(entityIDs)
				req.setSize(1000) // TODO: less random

				return req
			})())
				.then((animations: AnimationDTO.ListAnimationResp) => {
					// load all current animations
					animations.getAnimationsList().forEach((an: Animation.Animation) => {
						const id = ulid(an.getId_asU8())
						const sheetID = ulid(an.getSheetid_asU8())
						const duplicateID = ulid(an.getDuplicateid_asU8())
						const entityID = ulid(an.getEntityid_asU8())

						let anims = this.Entities.get(entityID)?.Animations!

						// loading self animations, change object + add named animations
						if (entityID == ulid(this.Entity.E.getId_asU8())) {
							anims = this.Entity.Animations

							console.log('set named animation:', an.getName(), id)
							anims.set(an.getName(), id)
						}


						// return if already loaded
						if (this.anims.exists(duplicateID)) {
							anims.set(id, duplicateID)

							return
						}

						const loadAnim = () => {
							// Create animation
							const seq = an.getSequenceList().length > 0 ? { frames: an.getSequenceList() } : { start: an.getStart(), end: an.getEnd() };
							const newAnim = this.anims.create({
								key: duplicateID,
								frames: this.anims.generateFrameNumbers(
									sheetID,
									seq,
								),
								frameRate: an.getRate(),
								repeat: -1,
							})
							if (!newAnim) {
								console.log('failed to load animation ' + duplicateID)

								return
							}

							console.log('set animation:', id, duplicateID)

							// Add animation to entity animations
							anims.set(id, duplicateID)
						}

						if (!this.SpriteSheets.get(sheetID)) {
							// load sprite sheet

							this.EntityLoader.spritesheet(sheetID, 'img/assets/' + sheetID + '.png', {
								frameWidth: an.getFramewidth(),
								frameHeight: an.getFrameheight(),
								startFrame: an.getFramestart(),
								endFrame: an.getFrameend(),
								margin: an.getFramemargin(),
								spacing: an.getFramespacing(),
							}).once('filecomplete-spritesheet-' + sheetID, () => {
								// Add spritesheet to loaded
								this.SpriteSheets.set(sheetID, 1)

								loadAnim()
							}).start()
						} else {
							loadAnim()
						}
					})
				})
		})
	}

	// Start local entity sync at regular intervals
	syncEntity() {
		this.SyncTimer = window.setInterval(() => {
			this.EntityClient.send(this.Entity.E)
		}, 500)
	}

	// Stop local entity sync at regular intervals
	stopSyncEntity() {
		clearInterval(this.SyncTimer)
		this.EntityClient.finishSend()
	}

	// Connect server update
	connectUpdate() {
		// call update entity
		this.EntityClient = grpc.client(API.UpdateEntity, {
			host: 'https://api.legacyfactory.com:8082',
			// transport: grpc.WebsocketTransport(),
		});

		let md = new grpc.Metadata()
		md.set('token', this.registry.get('token'))
		this.EntityClient.start(md)
	}

	// API PC
	connectPC(callback: (message: EntityDTO.ListEntityResp) => void) {
		let req = this.PC;
		let md = new grpc.Metadata()
		md.set('token', this.registry.get('token'))

		const prom = new Promise<string>((resolve, reject) => {
			grpc.invoke(API.ConnectPC, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onMessage: (message: EntityDTO.ListEntityResp) => {
					callback(message)
				},
				onEnd: (code: grpc.Code, message: string | undefined, trailers: grpc.Metadata) => {
					if (code !== grpc.Code.OK || !message) {
						reject('connectPC:' + message)

						return
					}

					resolve(message)
				}
			});
		})

		return prom
	}

	// API Cell
	listCell(req: CellDTO.ListCellReq) {
		let md = new grpc.Metadata()
		md.set('token', this.registry.get('token'))

		const prom = new Promise<CellDTO.ListCellResp>((resolve, reject) => {
			grpc.unary(API.ListCell, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject('listCell:' + res)

						return
					}

					resolve(message as CellDTO.ListCellResp)
				}
			});
		})

		return prom
	}

	// API Entity
	listEntity(req: EntityDTO.ListEntityReq) {
		let md = new grpc.Metadata()
		md.set('token', this.registry.get('token'))

		const prom = new Promise<EntityDTO.ListEntityResp>((resolve, reject) => {
			grpc.unary(API.ListEntity, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject('listEntity:' + res)

						return
					}

					resolve(message as EntityDTO.ListEntityResp)
				}
			});
		})

		return prom
	}

	// API Animation
	listAnimation(req: AnimationDTO.ListAnimationReq) {
		let md = new grpc.Metadata()
		md.set('token', this.registry.get('token'))

		const prom = new Promise<AnimationDTO.ListAnimationResp>((resolve, reject) => {
			grpc.unary(API.ListAnimation, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject('listAnimation:' + res)

						return
					}

					resolve(message as AnimationDTO.ListAnimationResp)
				}
			});
		})

		return prom
	}

	// API World
	listWorld(req: WorldDTO.ListWorldReq) {
		let md = new grpc.Metadata()
		md.set('token', this.registry.get('token'))

		const prom = new Promise<WorldDTO.ListWorldResp>((resolve, reject) => {
			grpc.unary(API.ListWorld, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject('listWorld:' + res)

						return
					}

					resolve(message as WorldDTO.ListWorldResp)
				}
			});
		})

		return prom
	}
}
