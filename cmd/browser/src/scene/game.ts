import {Loader, Scene} from "phaser";
import {grpc} from '@improbable-eng/grpc-web';

import * as jspb from "google-protobuf";

import {ulid} from '../lib/ulid'

import * as API from 'cmd/api/grpc/api_pb_service';

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

import { client } from "@improbable-eng/grpc-web/dist/typings/client";

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

type GraphicEntity = {
    E: Entity.E
    Sprite: Phaser.GameObjects.Sprite
}

type GraphicCell = {
    Cell: Cell.Cell
    Tilemap: Phaser.Tilemaps.Tilemap
    Layer: Phaser.Tilemaps.TilemapLayer
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
    Entity: GraphicEntity;
    EntityID: string;
    Cursors: Phaser.Types.Input.Keyboard.CursorKeys
    // Controls: Controls;

    // Cells & maps
    Cells: Map<Orientation, GraphicCell>
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

    // Others
    Animations: Map<string, string>

    // Spritesheets already loaded
    SpriteSheets: Map<string, integer>

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    init (pc: PCDTO.PC) {
        this.PC =  pc.getPc() as PC.PC;
        const e = pc.getEntity() as Entity.E;
        this.Entity = {
            E: e,
            Sprite: this.add.sprite(e.getX(), e.getY(), ulid(e.getId_asU8()))
        }
        this.EntityID = ulid(this.Entity.E.getId_asU8())

        this.Cells = new Map()
        this.CellLoader = new Phaser.Loader.LoaderPlugin(this)
        this.Border = new Map()

        this.Entities = new Map()
        this.EntityLoader = new Phaser.Loader.LoaderPlugin(this)
        this.EntityBuffer = new Array<string>()
        this.EntityBufferRendering = new Array<string>()

        this.Animations = new Map()

        this.SpriteSheets = new Map()

        this.World = new World.World()
    }
    preload() {}
    create() {
        // this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.W).on('down', () => {this.movePC(Orientation.Up)})
        // this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.S).on('down', () => {this.movePC(Orientation.Down)})
        // this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.A).on('down', () => {this.movePC(Orientation.Left)})
        // this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.D).on('down', () => {this.movePC(Orientation.Right)})

        // this.Controls.Up = this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.W)
        // this.Controls.Down = this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.S)
        // this.Controls.Left = this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.A)
        // this.Controls.Right = this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.D)

        this.Cursors = this.input.keyboard.createCursorKeys();

        // Connect for entity sync
        this.connect()
        .then(() => {
            console.log('disconnect from server')
        })

        // Connect for entity send
        this.connectUpdate()
        .then(() => {
            console.log('disconnect to server')
        })

        // Load world & cell sync
        this.listWorld((() => {
            const req = new WorldDTO.ListWorldReq()

            req.setIdsList([this.PC.getWorldid_asU8()])

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
        .catch((err)=>{
            console.log(err)
        })
    }

    update(time: number, deltaTime: number) {
        // Controls + local anim update
        let animationID = null
        let deltaX: number = 0
        let deltaY: number = 0
        const speed: number = 10

        // Move entity
        if (this.Cursors.up.isDown) {
            deltaY = -speed
            animationID = this.Animations.get(this.EntityID + ':' + 'up')
        } else if (this.Cursors.right.isDown) {
            deltaX = speed
            animationID = this.Animations.get(this.EntityID + ':' + 'right')
        } else if (this.Cursors.down.isDown) {
            deltaY = speed
            animationID = this.Animations.get(this.EntityID + ':' + 'down')
        } else if (this.Cursors.left.isDown) {
            deltaX = -speed
            animationID = this.Animations.get(this.EntityID + ':' + 'left')
        } else {
            animationID = this.Animations.get(this.EntityID + ':' + 'idle')
        }

        let o = Orientation.None
        const x = this.Entity.E.getX() + deltaX
        const y = this.Entity.E.getY() + deltaY
        const up = this.Border.get(Orientation.Up)!
        const right = this.Border.get(Orientation.Right)!
        const down = this.Border.get(Orientation.Down)!
        const left = this.Border.get(Orientation.Left)!

        // console.log(y, up, down)
        // console.log(x, right, left)

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

        if (o != Orientation.None && this.Loading != undefined) {
            const id = this.Cells.get(o)?.Cell.getId_asU8()
            this.Loading = id
            if (id) {
                this.Entity.E.setCellid(id)

                this.loadMap(o).then(() => {
                    console.log('loaded cell')
                })
            } else {
                // don't move entity, out of world
                // deltaX = 0
                // deltaY = 0
            }
        }

        // Move entity
        this.Entity.E.setX(this.Entity.E.getX() + deltaX)
        this.Entity.E.setY(this.Entity.E.getY() + deltaY)

        // Move entity client side
        this.Entity?.Sprite.setX(this.Entity.E.getX())
        this.Entity?.Sprite.setY(this.Entity.E.getY())
        if (animationID) {
            this.Entity?.Sprite.play(animationID, true)
        }

        // Send new entity to server
        // this.EntityClient?.send(this.Entity)

        // Swap buffers
        const temp = this.EntityBuffer
        this.EntityBuffer = this.EntityBufferRendering
        this.EntityBufferRendering = temp

        // Entity queue update
        // this.EntityBufferRendering.forEach((value: string) => {
        //     // Reconciliation for self
        //     if (value == ulid(this.Entity.getId_asU8())) {
        //         return
        //     }

        //     const e = this.Entities.get(value)
        //     e?.Sprite.play(ulid(e.E.getAnimationid_asU8()), true)
        // })

        this.EntityBufferRendering = []
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

            var cellIDs : Uint8Array[] = []

            switch (o) {
                case Orientation.None:
                    // create new blank graphic cell from loaded cell
                    const m = this.make.tilemap()
                    this.Cells.set(Orientation.None, {
                        Cell: c,
                        Tilemap: m,
                        Layer: m.createBlankLayer('blank', ''),
                    })
                    // assign new cells to load
                    cellIDs.push(
                        contig.get(Orientation.Up)!,
                        contig.get(Orientation.UpRight)!,
                        contig.get(Orientation.Right)!,
                        contig.get(Orientation.DownRight)!,
                        contig.get(Orientation.Down)!,
                        contig.get(Orientation.DownLeft)!,
                        contig.get(Orientation.Left)!,
                        contig.get(Orientation.UpLeft)!,
                    )
                    break;
                case Orientation.Up:
                    // shift preloaded
                    this.Cells.set(Orientation.DownLeft, this.Cells.get(Orientation.Left)!)
                    this.Cells.set(Orientation.Down, this.Cells.get(Orientation.None)!)
                    this.Cells.set(Orientation.DownRight, this.Cells.get(Orientation.Right)!)
                    this.Cells.set(Orientation.Left, this.Cells.get(Orientation.DownLeft)!)
                    this.Cells.set(Orientation.None, this.Cells.get(Orientation.Down)!)
                    this.Cells.set(Orientation.Right, this.Cells.get(Orientation.DownRight)!)
                    // assign new cells to load
                    cellIDs.push(
                        contig.get(Orientation.UpLeft)!,
                        contig.get(Orientation.Up)!,
                        contig.get(Orientation.UpRight)!,
                    )
                    break;
                case Orientation.UpRight:
                    // shift preloaded
                    this.Cells.set(Orientation.Left, this.Cells.get(Orientation.Up)!)
                    this.Cells.set(Orientation.DownLeft, this.Cells.get(Orientation.None)!)
                    this.Cells.set(Orientation.Down, this.Cells.get(Orientation.Right)!)
                    this.Cells.set(Orientation.None, this.Cells.get(Orientation.UpRight)!)
                    // assign new cells to load
                    cellIDs.push(
                        contig.get(Orientation.UpLeft)!,
                        contig.get(Orientation.Up)!,
                        contig.get(Orientation.UpRight)!,
                        contig.get(Orientation.Right)!,
                        contig.get(Orientation.DownRight)!,
                    )
                    break;
                case Orientation.Right:
                    // shift preloaded
                    this.Cells.set(Orientation.UpLeft, this.Cells.get(Orientation.Up)!)
                    this.Cells.set(Orientation.Up, this.Cells.get(Orientation.UpRight)!)
                    this.Cells.set(Orientation.Left, this.Cells.get(Orientation.None)!)
                    this.Cells.set(Orientation.None, this.Cells.get(Orientation.Right)!)
                    this.Cells.set(Orientation.DownLeft, this.Cells.get(Orientation.Down)!)
                    this.Cells.set(Orientation.Down, this.Cells.get(Orientation.DownRight)!)
                    // assign new cells to load
                    cellIDs.push(
                        contig.get(Orientation.UpRight)!,
                        contig.get(Orientation.Right)!,
                        contig.get(Orientation.DownRight)!
                    )
                    break;
                case Orientation.DownRight:
                    // shift preloaded
                    this.Cells.set(Orientation.Left, this.Cells.get(Orientation.Down)!)
                    this.Cells.set(Orientation.UpLeft, this.Cells.get(Orientation.None)!)
                    this.Cells.set(Orientation.Up, this.Cells.get(Orientation.Right)!)
                    this.Cells.set(Orientation.None, this.Cells.get(Orientation.DownRight)!)
                    // assign new cells to load
                    cellIDs.push(
                        contig.get(Orientation.DownLeft)!,
                        contig.get(Orientation.Down)!,
                        contig.get(Orientation.DownRight)!,
                        contig.get(Orientation.Right)!,
                        contig.get(Orientation.UpRight)!,
                    )
                    break;
                case Orientation.Down:
                    // shift preloaded
                    this.Cells.set(Orientation.UpLeft, this.Cells.get(Orientation.Left)!)
                    this.Cells.set(Orientation.Up, this.Cells.get(Orientation.None)!)
                    this.Cells.set(Orientation.UpRight, this.Cells.get(Orientation.Right)!)
                    this.Cells.set(Orientation.Left, this.Cells.get(Orientation.DownLeft)!)
                    this.Cells.set(Orientation.None, this.Cells.get(Orientation.Down)!)
                    this.Cells.set(Orientation.Right, this.Cells.get(Orientation.DownRight)!)
                    // assign new cells to load
                    cellIDs.push(
                        contig.get(Orientation.DownLeft)!,
                        contig.get(Orientation.Down)!,
                        contig.get(Orientation.DownRight)!,
                    )
                    break;
                case Orientation.DownLeft:
                    // shift preloaded
                    this.Cells.set(Orientation.Right, this.Cells.get(Orientation.Down)!)
                    this.Cells.set(Orientation.UpRight, this.Cells.get(Orientation.None)!)
                    this.Cells.set(Orientation.Up, this.Cells.get(Orientation.Left)!)
                    this.Cells.set(Orientation.None, this.Cells.get(Orientation.DownLeft)!)
                    // assign new cells to load
                    cellIDs.push(
                        contig.get(Orientation.DownRight)!,
                        contig.get(Orientation.Down)!,
                        contig.get(Orientation.DownLeft)!,
                        contig.get(Orientation.Left)!,
                        contig.get(Orientation.UpLeft)!,
                    )
                    break;
                case Orientation.Left:
                    // shift preloaded
                    this.Cells.set(Orientation.UpRight, this.Cells.get(Orientation.Up)!)
                    this.Cells.set(Orientation.Up, this.Cells.get(Orientation.UpLeft)!)
                    this.Cells.set(Orientation.Right, this.Cells.get(Orientation.None)!)
                    this.Cells.set(Orientation.None, this.Cells.get(Orientation.Left)!)
                    this.Cells.set(Orientation.DownRight, this.Cells.get(Orientation.Down)!)
                    this.Cells.set(Orientation.Down, this.Cells.get(Orientation.DownLeft)!)
                    // assign new cells to load
                    cellIDs.push(
                        contig.get(Orientation.UpLeft)!,
                        contig.get(Orientation.Left)!,
                        contig.get(Orientation.DownLeft)!,
                    )
                    break;
                case Orientation.UpLeft:
                    // shift preloaded
                    this.Cells.set(Orientation.Right, this.Cells.get(Orientation.Up)!)
                    this.Cells.set(Orientation.DownRight, this.Cells.get(Orientation.None)!)
                    this.Cells.set(Orientation.Down, this.Cells.get(Orientation.Left)!)
                    this.Cells.set(Orientation.None, this.Cells.get(Orientation.UpLeft)!)
                    // assign new cells to load
                    cellIDs.push(
                        contig.get(Orientation.UpRight)!,
                        contig.get(Orientation.Up)!,
                        contig.get(Orientation.UpLeft)!,
                        contig.get(Orientation.Left)!,
                        contig.get(Orientation.DownLeft)!,
                    )
                    break;
            }

            // update border after none cell is up to date
            const cn = this.Cells.get(Orientation.None)!
            this.Border.set(Orientation.Up, cn.Cell.getY() * this.World.getCellheight())
            this.Border.set(Orientation.Right, (cn.Cell.getX() + 1) * this.World.getCellwidth())
            this.Border.set(Orientation.Down, (cn.Cell.getY() + 1) * this.World.getCellheight())
            this.Border.set(Orientation.Left, cn.Cell.getX() * this.World.getCellwidth())

            return this.listCell((() => {
                const req = new CellDTO.ListCellReq()

                req.setIdsList(cellIDs.filter((v) => (!(v == undefined))))

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
            loadedCells.map((v) => {
                if (!contig.has(v)) {
                    // world border
                    console.log('contig has no:', v)
                    return
                }

                const c = cellMap.get(ulid(contig.get(v)!))!

                if (!c) {
                    // world border
                    // TO INVESTIGATE, shouldn't happen ?
                    console.log('cell not found:', v)
                    return
                }

                console.log(v)

                const tm = ulid(c.getTilemap_asU8())
                const ts = ulid(c.getTileset_asU8())

                this.CellLoader.image(ts, 'img/' + ts +'.png')
                this.CellLoader.tilemapTiledJSON(tm, 'json/' + tm +'.json')

                const m = this.make.tilemap()
                this.Cells.set(v, {
                    Cell: c,
                    Tilemap: m,
                    Layer: m.createBlankLayer('blank', ''),
                })

                let loadedTM: boolean, loadedTS: boolean = false
                const load = () => {
                    console.log('load ', v, loadedTM, loadedTS)
                    if (!loadedTM || !loadedTS) {
                        return
                    }

                    // create new cell
                    const map = this.make.tilemap({key: tm, width: this.World.getCellwidth(), height: this.World.getCellheight()})
                    const set = map.addTilesetImage(ts)
                    const layer = map.createLayer('Tile Layer 1', set, c.getX() * this.World.getCellwidth(), c.getY() * this.World.getCellheight())

                    const cc = this.Cells.get(v)
                    if (!cc) {
                        // should never happen
                        return
                    }

                    cc.Tilemap = map
                    cc.Layer = layer

                    // Loading reset
                    // Add safe concurrency for next 2 instructions
                    loaded++
                    console.log('loaded:', v, ':', loaded)
                    if (loaded == loadedCells.length) {
                        this.Loading = undefined
                    }
                }
                this.CellLoader.on('filecomplete-tilemapJSON-' + tm, () => {
                    console.log('json complete on ', v)
                    loadedTM = true
                    load()
                })
                this.CellLoader.on('filecomplete-image-' + ts, () => {
                    console.log('image complete on ', v)
                    loadedTS = true
                    load()
                })
            })

            this.CellLoader.start()
        })
    }

    // Connect
    async connect() {
        // call list entities on current cell
        return this.connectPC((message: EntityDTO.ListEntityResp) => {
            // load all entities
            var entityIDs : Uint8Array[] = []

            message.getEntitiesList().forEach((entry: Entity.E) => {
                const id = ulid(entry.getId_asU8())

                if (this.Entities.has(id)) {
                    // update state only
                    this.Entities.get(id)!.E = entry
                } else {
                    // create associated sprite
                    const sprite = this.add.sprite(entry.getX(), entry.getY(), ulid(entry.getId_asU8()))
                    sprite.setDepth(entitySpriteDepth)

                    this.Entities.set(id, {
                        E: entry,
                        // TODO: adjust x and y with cell position
                        Sprite: sprite
                    })

                    if (id == ulid(this.Entity.E.getId_asU8())) {
                        this.Entity.Sprite = sprite
                        // local player sprite loaded, start camera follow
                        this.cameras.main.startFollow(this.Entity.Sprite)
                    }

                    // load entity animations
                    entityIDs.push(entry.getId_asU8())
                }

            })

            if (entityIDs.length == 0) {
                this.EntityBuffer.push(...entityIDs.map((id: Uint8Array) => {
                    return ulid(id)
                }))

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
                    const animationID = ulid(an.getId_asU8())
                    const sheetID = ulid(an.getSheetid_asU8())

                    // return if already loaded
                    if (this.anims.exists(animationID)) {
                        // Play entity animation
                        this.EntityBuffer.push(ulid(an.getEntityid_asU8()))

                        return
                    }

                    const loadAnim = () => {
                            // Create animation
                            const newAnim = this.anims.create({
                                key: animationID,
                                frames: this.anims.generateFrameNumbers(
                                    sheetID,
                                    {start: an.getStart(), end: an.getEnd()},
                                ),
                                frameRate: an.getRate(),
                                repeat: -1,
                            })
                            if (!newAnim) {
                                console.log('failed to load animation ' + animationID)

                                return
                            }

                            console.log('set animation:'+ulid(an.getEntityid_asU8())+":"+an.getName())

                            // Add animation to mapper
                            this.Animations.set(ulid(an.getEntityid_asU8())+":"+an.getName(), animationID)

                            // Play entity animation
                            this.EntityBuffer.push(ulid(an.getEntityid_asU8()))
                    }

                    if (!this.SpriteSheets.get(sheetID)) {
                        // load sprite sheet
                        this.EntityLoader.spritesheet(sheetID, 'img/' + sheetID + '.png', {
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

    // Update server
    async connectUpdate() {
        // call update entity
        return this.updateEntity()
        .then((client: grpc.Client<Entity.E, Entity.E>) => {
            this.EntityClient = client
        })
    }

    // API PC
    connectPC(callback: (message: EntityDTO.ListEntityResp) => void) {
        let req = this.PC;
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        const prom = new Promise<string>((resolve, reject) => {
            grpc.invoke(API.API.ConnectPC, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onMessage: (message: EntityDTO.ListEntityResp) => {
                    callback(message)
                },
                onEnd: (code: grpc.Code, message: string | undefined, trailers: grpc.Metadata)  => {
                    if (code !== grpc.Code.OK || !message) {
                        reject(message)

                        return
                    }

                    resolve(message)
                }
            });
        })

        return prom
    }

    // API PC
    updateEntity() {
        let req = this.PC;
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        const prom = new Promise<grpc.Client<Entity.E, Entity.E>>((resolve, reject) => {
            grpc.client(API.API.UpdateEntity, {
                // metadata: md,
                host: 'http://localhost:8081',
            });
        })

        return prom
    }

    // API Cell
    listCell(req: CellDTO.ListCellReq) {
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        const prom = new Promise<CellDTO.ListCellResp>((resolve, reject) => {
            grpc.unary(API.API.ListCell, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onEnd: res => {
                    const { status, statusMessage, headers, message, trailers } = res;
                    if (status !== grpc.Code.OK || !message) {
                        reject(res)

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
            grpc.unary(API.API.ListEntity, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onEnd: res => {
                    const { status, statusMessage, headers, message, trailers } = res;
                    if (status !== grpc.Code.OK || !message) {
                        reject(res)

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
            grpc.unary(API.API.ListAnimation, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onEnd: res => {
                    const { status, statusMessage, headers, message, trailers } = res;
                    if (status !== grpc.Code.OK || !message) {
                        reject(res)

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
            grpc.unary(API.API.ListWorld, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onEnd: res => {
                    const { status, statusMessage, headers, message, trailers } = res;
                    if (status !== grpc.Code.OK || !message) {
                        reject(res)

                        return
                    }

                    resolve(message as WorldDTO.ListWorldResp)
                }
            });
        })

        return prom
    }
}
