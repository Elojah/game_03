import { Scene } from 'phaser';
import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'
import { parse } from '../lib/ulid'

import { API } from 'cmd/api/grpc/api_pb_service';
import { CreateSessionReq, CreateSessionResp } from 'pkg/user/dto/session_pb';
import { GetPCReq, PC } from 'pkg/entity/dto/pc_pb';

export class Loading extends Scene {

	HTMLlogin: Phaser.GameObjects.DOMElement;

	constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
		super(config);
	}

	preload() { }

	create() {
		const urlParams = new URLSearchParams(window.location.search)
		const worldID = parse(urlParams.get('world_id')!)
		const pcID = parse(urlParams.get('pc_id')!)
		let pc = new PC()

		const req = new GetPCReq()
		req.setId(pcID)
		req.setWorldid(worldID)
		this.getPC(req)
			.then((result) => {
				pc = result
			}).then(() => {
				const req = new CreateSessionReq()
				req.setPcid(pcID)
				req.setWorldid(worldID)
				return this.createSession(req)
			})
			.then((result) => {
				this.registry.set('token', result)
				this.scene.transition({
					target: 'game',
					duration: 1000,
					remove: true,
					data: pc,
				})
			})
			.catch((err) => {
				console.log(err)
			})
	}

	update() { }

	getPC(req: GetPCReq) {
		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)

		const prom = new Promise<PC>((resolve, reject) => {
			grpc.unary(API.GetPC, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as PC)
				}
			});
		})

		return prom
	}

	createSession(req: CreateSessionReq) {
		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)

		const prom = new Promise<CreateSessionResp>((resolve, reject) => {
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

					resolve(message as CreateSessionResp)
				}
			});
		})

		return prom
	}
}
