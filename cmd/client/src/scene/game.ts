import { Scene } from "phaser";
import { grpc } from '@improbable-eng/grpc-web';

import { Mutex } from 'async-mutex';

import * as jspb from "google-protobuf";

import { ulid } from '../lib/ulid'

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

enum Orientation {
	None = 0,
	Up,
	UpRight,
	Right,
	DownRight,
	Down,
	DownLeft,
	Left,
	UpLeft,
}


interface properties {
	name: string
	value: boolean
}

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
	Colliders: Map<string, Phaser.Physics.Arcade.Collider>
	Layers: Map<string, Phaser.Tilemaps.TilemapLayer | null>
}

type Controls = {
	Up: Phaser.Input.Keyboard.Key
	Down: Phaser.Input.Keyboard.Key
	Left: Phaser.Input.Keyboard.Key
	Right: Phaser.Input.Keyboard.Key
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
	// Loading mutex
	Loading: Uint8Array | undefined

	// Entities
	Entities: Map<string, GraphicEntity>
	// Specific loader for entities
	EntityLoader: Phaser.Loader.LoaderPlugin
	// 2 buffers for entity update
	// Rendering buffer is frozen and used by update only
	EntityBuffer: Array<string>
	EntityBufferRendering: Array<string>
	// Grpc client to send entity updates to server
	EntityClient: grpc.Client<Entity.E, Entity.E>
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
			Orientation: Orientation.Down
		}
		this.EntityID = ulid(this.Entity.E.getId_asU8())

		this.Cells = new Map()
		this.CellsByID = new Map()
		this.CellLoader = new Phaser.Loader.LoaderPlugin(this)
		this.Border = new Map()

		this.Entities = new Map()
		this.EntityLoader = new Phaser.Loader.LoaderPlugin(this)
		this.EntityBuffer = new Array<string>()
		this.EntityBufferRendering = new Array<string>()

		this.SpriteSheets = new Map()

		this.World = new World.World()

		this.LoadMapMutex = new Mutex()

		const m = this.make.tilemap()
		this.Blank = {
			Cell: new Cell.Cell,
			Tilemap: m,
			Layers: new Map(),
			Colliders: new Map()
		}
	}
	preload() { }
	create() {
		// this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.W).on('down', () => {this.movePC(Orientation.Up)})
		// this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.S).on('down', () => {this.movePC(Orientation.Down)})
		// this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.A).on('down', () => {this.movePC(Orientation.Left)})
		// this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.D).on('down', () => {this.movePC(Orientation.Right)})

		// this.Controls.Up = this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.W)
		// this.Controls.Down = this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.S)
		// this.Controls.Left = this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.A)
		// this.Controls.Right = this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.D)

		this.Cursors = this.input.keyboard!.createCursorKeys();

		// Connect for entity sync
		this.connect()
			.then(() => {
				console.log('disconnect from server')
			})

		// Connect for entity send
		this.connectUpdate()

		this.Loading = this.Entity.E.getCellid_asU8()

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
							console.log("failed to load world")
						} else {
							this.World = ws.getWorldsList()[0]
						}
					}).then(() => {
						return this.loadMap(Orientation.None)
					})
					.then(() => {
						this.load.start()
						this.load.on('complete', () => {
							console.log('load complete')
						})
					})
					.catch((err) => {
						console.log(err)
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
			this.Entity.Orientation = Orientation.Up
		} else if (this.Cursors.down.isDown) {
			animationID = this.Entity.Animations.get('walk_down')
			this.Entity.Body.body.setVelocityY(speed)
			this.Entity.Orientation = Orientation.Down
		}

		if (this.Cursors.right.isDown) {
			animationID = this.Entity.Animations.get('walk_right')
			this.Entity.Body.body.setVelocityX(speed)
			this.Entity.Orientation = Orientation.Right
		} else if (this.Cursors.left.isDown) {
			animationID = this.Entity.Animations.get('walk_left')
			this.Entity.Body.body.setVelocityX(-speed)
			this.Entity.Orientation = Orientation.Left
		}

		if (!animationID) {
			switch (this.Entity.Orientation) {
				case Orientation.Up:
					animationID = this.Entity.Animations.get('idle_up')
					break
				case Orientation.Right:
					animationID = this.Entity.Animations.get('idle_right')
					break
				case Orientation.Down:
					animationID = this.Entity.Animations.get('idle_down')
					break
				case Orientation.Left:
					animationID = this.Entity.Animations.get('idle_left')
					break
			}
		}

		if (animationID) {
			this.Entity?.Body?.play(animationID, true)
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

		let o = Orientation.None
		const x = this.Entity.Body.body.x
		const y = this.Entity.Body.body.y
		const up = this.Border.get(Orientation.Up)!
		const right = this.Border.get(Orientation.Right)!
		const down = this.Border.get(Orientation.Down)!
		const left = this.Border.get(Orientation.Left)!

		if (y < up) {
			if (x < left) {
				o = Orientation.UpLeft
			} else if (x < right) {
				o = Orientation.Up
			} else {
				o = Orientation.UpRight
			}
		} else if (y < down) {
			if (x < left) {
				o = Orientation.Left
			} else if (x < right) {
				o = Orientation.None
			} else {
				o = Orientation.Right
			}
		} else {
			if (x < left) {
				o = Orientation.DownLeft
			} else if (x < right) {
				o = Orientation.Down
			} else {
				o = Orientation.DownRight
			}
		}

		if (o != Orientation.None && this.Loading == undefined) {
			const id = this.Cells.get(o)?.Cell.getId_asU8()
			this.Loading = id

			console.log('loading new direction:', o, o.toString(), id)

			if (id) {
				this.Entity.E.setCellid(id)

				this.loadMap(o).then(() => {
					console.log('loaded cell')
				})
			} else {
				// don't move entity, out of world
			}
		}

		// Move entity
		this.Entity.E.setX(Math.round(x))
		this.Entity.E.setY(Math.round(y))

		// Swap buffers
		// this.EntityBufferRendering = this.EntityBuffer
		// this.EntityBuffer = []

		// Entity queue update
		// this.EntityBufferRendering.forEach((value: string) => {

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

		// this.EntityBufferRendering = []
	}

	// Orientation o uses preload surrounded cells
	// Orientation.None loads everything
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
					case Orientation.None:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Orientation.Up)!,
							contig.get(Orientation.UpRight)!,
							contig.get(Orientation.Right)!,
							contig.get(Orientation.DownRight)!,
							contig.get(Orientation.Down)!,
							contig.get(Orientation.DownLeft)!,
							contig.get(Orientation.Left)!,
							contig.get(Orientation.UpLeft)!,
						)

						// assign cells to delete
						// deletedCells.push()
						// create new blank graphic cell from loaded cell
						const m = this.make.tilemap()
						this.Cells.set(Orientation.None, {
							Cell: c,
							Tilemap: this.Blank.Tilemap,
							Layers: this.Blank.Layers,
							Colliders: this.Blank.Colliders
						})

						this.Cells.delete(Orientation.Up)
						this.Cells.delete(Orientation.UpRight)
						this.Cells.delete(Orientation.Right)
						this.Cells.delete(Orientation.DownRight)
						this.Cells.delete(Orientation.Down)
						this.Cells.delete(Orientation.DownLeft)
						this.Cells.delete(Orientation.Left)
						this.Cells.delete(Orientation.UpLeft)
						break;
					case Orientation.Up:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Orientation.UpLeft)!,
							contig.get(Orientation.Up)!,
							contig.get(Orientation.UpRight)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Orientation.DownLeft)!,
							this.Cells.get(Orientation.Down)!,
							this.Cells.get(Orientation.DownRight)!,
						)
						// shift preloaded
						this.Cells.set(Orientation.DownLeft, this.Cells.get(Orientation.Left)!)
						this.Cells.set(Orientation.Down, this.Cells.get(Orientation.None)!)
						this.Cells.set(Orientation.DownRight, this.Cells.get(Orientation.Right)!)
						this.Cells.set(Orientation.Left, this.Cells.get(Orientation.UpLeft)!)
						this.Cells.set(Orientation.None, this.Cells.get(Orientation.Up)!)
						this.Cells.set(Orientation.Right, this.Cells.get(Orientation.UpRight)!)
						this.Cells.delete(Orientation.UpLeft)
						this.Cells.delete(Orientation.Up)
						this.Cells.delete(Orientation.UpRight)
						break;
					case Orientation.UpRight:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Orientation.UpLeft)!,
							contig.get(Orientation.Up)!,
							contig.get(Orientation.UpRight)!,
							contig.get(Orientation.Right)!,
							contig.get(Orientation.DownRight)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Orientation.UpLeft)!,
							this.Cells.get(Orientation.Left)!,
							this.Cells.get(Orientation.DownLeft)!,
							this.Cells.get(Orientation.Down)!,
							this.Cells.get(Orientation.DownRight)!,
						)
						// shift preloaded
						this.Cells.set(Orientation.Left, this.Cells.get(Orientation.Up)!)
						this.Cells.set(Orientation.DownLeft, this.Cells.get(Orientation.None)!)
						this.Cells.set(Orientation.Down, this.Cells.get(Orientation.Right)!)
						this.Cells.set(Orientation.None, this.Cells.get(Orientation.UpRight)!)
						this.Cells.delete(Orientation.UpLeft)
						this.Cells.delete(Orientation.Up)
						this.Cells.delete(Orientation.UpRight)
						this.Cells.delete(Orientation.Right)
						this.Cells.delete(Orientation.DownRight)
						break;
					case Orientation.Right:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Orientation.UpRight)!,
							contig.get(Orientation.Right)!,
							contig.get(Orientation.DownRight)!
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Orientation.UpLeft)!,
							this.Cells.get(Orientation.Left)!,
							this.Cells.get(Orientation.DownLeft)!,
						)
						// shift preloaded
						this.Cells.set(Orientation.UpLeft, this.Cells.get(Orientation.Up)!)
						this.Cells.set(Orientation.Left, this.Cells.get(Orientation.None)!)
						this.Cells.set(Orientation.DownLeft, this.Cells.get(Orientation.Down)!)
						this.Cells.set(Orientation.Up, this.Cells.get(Orientation.UpRight)!)
						this.Cells.set(Orientation.None, this.Cells.get(Orientation.Right)!)
						this.Cells.set(Orientation.Down, this.Cells.get(Orientation.DownRight)!)
						this.Cells.delete(Orientation.UpRight)
						this.Cells.delete(Orientation.Right)
						this.Cells.delete(Orientation.DownRight)
						break;
					case Orientation.DownRight:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Orientation.DownLeft)!,
							contig.get(Orientation.Down)!,
							contig.get(Orientation.DownRight)!,
							contig.get(Orientation.Right)!,
							contig.get(Orientation.UpRight)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Orientation.UpLeft)!,
							this.Cells.get(Orientation.Left)!,
							this.Cells.get(Orientation.DownLeft)!,
							this.Cells.get(Orientation.Up)!,
							this.Cells.get(Orientation.UpRight)!,
						)
						// shift preloaded
						this.Cells.set(Orientation.Left, this.Cells.get(Orientation.Down)!)
						this.Cells.set(Orientation.UpLeft, this.Cells.get(Orientation.None)!)
						this.Cells.set(Orientation.Up, this.Cells.get(Orientation.Right)!)
						this.Cells.set(Orientation.None, this.Cells.get(Orientation.DownRight)!)
						this.Cells.delete(Orientation.DownLeft)
						this.Cells.delete(Orientation.Down)
						this.Cells.delete(Orientation.DownRight)
						this.Cells.delete(Orientation.Right)
						this.Cells.delete(Orientation.UpRight)
						break;
					case Orientation.Down:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Orientation.DownLeft)!,
							contig.get(Orientation.Down)!,
							contig.get(Orientation.DownRight)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Orientation.UpLeft)!,
							this.Cells.get(Orientation.Up)!,
							this.Cells.get(Orientation.UpRight)!,
						)
						// shift preloaded
						this.Cells.set(Orientation.UpLeft, this.Cells.get(Orientation.Left)!)
						this.Cells.set(Orientation.Up, this.Cells.get(Orientation.None)!)
						this.Cells.set(Orientation.UpRight, this.Cells.get(Orientation.Right)!)
						this.Cells.set(Orientation.Left, this.Cells.get(Orientation.DownLeft)!)
						this.Cells.set(Orientation.None, this.Cells.get(Orientation.Down)!)
						this.Cells.set(Orientation.Right, this.Cells.get(Orientation.DownRight)!)
						this.Cells.delete(Orientation.DownLeft)
						this.Cells.delete(Orientation.Down)
						this.Cells.delete(Orientation.DownRight)
						break;
					case Orientation.DownLeft:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Orientation.DownRight)!,
							contig.get(Orientation.Down)!,
							contig.get(Orientation.DownLeft)!,
							contig.get(Orientation.Left)!,
							contig.get(Orientation.UpLeft)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Orientation.UpLeft)!,
							this.Cells.get(Orientation.Up)!,
							this.Cells.get(Orientation.UpRight)!,
							this.Cells.get(Orientation.Right)!,
							this.Cells.get(Orientation.DownRight)!,
						)
						// shift preloaded
						this.Cells.set(Orientation.Right, this.Cells.get(Orientation.Down)!)
						this.Cells.set(Orientation.UpRight, this.Cells.get(Orientation.None)!)
						this.Cells.set(Orientation.Up, this.Cells.get(Orientation.Left)!)
						this.Cells.set(Orientation.None, this.Cells.get(Orientation.DownLeft)!)
						this.Cells.delete(Orientation.DownRight)
						this.Cells.delete(Orientation.Down)
						this.Cells.delete(Orientation.DownLeft)
						this.Cells.delete(Orientation.Left)
						this.Cells.delete(Orientation.UpLeft)
						break;
					case Orientation.Left:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Orientation.UpLeft)!,
							contig.get(Orientation.Left)!,
							contig.get(Orientation.DownLeft)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Orientation.UpRight)!,
							this.Cells.get(Orientation.Right)!,
							this.Cells.get(Orientation.DownRight)!,
						)
						// shift preloaded
						this.Cells.set(Orientation.UpRight, this.Cells.get(Orientation.Up)!)
						this.Cells.set(Orientation.Right, this.Cells.get(Orientation.None)!)
						this.Cells.set(Orientation.DownRight, this.Cells.get(Orientation.Down)!)
						this.Cells.set(Orientation.Up, this.Cells.get(Orientation.UpLeft)!)
						this.Cells.set(Orientation.None, this.Cells.get(Orientation.Left)!)
						this.Cells.set(Orientation.Down, this.Cells.get(Orientation.DownLeft)!)
						this.Cells.delete(Orientation.UpLeft)
						this.Cells.delete(Orientation.Left)
						this.Cells.delete(Orientation.DownLeft)
						break;
					case Orientation.UpLeft:
						// assign new cells to load
						newCellIDs.push(
							contig.get(Orientation.UpRight)!,
							contig.get(Orientation.Up)!,
							contig.get(Orientation.UpLeft)!,
							contig.get(Orientation.Left)!,
							contig.get(Orientation.DownLeft)!,
						)

						// assign cells to delete
						deletedCells.push(
							this.Cells.get(Orientation.UpRight)!,
							this.Cells.get(Orientation.Right)!,
							this.Cells.get(Orientation.DownRight)!,
							this.Cells.get(Orientation.Down)!,
							this.Cells.get(Orientation.DownLeft)!,
						)
						// shift preloaded
						this.Cells.set(Orientation.Right, this.Cells.get(Orientation.Up)!)
						this.Cells.set(Orientation.DownRight, this.Cells.get(Orientation.None)!)
						this.Cells.set(Orientation.Down, this.Cells.get(Orientation.Left)!)
						this.Cells.set(Orientation.None, this.Cells.get(Orientation.UpLeft)!)
						this.Cells.delete(Orientation.UpRight)
						this.Cells.delete(Orientation.Up)
						this.Cells.delete(Orientation.UpLeft)
						this.Cells.delete(Orientation.Left)
						this.Cells.delete(Orientation.DownLeft)
						break;
				}

				// update border after none cell is up to date
				const cn = this.Cells.get(Orientation.None)!
				this.Border.set(Orientation.Up, cn.Cell.getY() * this.World.getCellheight())
				this.Border.set(Orientation.Right, (cn.Cell.getX() + 1) * this.World.getCellwidth())
				this.Border.set(Orientation.Down, (cn.Cell.getY() + 1) * this.World.getCellheight())
				this.Border.set(Orientation.Left, cn.Cell.getX() * this.World.getCellwidth())

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
				if (o == Orientation.None) {
					const c = this.Cells.get(Orientation.None)!
					cellMap.set(ulid(c?.Cell.getId_asU8()), c.Cell)
				}

				const contig = this.Cells.get(Orientation.None)?.Cell.getContiguousMap() as jspb.Map<number, Uint8Array>
				if (o == Orientation.None) {
					contig.set(Orientation.None, this.Cells.get(Orientation.None)?.Cell.getId_asU8()!)
				}

				// must fit above loaded cells array
				const loadedCells = new Array<Orientation>()
				switch (o) {
					case Orientation.None:
						loadedCells.push(
							Orientation.None,
							Orientation.Up,
							Orientation.UpRight,
							Orientation.Right,
							Orientation.DownRight,
							Orientation.Down,
							Orientation.DownLeft,
							Orientation.Left,
							Orientation.UpLeft
						)
						break;
					case Orientation.Up:
						loadedCells.push(
							Orientation.Up,
							Orientation.UpRight,
							Orientation.UpLeft
						)
						break;
					case Orientation.UpRight:
						loadedCells.push(
							Orientation.Up,
							Orientation.UpRight,
							Orientation.Right,
							Orientation.DownRight,
							Orientation.UpLeft
						)
						break;
					case Orientation.Right:
						loadedCells.push(
							Orientation.UpRight,
							Orientation.Right,
							Orientation.DownRight,
						)
						break;
					case Orientation.DownRight:
						loadedCells.push(
							Orientation.UpRight,
							Orientation.Right,
							Orientation.DownRight,
							Orientation.Down,
							Orientation.DownLeft,
						)
						break;
					case Orientation.Down:
						loadedCells.push(
							Orientation.DownRight,
							Orientation.Down,
							Orientation.DownLeft,
						)
						break;
					case Orientation.DownLeft:
						loadedCells.push(
							Orientation.DownRight,
							Orientation.Down,
							Orientation.DownLeft,
							Orientation.Left,
							Orientation.UpLeft
						)
						break;
					case Orientation.Left:
						loadedCells.push(
							Orientation.DownLeft,
							Orientation.Left,
							Orientation.UpLeft
						)
						break;
					case Orientation.UpLeft:
						loadedCells.push(
							Orientation.Up,
							Orientation.UpRight,
							Orientation.DownLeft,
							Orientation.Left,
							Orientation.UpLeft
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

							map.layers.map((l) => {
								const layer = map.createLayer(l.name, sets, c.getX() * this.World.getCellwidth(), c.getY() * this.World.getCellheight())
								if (!layer) {
									console.log('failed to create layer:', l.name, c.getX(), c.getY())
									return
								}

								// TODO: required or not ?
								layer.setPipeline('main')

								console.log('created layer:', l.name, c.getX(), c.getY())
								const props = layer.layer.properties as properties[]
								if (props.find((prop) => { return prop.name == 'collides' && prop.value })) {
									const collider = this.physics.add.collider(this.Entity.Body, layer.setCollisionByExclusion([-1]))
									cc.Colliders.set(l.name, collider)
								}
								cc.Layers.set(l.name, layer)
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

				c.Tilemap.destroy()
			}
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

				if (this.Entities.has(id)) {
					// update state only
					const x = entry.getX() + (this.Cells.get(this.CellsByID.get(ulid(entry.getCellid_asU8()))!)?.Cell.getX()! * this.World.getCellwidth())
					const y = entry.getX() + (this.Cells.get(this.CellsByID.get(ulid(entry.getCellid_asU8()))!)?.Cell.getY()! * this.World.getCellheight())
					updateGraphicEntity(this.Entities.get(id)!, x, y)
					console.log(this.CellsByID.get(ulid(entry.getCellid_asU8()))!)
					console.log(id, x, y)

					this.Entities.get(id)!.E = entry
					// this.EntityBuffer.push(ulid(entry.getId_asU8()))
				} else {
					// create associated sprite
					if (id == ulid(this.Entity.E.getId_asU8())) {
						this.Entity.Body.destroy()
						this.Entity.Body = this.physics.add.sprite(entry.getX(), entry.getY(), id)

						console.log('set body from server info')
						this.Entity.E.setX(entry.getX())
						this.Entity.E.setY(entry.getY())

						// local player sprite loaded, start camera follow
						this.cameras.main.startFollow(this.Entity.Body)
						this.Entity.Body.setDepth(entitySpriteDepth - 1)

						this.Entities.set(id, {
							E: entry,
							Animations: new Map(),
							Direction: new Phaser.Geom.Point(entry.getX(), entry.getY()),
							Interpolation: 0.00,
							Sprite: this.Entity.Body,
						})

						this.syncEntity()
						release()
					} else {
						const sprite = this.add.sprite(entry.getX(), entry.getY(), id)
						sprite.setDepth(entitySpriteDepth - 1)

						this.Entities.set(id, {
							E: entry,
							Animations: new Map(),
							Direction: new Phaser.Geom.Point(entry.getX(), entry.getY()),
							Interpolation: 0.00, // default speed
							Sprite: sprite,
						})
					}

					// load entity animations
					entityIDs.push(entry.getId_asU8())
					// this.EntityBuffer.push(ulid(entry.getId_asU8()))
				}

			})

			if (entityIDs.length == 0) {
				return
			}

			// load all animations
			this.listAnimation((() => {
				const req = new AnimationDTO.ListAnimationReq()

				req.setEntityidsList(entityIDs)
				req.setSize(1000)

				return req
			})())
				.then((animations: AnimationDTO.ListAnimationResp) => {
					// load all current animations
					animations.getAnimationsList().forEach((an: Animation.Animation) => {
						const id = ulid(an.getId_asU8())
						const sheetID = ulid(an.getSheetid_asU8())
						const duplicateID = ulid(an.getDuplicateid_asU8())
						const entityID = ulid(an.getEntityid_asU8())

						if (entityID == ulid(this.Entity.E.getId_asU8())) {
							this.Entity.Animations.set(an.getName(), duplicateID)
						}

						// return if already loaded
						if (this.anims.exists(duplicateID)) {
							this.Entities.get(entityID)?.Animations.set(id, duplicateID)

							// Play entity animation
							// this.EntityBuffer.push(ulid(an.getEntityid_asU8()))

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

							console.log('set animation:' + ulid(an.getEntityid_asU8()) + ":" + an.getName())

							// Add animation to entity animations
							this.Entities.get(entityID)?.Animations.set(id, duplicateID)

							// Play entity animation
							// this.EntityBuffer.push(ulid(an.getEntityid_asU8()))
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
