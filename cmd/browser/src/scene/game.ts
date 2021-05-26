import {Scene} from "phaser";
import {grpc} from '@improbable-eng/grpc-web';

import * as jspb from "google-protobuf";

import {ulid} from '../lib/ulid'

import * as API from '@cmd/api/grpc/api_pb_service';

import * as Entity from '@pkg/entity/entity_pb';
import * as EntityDTO from '@pkg/entity/dto/entity_pb';

import * as PC from '@pkg/entity/pc_pb';
import * as PCDTO from '@pkg/entity/dto/pc_pb';

import * as Cell from '@pkg/room/cell_pb';
import * as CellDTO from '@pkg/room/dto/cell_pb';

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

export class Game extends Scene {


    PC: PC.PC;
    Entity: Entity.E;
    Cell: jspb.Map<Orientation, Cell.Cell>

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    init (pc: PCDTO.PC) {
        this.PC =  pc.getPc() as PC.PC;
        this.Entity =  pc.getEntity() as Entity.E;
        this.Cell = new jspb.Map<Orientation, Cell.Cell>([])
    }
    preload() {

        // load pc entity
        const tm = ulid(this.Entity.getTilemap_asU8())
        const ts = ulid(this.Entity.getTileset_asU8())
        this.load.image(tm, 'img/' + tm +'.png')
        this.load.tilemapTiledJSON(ts, 'json/' + ts +'.json')

        // call current pc cell
        this.listCell([this.Entity.getCellid_asU8()])
        .then((cells: CellDTO.ListCellResp) => {
            // load current pc cell
            if (cells.getCellsList().length != 1) {
                throw new Error('failed to load cell')
            }

            const c = cells.getCellsList()[0]
            this.Cell.set(Orientation.None, c)

            const tm = ulid(c.getTilemap_asU8())
            const ts = ulid(c.getTileset_asU8())
            this.load.image(tm, 'img/' + tm +'.png')
            this.load.tilemapTiledJSON(ts, 'json/' + ts +'.json')

            return c
        })
        .then((c: Cell.Cell) => {
            // call contiguous pc cells
            const contig = c.getContiguousMap() as jspb.Map<number, Uint8Array>

            var cellIDs : Uint8Array[] = []
            contig.forEach((entry) => {
                cellIDs.push(entry)
            })

            return this.listCell(cellIDs)
        })
        .then((cells: CellDTO.ListCellResp) => {
            // load contiguous pc cells
            const contig = this.Cell.get(Orientation.None)?.getContiguousMap() as jspb.Map<number, Uint8Array>

            // create reverted map to ease access
            const revertContig = new jspb.Map<Uint8Array, number>([])
            contig.forEach((entry, key) => {
                revertContig.set(entry, key)
            })

            cells.getCellsList().map((c: Cell.Cell) => {
                // Pre-assign contiguous cells
                const o = revertContig.get(c.getId_asU8())
                if (!o) {
                    return
                }

                this.Cell.set(o, c)

                const tm = ulid(c.getTilemap_asU8())
                const ts = ulid(c.getTileset_asU8())
                this.load.image(tm, 'img/' + tm +'.png')
                this.load.tilemapTiledJSON(ts, 'json/' + ts +'.json')
            })
        })
        .catch((err) => {
            console.log(err)
        })

    }
    create() {
        // Create tilemap for all cells
        this.Cell.forEach((entry:Cell.Cell) => {
            const tm = ulid(entry.getTilemap_asU8())
            const ts = ulid(entry.getTileset_asU8())
            const map = this.make.tilemap({ key: tm })
            map.addTilesetImage(ts, tm)
        })

        // Add entity tile
        const eid = ulid(this.Entity.getId_asU8())
        const etm = ulid(this.Entity.getTilemap_asU8())
        const ets = ulid(this.Entity.getTileset_asU8())
        const entityMap = this.make.tilemap({ key: eid })
        entityMap.addTilesetImage(ets, etm)

        // Create layer in right order
        // map.createLayer(cellID, cellID)
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
