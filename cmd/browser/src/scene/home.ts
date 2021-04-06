import {Scene} from "phaser";
import {grpc} from "@improbable-eng/grpc-web";

import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as API from "@cmd/api/grpc/api_pb_service";
import * as TwitchDTO from "@pkg/twitch/dto/follow_pb";
import { makeDefaultTransport } from "@improbable-eng/grpc-web/dist/typings/transports/Transport";

export class Home extends Scene {

    HTMLlogin: Phaser.GameObjects.DOMElement;

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    preload() {
        this.load.html('login', 'html/login.html')
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
            document.cookie = "auth-token=;expires=Thu, 01 Jan 1970 00:00:00 GMT";
        }

        this.add.image(0, 0, 'home_background_01').setOrigin(0)
        this.listFollow()
    }
    update() {}
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
                    console.log('grpc error: ', status, statusMessage, headers, message, trailers)
                    return
                }
                // Send a validate thing back
            }
        });
    }    
    listFollow() {
        const req = new TwitchDTO.ListFollowReq();
        let md = new grpc.Metadata()
        md.set('token', this.registry.get('token'))

        grpc.unary(API.API.ListFollow, {
            metadata: md,
            request: req,
            host: 'http://localhost:8081',
            onEnd: res => {
                const { status, statusMessage, headers, message, trailers } = res;
                if (status !== grpc.Code.OK || !message) {
                    console.log('grpc error: ', status, statusMessage, headers, message, trailers)
                    return
                }

                console.log('grpc ok: ', status, statusMessage, headers, message, trailers)
                // Send a validate thing back
            }
        });
    }    
}
