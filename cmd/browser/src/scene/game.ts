import {Scene} from "phaser";
import {grpc} from '@improbable-eng/grpc-web';

import * as API from '@cmd/api/grpc/api_pb_service';

import * as PC from '@pkg/entity/pc_pb';

import * as Cell from '@pkg/room/cell_pb';
import * as CellDTO from '@pkg/room/dto/cell_pb';

export class Game extends Scene {

    PC: PC.PC;

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    init (pc: PC.PC) {
        this.PC =  pc;
    }
    preload() {
        this.cache.addCustom('cell')

        // call current pc cell
        this.listCell([this.PC.getId() as Uint8Array])
        .then((cells: CellDTO.ListCellResp) => {
            // load current pc cell
            if (cells.getCellsList().length != 1) {
                throw new Error('failed to load cell')
            }

            const c = cells.getCellsList()[0]
            this.cache.custom['cell'].add(c.getId() as string, c)

            return c
        })
        .then((cell: Cell.Cell) => {
            // call contiguous pc cells
            const ids = cell.getContiguousMap() as Array<[number, Uint8Array]>

            return this.listCell(ids.map((val) => {return val[1]}))
        })
        .then((cells: CellDTO.ListCellResp) => {
            // load contiguous pc cells
            cells.getCellsList().map((c: Cell.Cell) => {
                this.cache.custom['cell'].add(c.getId() as string, c)
            })
        })
        .then(() => {

        })
        .catch((err) => {
            console.log(err)
        })

    }
    create() {}
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
}
