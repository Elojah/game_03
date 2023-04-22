import { Scene } from 'phaser';
import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'
import { parse } from '../lib/ulid'
import { Buffer } from 'buffer'

import { API } from 'cmd/api/grpc/api_pb_service';
import { CreateSessionReq, CreateSessionResp, SDP } from 'pkg/user/dto/session_pb';
import { GetPCReq, PC } from 'pkg/entity/dto/pc_pb';

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

		local.oniceconnectionstatechange = e => console.log(local.iceConnectionState)
		local.onicecandidate = event => {
			if (event.candidate === null) {
				console.log(JSON.stringify(local.localDescription))
			}
		}

		local.onnegotiationneeded = e => local.createOffer()
			.then(d => {
				console.log(d)
				local.setLocalDescription(d)
				remote.setRemoteDescription(d)

				const req = new SDP()
				req.setEncoded(Buffer.from(JSON.stringify(d), 'binary').toString('base64'))
				this.rtcConnectPC(req)
					.then(() => {
						console.log('sent rtc sdp')
					})
			})
			.then(() => remote.createAnswer())
			.then((answer) => {
				console.log(answer)
				remote.setLocalDescription(answer)
				local.setRemoteDescription(answer)
			})
			.catch((err) => { console.log('6', err) });

		const remote = new RTCPeerConnection({
			iceServers: [{
				urls: 'stun:stun.l.google.com:19302'
			}]
		});
		remote.ondatachannel = () => { console.log('3') }

		local.onicecandidate = (e) =>
			!e.candidate ||
			remote.addIceCandidate(e.candidate).catch(() => { console.log('4') });
		remote.onicecandidate = (e) =>
			!e.candidate ||
			local.addIceCandidate(e.candidate).catch(() => { console.log('5') });


		// window.copySDP = () => {
		// 	const browserSDP = document.getElementById('localSessionDescription')

		// 	browserSDP.focus()
		// 	browserSDP.select()

		// 	try {
		// 		const successful = document.execCommand('copy')
		// 		const msg = successful ? 'successful' : 'unsuccessful'
		// 		log('Copying SDP was ' + msg)
		// 	} catch (err) {
		// 		log('Unable to copy SDP ' + err)
		// 	}
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


	rtcConnectPC(req: SDP) {
		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)

		const prom = new Promise<null>((resolve, reject) => {
			grpc.unary(API.RTCConnectPC, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(null)
				}
			});
		})

		return prom
	}
}
