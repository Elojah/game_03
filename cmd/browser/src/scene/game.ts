import {Scene} from "phaser";
import {grpc} from '@improbable-eng/grpc-web';

import * as API from '@cmd/api/grpc/api_pb_service';

import * as PC from '@pkg/entity/pc_pb';

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

    }
    create() {}
    update() {}

    // API PC
    connectPC(PCID: Uint8Array) {
        let req = this.PC;
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        const prom = new Promise<CellDTO.Cell>((resolve, reject) => {
            grpc.invoke(API.API.ConnectPC, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onMessage: (message: CellDTO.Cell) => {
                    resolve(message)
                },
                onEnd: (code: grpc.Code, message: string | undefined, trailers: grpc.Metadata)  => {
                    if (code !== grpc.Code.OK || !message) {
                        reject(message)

                        return
                    }

                    console.log('stream ended:', code, message)
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
