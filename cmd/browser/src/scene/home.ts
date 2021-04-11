import {Scene} from "phaser";
import {grpc} from "@improbable-eng/grpc-web";

import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as API from "@cmd/api/grpc/api_pb_service";
import * as TwitchDTO from "@pkg/twitch/dto/follow_pb";
import * as Twitch from "@pkg/twitch/follow_pb";
import { makeDefaultTransport } from "@improbable-eng/grpc-web/dist/typings/transports/Transport";

export class Home extends Scene {

    HTMLlogin: Phaser.GameObjects.DOMElement;

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    preload() {
        this.load.html('follow', 'html/follow.html')
        this.load.html('follow_line', 'html/follow_line.html')
        this.load.image('home_background_01', 'img/home_background_01.png')
    }
    create() {
        if (!document.cookie.startsWith("auth-token=") && !this.registry.get('token')) {
            this.scene.transition({
                target: "login",
                duration: 1000,
                remove: true,
            })            

            return
        }

        if (!this.registry.get('token')) {
            this.registry.set('token', document.cookie.replace('auth-token=', ''))
        }

        this.add.image(0, 0, 'home_background_01').setOrigin(0);
        this.displayFollow();
    }
    update() {}
    displayFollow(){        
        this.listFollow()
        .then((follows: TwitchDTO.ListFollowResp)=> {
            if (follows.getCursor() && ) {

            }
            
            const ol = this.cache.html.get('follow')
            const li = this.cache.html.get('follow_line')

            const lines = follows.getFollowsList().reduce((acc:string, fol: Twitch.Follow) => acc + li.replace('{{login}}', fol.getTologin()), '')
            
            this.add.dom(10, 10).createFromHTML(ol.replace('{{lines}}', lines)).setOrigin(0)
        })
        .catch((err)=> {
            console.log(err)
        })

    }
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
    listFollow() {
        let req = new TwitchDTO.ListFollowReq();
        req.setFirst('20') // TODO: paginate correctly
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
}
