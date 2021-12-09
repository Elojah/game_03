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

    Cell: jspb.Map<Orientation, Cell.Cell>
    Cursors: Phaser.Types.Input.Keyboard.CursorKeys
    // Controls: Controls;

    // Others
    Animations: Map<string, string>
    Entities: Map<string, GraphicEntity>

    // Specific loader for entities
    // Default loader this.load reserved for cells
    EntityLoader: Phaser.Loader.LoaderPlugin

    // 2 buffers for entity update
    // Rendering buffer is frozen and used by update
    EntityBuffer: Array<string>
    EntityBufferRendering: Array<string>

    // Grpc client to send entity updates to server
    EntityClient: grpc.Client<Entity.E, Entity.E>

    // Spritesheets already loaded
    SpriteSheets: Map<string, integer>

    // Current world loaded
    World: World.World

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
        this.Cell = new jspb.Map<Orientation, Cell.Cell>([])

        this.Animations = new Map()
        this.Entities = new Map()

        this.EntityLoader = new Phaser.Loader.LoaderPlugin(this)
        this.EntityBuffer = new Array<string>()
        this.EntityBufferRendering = new Array<string>()

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
            return this.loadCell()
        })
        .then(() => {
            this.load.start()
            this.load.on('complete', this.displayMap)
        })
        .catch((err)=>{
            console.log(err)
        })
    }

    update(time: number, deltaTime: number) {
        // Controls + local anim update
        let animationID = null

        if (this.Cursors.up.isDown) {
            this.Entity.E.setY(this.Entity.E.getY() - 1)
            animationID = this.Animations.get(this.EntityID + ':' + 'up')

            if (this.Entity.E.getY() < 0) {
                // TO DO WITH COLLISION
                const id = this.Cell.get(Orientation.Up)?.getId_asU8()
                if (id) {
                    this.Entity.E.setCellid(id)
                    this.Entity.E.setY(this.World.getCellheight() - 1)

                    this.loadCell().then(() => {
                        this.displayMap()
                        console.log('loaded cell')
                    })
                } else {
                    this.Entity.E.setY(0)
                }
            }
        } else if (this.Cursors.right.isDown) {
            this.Entity.E.setX(this.Entity.E.getX() + 1)
            animationID = this.Animations.get(this.EntityID + ':' + 'right')

            if (this.Entity.E.getX() > this.World.getCellwidth()) {
                const id = this.Cell.get(Orientation.Right)?.getId_asU8()
                if (id) {
                    this.Entity.E.setCellid(id)
                    this.Entity.E.setX(0)

                    this.loadCell().then(() => {
                        this.displayMap()
                        console.log('loaded cell')
                    })
                } else {
                    this.Entity.E.setX(this.World.getCellwidth() - 1)
                }
            }
        } else if (this.Cursors.down.isDown) {
            this.Entity.E.setY(this.Entity.E.getY() + 1)
            animationID = this.Animations.get(this.EntityID + ':' + 'down')

            if (this.Entity.E.getY() > this.World.getCellheight()) {
                const id = this.Cell.get(Orientation.Down)?.getId_asU8()
                if (id) {
                    this.Entity.E.setCellid(id)
                    this.Entity.E.setY(0)

                    this.loadCell().then(() => {
                        this.displayMap()
                        console.log('loaded cell')
                    })
                } else {
                    this.Entity.E.setY(this.World.getCellheight() - 1)
                }
            }
        } else if (this.Cursors.left.isDown) {
            this.Entity.E.setX(this.Entity.E.getX() - 1)
            animationID = this.Animations.get(this.EntityID + ':' + 'left')

            if (this.Entity.E.getX() < 0) {
                const id = this.Cell.get(Orientation.Left)?.getId_asU8()
                if (id) {
                    this.Entity.E.setCellid(id)
                    this.Entity.E.setX(this.World.getCellwidth() - 1)

                    this.loadCell().then(() => {
                        this.displayMap()
                        console.log('loaded cell')
                    })
                } else {
                    this.Entity.E.setX(0)
                }
            }
        } else {
            animationID = this.Animations.get(this.EntityID + ':' + 'idle')
        }

        // Move entity locally
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

    displayMap() {
            // Create tilemap for all cells
            this.Cell.forEach((entry:Cell.Cell, key: Orientation) => {
                const ts = ulid(entry.getTileset_asU8())
                const tm = ulid(entry.getTilemap_asU8())

                const map = this.make.tilemap({ key: tm })
                const set = map.addTilesetImage(ts)

                let x, y = 0
                switch (key) {
                    case Orientation.None:
                        // this.cameras.main.setBounds(0, 0, map.widthInPixels, map.heightInPixels);
                        break;
                    case Orientation.Up:
                        y = -map.heightInPixels
                        break;
                    case Orientation.UpRight:
                        x = map.widthInPixels
                        y = -map.heightInPixels
                        break;
                    case Orientation.Right:
                        x = map.widthInPixels
                        break;
                    case Orientation.DownRight:
                        x = map.widthInPixels
                        y = map.heightInPixels
                        break;
                    case Orientation.Down:
                        y = map.heightInPixels
                        break;
                    case Orientation.DownLeft:
                        x = -map.widthInPixels
                        y = map.heightInPixels
                        break;
                    case Orientation.Left:
                        x = -map.widthInPixels
                        break;
                    case Orientation.UpLeft:
                        x = -map.widthInPixels
                        y = -map.heightInPixels
                        break;
                }

                console.log('display layer ' + key.toString())
                const layer = map.createLayer('Tile Layer 1', set, x, y);
            })
    }

    // Loader Cell
    async loadCell() {
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

            // Clear previous cells
            this.Cell.clear()

            const c = cells.getCellsList()[0]
            this.Cell.set(Orientation.None, c)

            const ts = ulid(c.getTileset_asU8())
            const tm = ulid(c.getTilemap_asU8())
            this.load.image(ts, 'img/' + ts +'.png')
            this.load.tilemapTiledJSON(tm, 'json/' + tm +'.json')

            return c
        })
        .then((c: Cell.Cell) => {
            // call contiguous pc cells
            const contig = c.getContiguousMap() as jspb.Map<number, Uint8Array>

            var cellIDs : Uint8Array[] = []
            contig.forEach((entry: Uint8Array) => {
                cellIDs.push(entry)
            })

            return this.listCell((() => {
                const req = new CellDTO.ListCellReq()

                req.setIdsList(cellIDs)

                return req
            })())
        })
        .then((cells: CellDTO.ListCellResp) => {
            // load contiguous pc cells
            const contig = this.Cell.get(Orientation.None)?.getContiguousMap() as jspb.Map<number, Uint8Array>

            // create reverted map to ease access
            const revertContig = new jspb.Map<Uint8Array, number>([])
            contig.forEach((entry: Uint8Array, key: number) => {
                revertContig.set(entry, key)
            })

            cells.getCellsList().map((c: Cell.Cell) => {
                // Pre-assign contiguous cells
                const o = revertContig.get(c.getId_asU8())
                if (!o) {
                    return
                }

                this.Cell.set(o, c)

                const ts = ulid(c.getTileset_asU8())
                const tm = ulid(c.getTilemap_asU8())
                this.load.image(ts, 'img/' + ts +'.png')
                this.load.tilemapTiledJSON(tm, 'json/' + tm +'.json')
            })
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
