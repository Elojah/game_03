import {Scene} from 'phaser';
import {grpc} from '@improbable-eng/grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';

import {ulid} from '../lib/ulid'

import * as API from '@cmd/api/grpc/api_pb_service';

import * as TwitchDTO from '@pkg/twitch/dto/follow_pb';

import * as RoomDTO from '@pkg/room/dto/room_pb';
import * as Room from '@pkg/room/room_pb';

import * as PCDTO from '@pkg/entity/dto/pc_pb';
import * as PC from '@pkg/entity/pc_pb';

export class Home extends Scene {

    HTMLlogin: Phaser.GameObjects.DOMElement;

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    preload() {
        this.load.html('follow', 'html/follow.html')
        this.load.html('follow_line', 'html/follow_line.html')
        this.load.html('room', 'html/room.html')
        this.load.html('room_line', 'html/room_line.html')
        this.load.html('pc', 'html/pc.html')
        this.load.html('pc_line', 'html/pc_line.html')
        this.load.image('home_background_01', 'img/home_background_01.png')
    }
    create() {
        if (!document.cookie.startsWith('auth-token=') && !this.registry.get('token')) {
            this.scene.transition({
                target: 'login',
                duration: 1000,
                remove: true,
            })

            return
        }

        if (!this.registry.get('token')) {
            this.registry.set('token', document.cookie.replace('auth-token=', ''))
        }

        this.add.image(0, 0, 'home_background_01').setOrigin(0);

        this.cache.addCustom('home')

        this.initFollow()
        this.loadFollow('', 20);

        this.initRoom()
        this.loadRoom(20, new Uint8Array);

        this.initPC()
    }
    update() {}

    // Follow list methods
    initFollow() {
        // init html follow list
        const follow = this.add.dom(10, 10).createFromCache('follow').setOrigin(0)
        follow.setInteractive()
    }
    loadFollow(after: string, first: number) {
        this.listFollow((() => {
            const req = new TwitchDTO.ListFollowReq()

            req.setAfter(after)
            req.setFirst(first.toString())

            return req
        })())
        .then((follows: TwitchDTO.ListFollowResp) => {
            // update cursor
            const flm = document.getElementById('follow-load-more')

            // remove previous event listener
            const prevLoadMore = this.cache.custom['home'].get('load_more_func')
            if (prevLoadMore) {
                flm?.removeEventListener('click', prevLoadMore)
            }

            // add new event listener
            const loadMore = () => {
                if (follows.getCursor()){
                    this.loadFollow(follows.getCursor(), 20)
                } else {
                    console.log('no more loading')
                }
            }
            flm?.addEventListener('click', loadMore)
            this.cache.custom['home'].add('load_more_func', loadMore)

            // update html
            const followList = document.getElementById('follow-list')
            const line = this.cache.html.get('follow_line') as string
            follows.getFollowsList().map((fol) => {
                const tmp = document.createElement('template')
                tmp.innerHTML = line.replace('{{login}}', fol.getTologin()).replace('{{id}}', fol.getToid())
                const li = tmp.content.firstChild as Node
                followList?.appendChild(li)
            })
        })
        .catch((err) => {
            console.log(err)
        })
    }

    // Room list methods
    initRoom() {
        // init html room list
        const room = this.add.dom(500, 10).createFromCache('room').setOrigin(0)
        room.setInteractive()

        // init html create room button
        const rc = document.getElementById('room-create')
        rc?.addEventListener('click', () => {
            this.createRoom((() => {
                const req = new Room.R()

                req.setName((document.getElementById('room-create-name') as HTMLInputElement).value)

                return req
            })())
            .then((ro: Room.R) => {
                console.log('successfully created room ', ro.getId_asB64())

                // reset room list
                const roomList = document.getElementById('room-list')
                if (roomList) {
                    roomList.innerHTML = ''
                }
                this.loadRoom(20, new Uint8Array)
            })
            .catch((err) => {
                console.log('failed to create room')
            })
        })
    }
    loadRoom(size: number, state: Uint8Array) {
        this.listRoom((() => {
            const req = new RoomDTO.ListRoomReq()

            req.setSize(size)
            req.setState(state)

            return req
        })())
        .then((rooms: RoomDTO.ListRoomResp) => {
            // update cursor
            const rlm = document.getElementById('room-load-more')

            // remove previous event listener
            const prevLoadMore = this.cache.custom['home'].get('room_load_more_func')
            if (prevLoadMore) {
                rlm?.removeEventListener('click', prevLoadMore)
            }

            // add new event listener
            const loadMore = () => {
                if (rooms.getState()){
                    this.loadRoom(20, rooms.getState_asU8())
                } else {
                    console.log('no more loading')
                }
            }
            rlm?.addEventListener('click', loadMore)
            this.cache.custom['home'].add('room_load_more_func', loadMore)

            // update html
            const roomList = document.getElementById('room-list')
            const line = this.cache.html.get('room_line') as string
            rooms.getRoomsList().map((ro) => {
                const tmp = document.createElement('template')
                tmp.innerHTML = line.replace('{{name}}', ro.getRoom()?.getName() as string).replace('{{id}}', ro.getRoom()?.getId_asB64() as string)

                const li = tmp.content.firstChild as Node
                li.addEventListener('click', () => {                    // reset pc list
                    const pcList = document.getElementById('pc-list')
                    if (pcList) {
                        pcList.innerHTML = ''
                    }

                    this.loadPC(ro.getRoom()?.getId_asU8() as Uint8Array, 20, new Uint8Array)
                })
                roomList?.appendChild(li)
            })
        })
        .catch((err) => {
            console.log(err)
        })
    }


