import {Scene} from "phaser";
import {grpc} from '@improbable-eng/grpc-web';

import * as API from '@cmd/api/grpc/api_pb_service';

import * as Entity from '@pkg/entity/entity_pb';
import * as EntityDTO from '@pkg/entity/dto/entity_pb';

import * as PC from '@pkg/entity/pc_pb';

import * as Cell from '@pkg/room/cell_pb';
import * as CellDTO from '@pkg/room/dto/cell_pb';


export class Game extends Scene {

    PC: PC.PC;
    Entity: Entity.E;

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    init (pc: PC.PC) {
        this.PC =  pc;
    }
    preload() {
        this.cache.addCustom('entity')
        this.cache.addCustom('cell')

        // call current pc cell
        this.listCell([this.PC.getId_asU8()])
        .then((cells: CellDTO.ListCellResp) => {
            // load current pc cell
            if (cells.getCellsList().length != 1) {
                throw new Error('failed to load cell')
            }

            const c = cells.getCellsList()[0]
            this.cache.custom['cell'].add(c.getId_asB64(), c)

            this.load.image(c.getId_asB64(), 'img/' + c.getId_asB64() +'.png')
            this.load.tilemapTiledJSON(c.getId_asB64(), 'json/' + c.getId_asB64() +'.json')

            return c
        })
        .then((c: Cell.Cell) => {
            // call contiguous pc cells
            const contig = c.getContiguousMap() as Array<[number, Uint8Array]>

            return this.listCell(contig.map((val) => {return val[1]}))
        })
        .then((cells: CellDTO.ListCellResp) => {
            // load contiguous pc cells
            cells.getCellsList().map((c: Cell.Cell) => {
                this.cache.custom['cell'].add(c.getId_asB64(), c)

                this.load.image(c.getId_asB64(), 'img/' + c.getId_asB64() +'.png')
                this.load.tilemapTiledJSON(c.getId_asB64(), 'json/' + c.getId_asB64() +'.json')
            })
        })
        .then(() => {
            return this.listEntity([this.PC.getEntityid_asU8()])
        })
        .then((entities: EntityDTO.ListEntityResp) => {
            // load current entity tileset
            if (entities.getEntitiesList().length != 1) {
                throw new Error('failed to load entity')
            }

            const e = entities.getEntitiesList()[0]

            this.cache.custom['entity'].add(e.getId_asB64(), e)
            this.Entity = e

            this.load.image(e.getId_asB64(), 'img/' + e.getId_asB64() +'.png')
            this.load.tilemapTiledJSON(e.getId_asB64(), 'json/' + e.getId_asB64() +'.json')
        })
        .catch((err) => {
            console.log(err)
        })

    }
    create() {
        // Add cell tiles as map
        const map = this.make.tilemap({ key: this.Entity.getCellid_asB64() })
        map.addTilesetImage(this.Entity.getCellid_asB64(), this.Entity.getCellid_asB64())
        const c = this.cache.custom['cell'].get(this.Entity.getCellid_asB64()) as Cell.Cell
        // add contiguous
        const contig = c.getContiguousMap() as Array<[number, Uint8Array]>
        contig.map((val) => {
            map.addTilesetImage(val[1].toString(), val[1].toString())
        })

        // Add entity tile
        const e = this.make.tilemap({ key: this.Entity.getId_asB64() })
        e.addTilesetImage('entity', this.Entity.getId_asB64())

    }
    update() {}

    // API PC
    connectPC(callback: (message: CellDTO.Cell) => void) {
        let req = this.PC;
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        const prom = new Promise<string>((resolve, reject) => {
            grpc.invoke(API.API.ConnectPC, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onMessage: (message: CellDTO.Cell) => {
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

    // API Cell
    listCell(IDs: Uint8Array[]) {
        let req = new CellDTO.ListCellReq();
        req.setIdsList(IDs)
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
    listEntity(IDs: Uint8Array[]) {
        let req = new EntityDTO.ListEntityReq();
        req.setIdsList(IDs)
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
}
