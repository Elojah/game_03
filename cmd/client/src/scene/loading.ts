import { Scene } from 'phaser';
import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'
import { parse } from '../lib/ulid'
import { Buffer } from 'buffer'

import { API } from 'cmd/api/grpc/api_pb_service';
import { Core } from 'cmd/core/grpc/core_pb_service';
import { ICECandidate, SDP } from 'pkg/rtc/dto/rtc_pb';
import { CreateSessionReq, CreateSessionResp } from 'pkg/user/dto/session_pb';
import { GetPCReq, PC } from 'pkg/entity/dto/pc_pb';

import { Empty } from "google-protobuf/google/protobuf/empty_pb";

export class Loading extends Scene {

	HTMLlogin: Phaser.GameObjects.DOMElement;

	constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
		super(config);
	}

	preload() { }

	create() {
		this.setupRTC()

		return

		this.add.image(0, 0, 'home_background_00').setOrigin(0)

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
				this.cache.html.destroy()
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

	setupRTC() {
		const local = new RTCPeerConnection({
			iceServers: [{
				urls: 'stun:stun.l.google.com:19302'
			}]
		});

		const dc = local.createDataChannel('update_entity');
		dc.onopen = () => { console.log('channel opened') }
		dc.onclose = () => { console.log('channel closed') }
		dc.onmessage = (m) => { console.log('message received:', m) }

		local.oniceconnectionstatechange = e => console.log('peer connection state change', local.iceConnectionState)

		local.onnegotiationneeded = (e) => local.createOffer()
			.then(d => {
				local.setLocalDescription(d)

				const req = new SDP()
				req.setEncoded(Buffer.from(JSON.stringify(d), 'binary').toString('base64'))
				this.connect(req)
					.then((resp) => {
						const answer = Buffer.from(resp.getEncoded(), 'base64').toString('ascii')
						local.setRemoteDescription(JSON.parse(answer))

						console.log('connect success')
					})
					.then(() => {
						this.receiveICE((candidate) => {
							const ic = Buffer.from(candidate.getEncoded(), 'base64').toString('ascii')

							console.log('ice candidate received from signal', ic)

							local.addIceCandidate(JSON.parse(ic))
						})
					})
					.then(() => {
						const send = this.sendICE()
						local.onicecandidate = (ic) => {
							if (!ic.candidate) {
								return
							}

							console.log('ice candidate received', ic.candidate)

							const req = new ICECandidate()
							req.setEncoded(Buffer.from(JSON.stringify(ic.candidate), 'binary').toString('base64'))
							send.send(req)

							console.log('ice candidate sent to signal', ic.candidate)
						}

					})
					.catch((err) => { console.log('failed to connect rtc', err) });
			})
			.catch((err) => { console.log('failed to setup negotiation', err) });
	}

	sendICE(): grpc.Client<grpc.ProtobufMessage, grpc.ProtobufMessage> {
		let client = grpc.client(Core.SendICE, {
			host: 'https://core.legacyfactory.com:8083',
			// transport: grpc.WebsocketTransport(),
		});

		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)
		client.start(md)

		return client
	}

	receiveICE(callback: (message: ICECandidate) => void) {
		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)

		const prom = new Promise<string>((resolve, reject) => {
			grpc.invoke(Core.ReceiveICE, {
				metadata: md,
				request: new Empty(),
				host: 'https://core.legacyfactory.com:8083',
				onMessage: (message: ICECandidate) => {
					callback(message)
				},
				onEnd: (code: grpc.Code, message: string | undefined, trailers: grpc.Metadata) => {
					if (code !== grpc.Code.OK || !message) {
						reject('receive ICE:' + message)

						return
					}

					resolve(message)
				}
			});
		})

		return prom
	}

	connect(req: SDP) {
		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)

		const prom = new Promise<SDP>((resolve, reject) => {
			grpc.unary(Core.Connect, {
				metadata: md,
				request: req,
				host: 'https://core.legacyfactory.com:8083',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as SDP)
				}
			});
		})

		return prom
	}


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
