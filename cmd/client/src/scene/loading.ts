import { Scene } from "phaser";
import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'
import { parse } from '../lib/ulid'

import { API } from 'cmd/api/grpc/api_pb_service';
import * as session from 'pkg/user/session_pb';
import * as dtosession from 'pkg/user/dto/session_pb';

export class Loading extends Scene {

	HTMLlogin: Phaser.GameObjects.DOMElement;

	constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
		super(config);
	}

	preload() { }

	create() {
		this.add.image(0, 0, 'home_background_00').setOrigin(0)

		const urlParams = new URLSearchParams(window.location.search)
		const pcID = urlParams.get('pc_id')

		const req = new dtosession.CreateSessionReq()
		req.setPcid(parse(pcID!))
		this.createSession(req)
			.then((result) => {
				this.registry.set('token', result)
			})
			.catch((err) => {
				console.log(err)
			})
	}

	update() { }

	createSession(req: dtosession.CreateSessionReq) {
		let md = new grpc.Metadata()
		md.set('token', getCookie('token')!)

		const prom = new Promise<session.Session>((resolve, reject) => {
			grpc.unary(API.CreateSession, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as session.Session)
				}
			});
		})

		return prom
	}

}
