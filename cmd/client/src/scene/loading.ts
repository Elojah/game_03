import { Scene } from 'phaser';
import * as grpc from 'grpc-web';

import { getCookie } from 'typescript-cookie'
import { parse } from '../lib/ulid'

import { APIClient } from 'cmd/api/grpc/ApiServiceClientPb';
import { CreateSessionReq, CreateSessionResp } from 'pkg/user/dto/session_pb';
import { GetPCReq, PC } from 'pkg/entity/dto/pc_pb';

export class Loading extends Scene {

  APIClient: APIClient

  Metadata: grpc.Metadata

  HTMLlogin: Phaser.GameObjects.DOMElement;

  constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
    super(config);
  }

  preload() { }

  create() {
    this.Metadata['token'] = getCookie('access')!

    this.APIClient = new APIClient('https://api.legacyfactory.com:8082', null)

    const urlParams = new URLSearchParams(window.location.search)
    const worldID = parse(urlParams.get('world_id')!)
    const pcID = parse(urlParams.get('pc_id')!)
    let pc = new PC()

    const req = new GetPCReq()
    req.setId(pcID)
    req.setWorldid(worldID)
    this.APIClient.getPC(req, this.Metadata)
      .then((result) => {
        pc = result
      }).then(() => {
        const req = new CreateSessionReq()
        req.setPcid(pcID)
        req.setWorldid(worldID)
        return this.APIClient.createSession(req, this.Metadata)
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
}
