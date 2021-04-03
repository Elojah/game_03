import {Scene} from "phaser";
import {grpc} from "@improbable-eng/grpc-web";

import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as API from "@cmd/api/grpc/api_pb_service";

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
        if (!document.cookie.startsWith("auth-token=")) {
            this.scene.transition({
                target: "login",
                duration: 1000,
                remove: true,
            })            

            return
        }

        this.add.image(0, 0, 'home_background_01').setOrigin(0)
        this.registry.set('token', document.cookie.replace('auth-token=', ''))
        document.cookie = ''
    }
    update() {}
    ping() {
        const req = new google_protobuf_empty_pb.Empty();
        grpc.unary(API.API.Ping, {
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
    // login(token: string | null) {
    //     const req = new google_protobuf_wrappers_pb.StringValue();
    //     req.setValue(token == null ? '' : token.trim())

    //     grpc.unary(API.API.Login, {
    //         request: req,
    //         host: 'http://localhost:8081',
    //         onEnd: res => {
    //             const { status, statusMessage, headers, message, trailers } = res;
    //             if (status !== grpc.Code.OK || !message) {
    //                 console.log('grpc error: ', status, statusMessage, headers, message, trailers)
    //                 return
    //             }
    //             // Send a validate thing back
    //         }
    //     });
    // }    
}
