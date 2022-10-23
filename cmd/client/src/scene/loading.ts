import { Scene } from "phaser";
import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import * as API from 'cmd/api/grpc/api_pb_service';

import * as dtoroom from 'pkg/room/dto/room_pb';

export class Loading extends Scene {

	HTMLlogin: Phaser.GameObjects.DOMElement;

	constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
		super(config);
	}
	preload() { }
	create() {
		this.add.image(0, 0, 'home_background_00').setOrigin(0)

		// this.getSession()
	}
	update() { }
}
