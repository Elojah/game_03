import {Scene} from "phaser";
import {grpc} from "@improbable-eng/grpc-web";

import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as API from "@cmd/api/grpc/api_pb_service";

export class Home extends Scene {

    HTMLlogin: Phaser.GameObjects.DOMElement;

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    preload() {
        this.load.html('login', 'html/login.html')
        this.load.image('home_background_00', 'img/home_background_00.png')
        this.load.image('home_background_01', 'img/home_background_01.png')
    }
    create() {
        // this.ping();
        
        // console.log(document.cookie.substring("oauth-session=".length))

        if (document.cookie.startsWith("oauth-session=")) {
            this.add.image(0, 0, 'home_background_01').setOrigin(0)
        } else {
            this.add.image(0, 0, 'home_background_00').setOrigin(0)
            this.HTMLlogin = this.add.dom(60, 120).createFromCache('login').setOrigin(0);
        }
    }
    update() {}
    ping() {
        const req = new google_protobuf_empty_pb.Empty();
        grpc.unary(API.API.Ping, {
            request: req,
            host: "https://localhost:8081",
            onEnd: res => {
                const { status, statusMessage, headers, message, trailers } = res;
                if (status !== grpc.Code.OK || !message) {
                    console.log("grpc error: ", status, statusMessage, headers, message, trailers)
                    return
                }
                // Send a validate thing back
            }
        });
    }    
}
