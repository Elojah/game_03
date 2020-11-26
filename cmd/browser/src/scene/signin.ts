import {Scene} from "phaser";
import {grpc} from "@improbable-eng/grpc-web";

import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as API from "@cmd/api/grpc/api_pb_service";

export class Signin extends Scene {

    Alpha: number;
    Background: Phaser.GameObjects.Image;
    HTMLsignin: Phaser.GameObjects.DOMElement;

    constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
        super(config);
    }
    preload() {
        this.load.html('signin', 'html/signin.html')
        // this.load.image('background_menu', 'background_menu.png')
    }
    create() {
        this.Alpha = 1;
        this.ping();      
        // this.Background = this.add.image(0, 0, 'background_menu').setOrigin(0)

        this.HTMLsignin = this.add.dom(60, 120).createFromCache('signin').setOrigin(0);
        this.HTMLsignin.addListener('click');
        this.HTMLsignin.on('click', this.signin, this)
    }
    update() {}
    signin(event: MouseEvent) {
        if ((<HTMLInputElement>event.target).name !== 'signin') {
            return
        }
        const req = new google_protobuf_empty_pb.Empty();
        // sub.setAlias(alias.value)
        // sub.setEmail(email.value)
        // sub.setPassword(password.value)
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
