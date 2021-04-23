import {Scene} from 'phaser';
import {grpc} from '@improbable-eng/grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb';

import * as API from '@cmd/api/grpc/api_pb_service';

import * as TwitchDTO from '@pkg/twitch/dto/follow_pb';
import * as Twitch from '@pkg/twitch/follow_pb';

import * as RoomDTO from '@pkg/room/dto/room_pb';
import * as Room from '@pkg/room/room_pb';

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
        this.load.html('load_more', 'html/load_more.html')
        this.load.html('create_room', 'html/create_room.html')
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
    }
    update() {}

    // Follow list methods
    initFollow() {
        // init data
        this.cache.custom['home'].add('follow', new TwitchDTO.ListFollowResp())

        // init html follow list
        const follow = this.add.dom(10, 15).createFromCache('follow').setOrigin(0)
        this.cache.custom['home'].add('follow_html', follow)

        // init html load more follow
        const lm = this.add.dom(10, 5).createFromCache('load_more').setOrigin(0)
        lm.setInteractive()
        lm.addListener('click').on('click', () => {})
        this.cache.custom['home'].add('load_more_follow_html', lm)
    }
    loadFollow(after: string, first: number) {
        this.listFollow(after, first)
        .then((follows: TwitchDTO.ListFollowResp) => {
            // update data
            const fol = this.cache.custom['home'].get('follow') as TwitchDTO.ListFollowResp
            fol.setCursor(follows.getCursor())
            fol.setFollowsList(fol.getFollowsList().concat(follows.getFollowsList()))

            this.displayFollow(fol)
        })
        .catch((err) => {
            console.log(err)
        })
    }
    displayFollow(follows: TwitchDTO.ListFollowResp) {
        // build new follow HTML
        const ol = this.cache.html.get('follow')
        const li = this.cache.html.get('follow_line')
        const lines = follows.getFollowsList().reduce((acc:string, fol: Twitch.Follow) => acc + li.replace('{{login}}', fol.getTologin()), '')

        // update load more button
        const lm = this.cache.custom['home'].get('load_more_follow_html') as Phaser.GameObjects.DOMElement
        lm.removeAllListeners()
        if (follows.getCursor() != '') {
            lm.addListener('click').on('click', () => {
                this.loadFollow(follows.getCursor(), 20)
            })
        }

        // update follows list
        const fHTML = this.cache.custom['home'].get('follow_html') as Phaser.GameObjects.DOMElement
        fHTML.setHTML(ol.replace('{{lines}}', lines))
    }


    // Room list methods
    initRoom() {
        // init data
        this.cache.custom['home'].add('room', new RoomDTO.ListRoomResp())

        // init html room list
        const room = this.add.dom(1000, 15).createFromCache('room').setOrigin(0)
        this.cache.custom['home'].add('room_html', room)

        // init html load more room
        const lm = this.add.dom(1000, 5).createFromCache('load_more').setOrigin(0)
        lm.setInteractive()
        lm.addListener('click').on('click', () => {})
        this.cache.custom['home'].add('load_more_room_html', lm)

        // init html create room
        const cm = this.add.dom(1100, 5).createFromCache('create_room').setOrigin(0)
        cm.setInteractive()
        cm.addListener('click').on('click', () => {})
        this.cache.custom['home'].add('create_room_html', cm)
    }
    loadRoom(size: number, state: Uint8Array) {
        this.listRoom(size, state)
        .then((rooms: RoomDTO.ListRoomResp) => {
            // update data
            const ro = this.cache.custom['home'].get('room') as RoomDTO.ListRoomResp
            ro.setState(rooms.getState())
            ro.setRoomsList(ro.getRoomsList().concat(rooms.getRoomsList()))

            this.displayRoom(ro)
        })
        .catch((err) => {
            console.log(err)
        })
    }
    displayRoom(rooms: RoomDTO.ListRoomResp) {
        // build new room HTML
        const ol = this.cache.html.get('room')
        const li = this.cache.html.get('room_line')
        const lines = rooms.getRoomsList().reduce((acc:string, ro: RoomDTO.Room) => acc + li.replace('{{name}}', ro.getRoom()!.getName()), '')

        // update load more button
        const lm = this.cache.custom['home'].get('load_more_room_html') as Phaser.GameObjects.DOMElement
        lm.removeAllListeners()
        if (rooms.getState() != '') {
            lm.addListener('click').on('click', () => {
                this.loadRoom(20, rooms.getState() as Uint8Array)
            })
        }

        // update rooms list
        const rHTML = this.cache.custom['home'].get('room_html') as Phaser.GameObjects.DOMElement
        rHTML.setHTML(ol.replace('{{lines}}', lines))
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
    listFollow(after: string, first: number) {
        let req = new TwitchDTO.ListFollowReq();
        req.setFirst(first.toString())
        req.setAfter(after)
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
    listRoom(size: number, state: Uint8Array) {
        let req = new RoomDTO.ListRoomReq();
        req.setSize(size)
        req.setState(state)
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
    createRoom(name: string) {
        let req = new Room.R();
        req.setName(name)
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
}