    // PC list methods
    initPC() {
        // init html pc list
        const pc = this.add.dom(1000, 10).createFromCache('pc').setOrigin(0)
        pc.setInteractive()
    }
    loadPC(roomID: Uint8Array,size: number, state: Uint8Array) {
        this.listPC((() => {
            const req = new PCDTO.ListPCReq()

            req.setRoomid(roomID)
            req.setSize(size)
            req.setState(state)

            return req
        })())
        .then((pcs: PCDTO.ListPCResp) => {

            // update create pc
            const pcc = document.getElementById('pc-create')

            // remove previous event listener
            const prevCreate = this.cache.custom['home'].get('pc_create_func')
            if (prevCreate) {
                pcc?.removeEventListener('click', prevCreate)
            }

            // add new event listener
            const createPC = () => {
                this.createPC((() => {
                    const req = new PCDTO.CreatePCReq()

                    req.setRoomid(roomID)

                    return req
                })())
                .then((pc: PC.PC) => {
                    console.log('successfully created pc ', pc.getId_asB64())

                    // reset pc list
                    const pcList = document.getElementById('pc-list')
                    if (pcList) {
                        pcList.innerHTML = ''
                    }
                    this.loadPC(roomID, 20, new Uint8Array)
                })
                .catch((err) => {
                    console.log('failed to create pc:' + err)
                })
            }
            pcc?.addEventListener('click', createPC)
            this.cache.custom['home'].add('pc_create_func', createPC)


            // update cursor
            const pclm = document.getElementById('pc-load-more')

            // remove previous event listener
            const prevLoadMore = this.cache.custom['home'].get('pc_load_more_func')
            if (prevLoadMore) {
                pclm?.removeEventListener('click', prevLoadMore)
            }

            // add new event listener
            const loadMore = () => {
                if (pcs.getState()){
                    this.loadPC(roomID, 20, pcs.getState_asU8())
                } else {
                    console.log('no more loading')
                }
            }
            pclm?.addEventListener('click', loadMore)
            this.cache.custom['home'].add('pc_load_more_func', loadMore)

            // update html
            const pcList = document.getElementById('pc-list')
            const line = this.cache.html.get('pc_line') as string
            pcs.getPcsList().map((pc) => {
                const tmp = document.createElement('template')
                const id = ulid(pc.getPc()?.getId_asU8() as Uint8Array)

                tmp.innerHTML = line.replace('{{name}}', id).replace('{{id}}', id)
                const li = tmp.content.firstChild as Node
                li.addEventListener('click', () => {
                    this.cache.custom['home'].destroy()
                    this.cache.html.destroy()
                    this.scene.transition({
                        target: "game",
                        duration: 1000,
                        remove: true,
                        data: pc,
                    })
                })

                pcList?.appendChild(li)
            })
        })
        .catch((err) => {
            console.log(err)
        })
    }


    // API call methods
    ping() {
        const req = new google_protobuf_empty_pb.Empty();
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        grpc.unary(API.API.Ping, {
            metadata: md,
            request: req,
            host: 'http://localhost:8081',
            onEnd: res => {
                const { status, statusMessage, headers, message, trailers } = res;
                if (status !== grpc.Code.OK || !message) {
                    return
                }
            }
        });
    }

    // API Follow
    listFollow(req: TwitchDTO.ListFollowReq) {
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        const prom = new Promise<TwitchDTO.ListFollowResp>((resolve, reject) => {
            grpc.unary(API.API.ListFollow, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onEnd: res => {
                    const { status, statusMessage, headers, message, trailers } = res;
                    if (status !== grpc.Code.OK || !message) {
                        reject(res)

                        return
                    }

                    resolve(message as TwitchDTO.ListFollowResp)
                }
            });
        })

        return prom
    }

    // API Room
    listRoom(req: RoomDTO.ListRoomReq) {
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        const prom = new Promise<RoomDTO.ListRoomResp>((resolve, reject) => {
            grpc.unary(API.API.ListRoom, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onEnd: res => {
                    const { status, statusMessage, headers, message, trailers } = res;
                    if (status !== grpc.Code.OK || !message) {
                        reject(res)

                        return
                    }

                    resolve(message as RoomDTO.ListRoomResp)
                }
            });
        })

        return prom
    }
    createRoom(req: Room.R) {
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        const prom = new Promise<Room.R>((resolve, reject) => {
            grpc.unary(API.API.CreateRoom, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onEnd: res => {
                    const { status, statusMessage, headers, message, trailers } = res;
                    if (status !== grpc.Code.OK || !message) {
                        reject(res)

                        return
                    }

                    resolve(message as Room.R)
                }
            });
        })

        return prom
    }

    // API PC
    listPC(req: PCDTO.ListPCReq) {
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        const prom = new Promise<PCDTO.ListPCResp>((resolve, reject) => {
            grpc.unary(API.API.ListPC, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onEnd: res => {
                    const { status, statusMessage, headers, message, trailers } = res;
                    if (status !== grpc.Code.OK || !message) {
                        reject(res)

                        return
                    }

                    resolve(message as PCDTO.ListPCResp)
                }
            });
        })

        return prom
    }
    createPC(req: PCDTO.CreatePCReq) {
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        const prom = new Promise<PC.PC>((resolve, reject) => {
            grpc.unary(API.API.CreatePC, {
                metadata: md,
                request: req,
                host: 'http://localhost:8081',
                onEnd: res => {
                    const { status, statusMessage, headers, message, trailers } = res;
                    if (status !== grpc.Code.OK || !message) {
                        reject(res)

                        return
                    }

                    resolve(message as PC.PC)
                }
            });
        })

        return prom
    }
}
