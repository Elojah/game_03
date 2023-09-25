import { GameObjects, Scene } from 'phaser';
import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import { Mutex, MutexInterface } from 'async-mutex';

import { Buffer } from 'buffer'

import { Map as pbmap } from 'google-protobuf';

import { ulid, parse } from '../lib/ulid'

import { grpc as grpcapi } from 'cmd/api/grpc/api';
import { grpc as grpccore } from 'cmd/core/grpc/core';

// import { ConnectReq, ICECandidate, SDP } from 'pkg/rtc/dto/rtc';
import { dto as dtortc } from 'pkg/rtc/dto/rtc';

import { ability } from 'pkg/ability/ability';
import { dto as abilitydto } from 'pkg/ability/dto/ability';
import { ability as cast } from 'pkg/ability/cast';

import { entity as animation } from 'pkg/entity/animation';
import { dto as animationdto } from 'pkg/entity/dto/animation';

import { entity } from 'pkg/entity/entity';
import { dto as entitydto } from 'pkg/entity/dto/entity';

import { entity as pc } from 'pkg/entity/pc';
import { dto as pcdto } from 'pkg/entity/dto/pc';

import { entity as pcpreferences } from 'pkg/entity/pc_preferences';

import { room as cell } from 'pkg/room/cell';
import { dto as celldto } from 'pkg/room/dto/cell';

import { room as world } from 'pkg/room/world';
import { dto as worlddto } from 'pkg/room/dto/world';

import { Empty } from "google-protobuf/google/protobuf/empty";
import { geometry } from 'pkg/geometry/geometry';

const entitySpriteDepth = 42

enum Action {
  None = 0,
  Up = 1,
  Right = 2,
  Down = 3,
  Left = 4,
  Hotkey00 = 5,
  Hotkey01 = 6,
  Hotkey02 = 7,
  Hotkey03 = 8,
  Hotkey04 = 9,
  Hotkey05 = 10,
  Hotkey10 = 11,
  Hotkey11 = 12,
  Hotkey12 = 13,
  Hotkey13 = 14,
  Hotkey14 = 15,
  Hotkey15 = 16,
};

type GraphicTarget = {
  Type: ability.TargetType
  GroupID: string
  Targets: Map<string, ability.Target>
}

type GraphicEffect = {
  ID: string
  Targets: Array<GraphicTarget>
}

type GraphicAbility = {
  A: ability.A
  // Targets by effect
  // Targets array define order in which target groups are defined
  Effects: Array<GraphicEffect>
}

type BodyEntity = {
  E: entity.E
  Body: Phaser.Types.Physics.Arcade.SpriteWithDynamicBody
  Animations: Map<string, string>
  Orientation: cell.Orientation
  Abilities: Map<string, GraphicAbility>
}

type GraphicEntity = {
  E: entity.E
  Animations: Map<string, string>
  Direction: Phaser.Geom.Point
  Interpolation: number
  Sprite: Phaser.GameObjects.Sprite

  Objects: Map<string, Phaser.Physics.Arcade.StaticGroup>
  Colliders: Map<string, Phaser.Physics.Arcade.Collider>
}

function updateGraphicEntity(ge: GraphicEntity, x: number, y: number) {
  // update direction to last known position
  if (ge.Direction.x != x || ge.Direction.y != y) {
    // TODO update interpolation as speed to catch up latency
    ge.Interpolation = 0
  }

  // TODO Add prediction/reconciliation ?

  ge.Direction.setTo(x, y)
}

function interpolateGraphicEntity(ge: GraphicEntity) {
  if (ge.Interpolation < 1) {
    ge.Interpolation = ge.Interpolation + 0.05
  }
}

type GraphicCell = {
  Cell: cell.Cell
  Tilemap: Phaser.Tilemaps.Tilemap
  Objects: Map<string, Phaser.Physics.Arcade.StaticGroup>
  Colliders: Map<string, Phaser.Physics.Arcade.Collider>
  Layers: Map<string, Phaser.Tilemaps.TilemapLayer | null>
}

export class Game extends Scene {

  // Self
  PC: pc.PC;
  Entity: BodyEntity;
  EntityID: string;

  // Cells & maps
  Cells: Map<cell.Orientation, GraphicCell>
  CellsByID: Map<string, cell.Orientation>
  // Specific loader for cells
  CellLoader: Phaser.Loader.LoaderPlugin
  // Current world loaded
  World: world.World
  // Client cell loading
  Border: Map<cell.Orientation, number>
  // Background image
  Background: Phaser.GameObjects.TileSprite
  // Loading mutex
  Loading: Uint8Array | undefined

  // Entities
  Entities: Map<string, GraphicEntity>
  // Specific loader for entities
  EntityLoader: Phaser.Loader.LoaderPlugin
  // Entity update map to detect disconnect
  EntityLastTick: Map<string, number>
  cleanEntitiesDelay: number

  CurrentAnimationID: string | undefined

  // Abilities for entity
  Abilities: Map<string, ability.A>

  // Keys assignment
  Keys: Map<Action, Phaser.Input.Keyboard.Key>

  // Targets live variable
  Targeting: GraphicTarget | undefined
  CastTargets: Map<string, cast.CastTarget>

  // Spritesheets already loaded
  SpriteSheets: Map<string, integer>

  // Static blank cell for filling
  Blank: GraphicCell

  // Start load pc mutex before map
  LoadMapMutex: Mutex

  SyncTimer: number

  Connected: MutexInterface.Releaser

  SendChannel: RTCDataChannel

  APIClient: grpcapi.APIClient

  CoreClient: grpccore.CoreClient

  constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
    super(config);
  }

  init(p: pcdto.PC) {
    this.APIClient = new grpcapi.APIClient('https://api.legacyfactory.com:8082', null)
    this.CoreClient = new grpccore.CoreClient('https://core.legacyfactory.com:8083', null)

    this.PC = p.PC;
    const e = p.Entity;
    this.Entity = {
      E: e,
      Body: this.physics.add.sprite(0, 0, ''),
      Animations: new Map(),
      Orientation: cell.Orientation.Down,
      Abilities: new Map()
    }
    this.EntityID = ulid(this.Entity.E.ID)

    this.Cells = new Map()
    this.CellsByID = new Map()
    this.CellLoader = new Phaser.Loader.LoaderPlugin(this)
    this.Border = new Map()

    this.Entities = new Map()
    this.EntityLoader = new Phaser.Loader.LoaderPlugin(this)
    this.EntityLastTick = new Map()
    this.cleanEntitiesDelay = 2 * 1000 // 2 seconds without tick (async) triggers entity removal

    this.Abilities = new Map()

    this.Keys = new Map()

    this.SpriteSheets = new Map()

    this.World = new world.World()

    this.LoadMapMutex = new Mutex()

    const m = this.make.tilemap()
    this.Blank = {
      Cell: new cell.Cell,
      Tilemap: m,
      Layers: new Map(),
      Colliders: new Map(),
      Objects: new Map(),
    }
  }

  async createConnection() {
    this.Connected = await this.LoadMapMutex.acquire()

    const local = new RTCPeerConnection({
      iceServers: [{
        urls: [
          'stun:stun.l.google.com:19302',
          'stun:stun1.l.google.com:19302',
          'stun:stun2.l.google.com:19302',
          'stun:stun3.l.google.com:19302',
          'stun:stun4.l.google.com:19302',
          'stun:stun.ekiga.net',
          'stun:stun.ideasip.com',
          'stun:stun.rixtelecom.se',
          'stun:stun.schlund.de',
          'stun:stun.stunprotocol.org:3478',
          'stun:stun.voiparound.com',
          'stun:stun.voipbuster.com',
          'stun:stun.voipstunt.com',
          'stun:stun.voxgratia.org',
        ]
      }]
    });

    this.SendChannel = local.createDataChannel('send_entity')
    this.SendChannel.onopen = () => {
      console.log('channel send_entity opened')
      this.LoadMapMutex.waitForUnlock().then(() => {
        console.log('ticker send_entity start')
        const id = ulid(this.Entity.E.ID)
        this.SyncTimer = window.setInterval(() => {

          const c = new cast.Cast()

          const pos = new geometry.Circle()
          pos.X = this.Entity.E.X
          pos.Y = this.Entity.E.Y

          const target = new cast.CastTarget()
          target.Circle = pos

          c.Targets.set(id, target)
          const now = Date.now()
          c.At = now

          this.SendChannel.send(c.serializeBinary())

          console.log('send entity position at ' + now)
        }, 500)
      })
    }
    this.SendChannel.onclose = () => { console.log('channel send_entity closed') }
    this.SendChannel.onmessage = (m) => { console.log('message received on send_entity:', m) }

    const re = local.createDataChannel('receive_entity')
    re.onopen = () => { console.log('channel receive_entity opened') }
    re.onclose = () => { console.log('channel receive_entity closed') }
    re.onmessage = (m) => {
      console.log('message received on receive_entity')
      const resp = entitydto.ListEntityResp.deserializeBinary(m.data)

      this.displayEntities(resp)
    }

    local.oniceconnectionstatechange = e => console.log('peer connection state change', local.iceConnectionState)

    local.onnegotiationneeded = (e) => local.createOffer()
      .then(d => {
        local.setLocalDescription(d)

        const req = new dtortc.ConnectReq()
        const sdp = new dtortc.SDP()
        sdp.Encoded = Buffer.from(JSON.stringify(d), 'binary').toString('base64')
        req.SDP = sdp
        req.PCID = this.PC.ID
        req.WorldID = this.PC.WorldID

        this.coreConnect(req)
          .then((resp) => {
            const answer = Buffer.from(resp.getEncoded(), 'base64').toString('ascii')
            local.setRemoteDescription(JSON.parse(answer))

            console.log('connect success')
          })
          .then(() => {
            // receive ICE
            this.receiveICE((candidate) => {
              const ic = Buffer.from(candidate.getEncoded(), 'base64').toString('ascii')

              console.log('ice candidate received from signal', ic)

              local.addIceCandidate(JSON.parse(ic))
            })

            // send ICE
            const send = this.sendICE()
            local.onicecandidate = (ic) => {
              if (!ic.candidate) {
                return
              }

              console.log('ice candidate received', ic.candidate)

              const req = new dtortc.ICECandidate()
              req.Encoded = Buffer.from(JSON.stringify(ic.candidate)).toString('base64')
              send.send(req)

              console.log('ice candidate sent to signal', ic.candidate)
            }
          })
          .then(() => {
            console.log('ICE trickle setup')
          })
          .catch((err) => { console.log('failed to connect rtc', err) });
      })
      .catch((err) => { console.log('failed to setup negotiation', err) });
  }

  preload() {
    this.load.html('menu', 'html/menu.html')
    this.load.html('hotkey', 'html/hotkey.html')
    this.load.html('window-entity', 'html/window-entity.html')
    this.load.html('window-ability', 'html/window-ability.html')
    this.load.html('window-ability-line', 'html/window-ability-line.html')
    this.load.image('game_background', 'img/game_background.png')
  }

  create() {
    // Set background & menu
    this.Background = this.add.tileSprite(0, 0, 1024, 512, 'game_background', 'img/game_background.png').setOrigin(0).setScrollFactor(0);
    this.Background.setDisplaySize(this.game.scale.width, this.game.scale.height)

    this.createUI()
    this.applyPCPreferences()

    this.Loading = this.Entity.E.CellID

    // Load world & cell sync
    // Connect for entity send/receive
    this.createConnection()
      .then(() => {
        return this.LoadMapMutex.waitForUnlock()
      })
      .then(() => {
        return this.listWorld((() => {
          const req = new worlddto.ListWorldReq()

          req.IDs = [this.PC.WorldID]
          req.Size = 1

          return req
        })())
      })
      .then((ws: worlddto.ListWorldResp) => {
        if (ws.Worlds.length != 1) {
          console.log('failed to load world')
        } else {
          this.World = ws.Worlds[0]
        }

        return this.loadMap(cell.Orientation.None)
      }).then(() => {
        // post first loading launch
        // this.Cursors = this.input.keyboard!.createCursorKeys();
        this.createKeys()

        setInterval(() => {
          this.cleanEntities()
        }, this.cleanEntitiesDelay)
      })
      .catch((err) => {
        console.log('setup error', err)
      })
  }

  createKeys() {
    this.Keys.set(Action.Up, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.W)!)
    this.Keys.set(Action.Right, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.D)!)
    this.Keys.set(Action.Down, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.S)!)
    this.Keys.set(Action.Left, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.A)!)
    this.Keys.set(Action.Hotkey00, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.ONE)!)
    this.Keys.set(Action.Hotkey01, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.TWO)!)
    this.Keys.set(Action.Hotkey02, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.THREE)!)
    this.Keys.set(Action.Hotkey03, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.FOUR)!)
    this.Keys.set(Action.Hotkey04, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.FIVE)!)
    this.Keys.set(Action.Hotkey05, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.SIX)!)
    this.Keys.set(Action.Hotkey10, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.SEVEN)!)
    this.Keys.set(Action.Hotkey11, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.EIGHT)!)
    this.Keys.set(Action.Hotkey12, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.NINE)!)
    this.Keys.set(Action.Hotkey13, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.ZERO)!)
    this.Keys.set(Action.Hotkey14, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.COMMA)!)
    this.Keys.set(Action.Hotkey15, this.input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.PERIOD)!)

    this.Keys.forEach((v, k) => {
      if (k == Action.Up || k == Action.Right || k == Action.Down || k == Action.Left) {
        // used in updateBodyEntity, no more usage here ftm ?
        return
      }

      v.once('down', () => {
        this.prepareAbility(v, k)
      })
    })
  }

  async prepareAbility(key: Phaser.Input.Keyboard.Key, a: Action) {
    let eid = ''

    switch (a) {
      case Action.Hotkey00:
        eid = 'hotkey-0-0-icon'
        break
      case Action.Hotkey01:
        eid = 'hotkey-0-1-icon'
        break
      case Action.Hotkey02:
        eid = 'hotkey-0-2-icon'
        break
      case Action.Hotkey03:
        eid = 'hotkey-0-3-icon'
        break
      case Action.Hotkey04:
        eid = 'hotkey-0-4-icon'
        break
      case Action.Hotkey05:
        eid = 'hotkey-0-5-icon'
        break
      case Action.Hotkey10:
        eid = 'hotkey-1-0-icon'
        break
      case Action.Hotkey11:
        eid = 'hotkey-1-1-icon'
        break
      case Action.Hotkey12:
        eid = 'hotkey-1-2-icon'
        break
      case Action.Hotkey13:
        eid = 'hotkey-1-3-icon'
        break
      case Action.Hotkey14:
        eid = 'hotkey-1-4-icon'
        break
      case Action.Hotkey15:
        eid = 'hotkey-1-5-icon'
        break
    }

    const abilityIDStr = document.getElementById(eid)?.dataset['abilityID']
    if (!abilityIDStr) {
      // no ability in hotkey
      return key.once('down', () => { this.prepareAbility(key, a) })
    }

    const abilityID = parse(abilityIDStr)

    const c = new cast.Cast()

    c.AbilityID = abilityID

    // overrided, used to bypass parse error
    c.ID = abilityID

    const la = this.Entity.Abilities.get(abilityIDStr)

    if (!la) {
      // ability not initialized locally
      return key.once('down', () => { this.prepareAbility(key, a) })
    }

    // loop for multiple effects
    for (const e of la.Effects) {
      // loop for multiple target groups
      // group by groupID
      // order defined in custom ?
      // Allies - Foes - Circle - Rect
      for (const t of e.Targets) {
        const ok = await this.targetSelect(t, key)
        if (!ok) {
          // target selection failed
          return key.once('down', () => { this.prepareAbility(key, a) })
        }

        // don't launch ability if we have zero targets
        if (this.CastTargets.size == 0) {
          return key.once('down', () => { this.prepareAbility(key, a) })
        }

        t.Targets.forEach((target, key) => {
          if (this.CastTargets.size == 0) {
            return
          }

          // assign randomly and depile current cast targets
          const ct = this.CastTargets.entries().next().value
          this.CastTargets.delete(ct[0])
          c.Targets.set(key, ct[1])
        })
      }
    }

    this.launchAbility(c)

    return key.once('down', () => { this.prepareAbility(key, a) })
  }

  launchAbility(c: cast.Cast) {
    const ab = this.Abilities.get(ulid(c.ID))
    if (!ab) {
      console.log('ability not found', ulid(c.ID))
      return
    }

    const la = this.Entity.Abilities.get(ulid(c.ID))
    if (!la) {
      return
    }

    const now = Date.now()
    c.At = now

    // Play cast animation on entity
    // this.Entity.Body.body.setVelocity(0)
    // this.Entity.E.AnimationID =parse(animationID)
    // const duplicateID = this.Entity.Animations.get(animationID)
    // if (duplicateID) {
    // 	this.Entity?.Body?.play(duplicateID, true)
    // }


    // play animation on targets
    for (const e of la.Effects) {
      for (const t of e.Targets) {
        switch (t.Type) {
          case ability.TargetType.NoneTarget: { break }
          case ability.TargetType.Self: { break }
          case ability.TargetType.Ally: {
            t.Targets.forEach((_, id) => {
              const target = c.Targets.get(id)
              if (!target) {
                return
              }
              const e = this.Entities.get(ulid(target.ID))
              if (!e) {
                return
              }

              this.playCastAnimation(ulid(ab.Animation!), e.E.X, e.E.Y)
            })
            break
          }
          case ability.TargetType.Foe: {
            t.Targets.forEach((_, id) => {
              const target = c.Targets.get(id)
              if (!target) {
                return
              }
              const e = this.Entities.get(ulid(target.ID))
              if (!e) {
                return
              }

              this.playCastAnimation(ulid(ab.Animation!), e.E.X, e.E.Y)
            })
            break
          }
          case ability.TargetType.Rect: {
            t.Targets.forEach((_, id) => {
              const target = c.Targets.get(id)
              if (!target) {
                return
              }

              const x = target.Rect?.X!
              const y = target.Rect?.Y!

              this.playCastAnimation(ulid(ab.Animation!), x, y)
            })
            break
          }
          case ability.TargetType.Circle: {
            t.Targets.forEach((_, id) => {
              const target = c.Targets.get(id)
              if (!target) {
                return
              }

              const x = target.Circle?.X!
              const y = target.Circle?.Y!

              this.playCastAnimation(ulid(ab.Animation!), x, y)
            })
            break
          }
        }
      }
    }

    this.SendChannel.send(c.serializeBinary())

    console.log('send ability cast at ' + now)
  }

  playCastAnimation(animationID: string, x: number, y: number) {
    const animID = this.Entity.Animations.get(animationID)
    if (!animID) {
      return
    }

    const tmpanim = this.add.sprite(x, y, '').setVisible(false)
    tmpanim.play(animID)
    tmpanim.on('animationcomplete', () => {
      tmpanim.destroy()
    })
  }

  async targetSelect(gt: GraphicTarget, key: Phaser.Input.Keyboard.Key): Promise<boolean> {
    // check if already targeting
    if (this.Targeting) {
      return false
    }

    this.Targeting = gt
    this.CastTargets = new Map()

    // change cursor
    // this.input.setDefaultCursor

    switch (gt.Type) {
      case ability.TargetType.NoneTarget: { break }
      case ability.TargetType.Self: {
        const ct = new cast.CastTarget()
        ct.ID = this.EntityID
        this.CastTargets.set(this.EntityID, ct)

        // instantly returns instead of waiting confirmation
        return true
        // break
      }
      case ability.TargetType.Ally:
      case ability.TargetType.Foe: {
        // allies and foes are already on over targeting at creation
        // and enabled using this.Targeting set above
        break
      }
      case ability.TargetType.Rect: {
        let n = 0
        const target = this.Targeting?.Targets.entries().next().value

        this.input.on('pointerdown', (pointer: Phaser.Input.Pointer) => {
          const ct = new cast.CastTarget()
          const rect = new geometry.Rect()
          rect.X = Math.round(pointer.worldX)
          rect.Y = Math.round(pointer.worldY)
          rect.Width = target[1].getWidth()
          rect.Height = target[1].getHeight()
          ct.Rect = rect
          // ct.setCellid()
          this.CastTargets.set((n++).toString(), ct)
          if (n >= this.Targeting?.Targets.size!) {
            key.emit('down')

            return
          }
        })
        break
      }
      case ability.TargetType.Circle: {
        let n = 0
        const target = this.Targeting?.Targets.entries().next().value

        this.input.on('pointerdown', (pointer: Phaser.Input.Pointer) => {
          // add new circle
          const ct = new cast.CastTarget()
          const circle = new geometry.Circle()
          circle.X = Math.round(pointer.worldX)
          circle.Y = Math.round(pointer.worldY)
          circle.Radius = target[1].getRadius()
          ct.Circle = circle
          // ct.setCellid()
          this.CastTargets.set((n++).toString(), ct)
          if (n >= this.Targeting?.Targets.size!) {
            key.emit('down')

            return
          }

        })
        break
      }
    }

    // wait for ability click
    return new Promise<boolean>((accept, reject) => {
      key.once('down', () => {
        this.input.removeListener('pointerdown')
        this.Targeting = undefined
        accept(true)
      })
    })
  }

  createUI() {
    this.add.dom(60, 20).createFromCache('menu').setScrollFactor(0)

    document.getElementById('menu-entity')?.addEventListener('click', (ev) => {
      const w = document.getElementById('window-entity')
      if (w) {
        w.remove()

        return
      }

      const we = this.add.dom(200, 300).createFromCache('window-entity').setScrollFactor(0).setVisible(false)

      document.getElementById('window-entity-name')!.innerHTML += this.Entity.E.Name
      document.getElementById('window-entity-damage')!.innerHTML += this.Entity.E.Stats!.Damage.toString()
      document.getElementById('window-entity-defense')!.innerHTML += this.Entity.E.Stats!.Defense.toString()
      document.getElementById('window-entity-move-speed')!.innerHTML += this.Entity.E.Stats!.MoveSpeed.toString()
      document.getElementById('window-entity-cast-speed')!.innerHTML += this.Entity.E.Stats!.CastSpeed.toString()
      document.getElementById('window-entity-cooldown-reduction')!.innerHTML += this.Entity.E.Stats!.CooldownReduction.toString()
      document.getElementById('window-entity-max-hp')!.innerHTML += this.Entity.E.Stats!.MaxHP.toString()
      document.getElementById('window-entity-max-mp')!.innerHTML += this.Entity.E.Stats!.MaxMP.toString()

      we.setVisible(true)
    })

    document.getElementById('menu-ability')?.addEventListener('click', (ev) => {
      const w = document.getElementById('window-ability')
      if (w) {
        w.remove()

        return
      }

      const wa = this.add.dom(200, 300).createFromCache('window-ability').setScrollFactor(0).setVisible(false)

      const table = document.getElementById('window-ability-table')!
      const line = this.cache.html.get('window-ability-line') as string
      this.Abilities.forEach((ab, id) => {
        const iconID = ulid(ab.Icon)

        const tmp = document.createElement('template')
        tmp.innerHTML = line.replace('{{icon}}', iconID).replace('{{name}}', ab.Name).replace('{{id}}', id)
        const li = tmp.content.firstChild as Node
        table.appendChild(li)

        const icon = document.getElementById('window-ability-line-' + id)!
        icon.draggable = true
        icon.ondragstart = (e) => {
          const target = e.target as HTMLElement

          e.dataTransfer?.setData('id', target.id.split('-').at(3)!)
          const iconID = target.getAttribute('src')?.split('/').at(-1)?.split('.').at(0)
          if (iconID) {
            e.dataTransfer?.setData('icon', iconID)
          }
        }
      })

      wa.setVisible(true)
    })

    document.getElementById('menu-settings')?.addEventListener('click', (ev) => {

    })

    // edit hotkeys to drop abilities into
    this.add.dom(1000, 500).createFromCache('hotkey').setScrollFactor(0)

    const keys = ['0', '1', '2', '3', '4', '5']
    keys.map((n, i) => {
      const elem = document.getElementById('hotkey-0-' + n)!
      elem.ondragover = (e) => {
        e.preventDefault()
      }
      elem.ondrop = (e) => {
        e.preventDefault()

        const elem = e.target as HTMLElement

        const icon = document.createElement('img')

        const id = e.dataTransfer?.getData('id')
        const iconID = e.dataTransfer?.getData('icon')
        icon.src = 'img/assets/' + iconID + '.png'
        icon.id = elem.id + '-icon'
        icon.dataset['abilityID'] = id
        icon.draggable = true
        icon.ondragstart = (e) => {
          const target = e.target as HTMLElement

          e.dataTransfer?.setData('id', target.dataset['abilityID']!)
          const iconID = target.getAttribute('src')?.split('/').at(-1)?.split('.').at(0)
          if (iconID) {
            e.dataTransfer?.setData('icon', iconID)
          }
        }
        icon.ondragend = (e) => {
          const target = e.target as HTMLElement

          target.remove()
        }

        elem.appendChild(icon)
      }
    })
  }

  applyPCPreferences() {
    this.getPCPreferences(this.PC).
      then((result: pcpreferences.PCPreferences) => {
        result.AbilityHotbars.forEach((abilityID: Uint8Array, hotbar: string) => {
          const target = document.getElementById(hotbar)
          if (!target) {
            return
          }

          const icon = document.createElement('img')

          const id = ulid(abilityID)
          const iconID = this.Abilities.get(id)?.Icon
          icon.src = 'img/assets/' + iconID + '.png'
          icon.id = target.id + '-icon'
          icon.dataset['abilityID'] = id
          icon.draggable = true
          icon.ondragstart = (e) => {
            const target = e.target as HTMLElement

            e.dataTransfer?.setData('id', target.dataset['abilityID']!)
            const iconID = target.getAttribute('src')?.split('/').at(-1)?.split('.').at(0)
            if (iconID) {
              e.dataTransfer?.setData('icon', iconID)
            }
          }
          icon.ondragend = (e) => {
            const target = e.target as HTMLElement

            target.remove()
          }

          target.appendChild(icon)
        })
      }).
      catch((error) => {
        console.log(error)
      })
  }

  updateBodyEntity() {
    // Controls + local anim update
    const speed: number = 200
    let animationID = null
    this.Entity.Body.body.setVelocity(0)

    // Move entity
    if (this.Keys.get(Action.Up)?.isDown) {
      animationID = this.Entity.Animations.get('walk_up')
      this.Entity.Body.body.setVelocityY(-speed)
      this.Entity.Orientation = cell.Orientation.Up
    } else if (this.Keys.get(Action.Down)?.isDown) {
      animationID = this.Entity.Animations.get('walk_down')
      this.Entity.Body.body.setVelocityY(speed)
      this.Entity.Orientation = cell.Orientation.Down
    }

    if (this.Keys.get(Action.Right)?.isDown) {
      animationID = this.Entity.Animations.get('walk_right')
      this.Entity.Body.body.setVelocityX(speed)
      this.Entity.Orientation = cell.Orientation.Right
    } else if (this.Keys.get(Action.Left)?.isDown) {
      animationID = this.Entity.Animations.get('walk_left')
      this.Entity.Body.body.setVelocityX(-speed)
      this.Entity.Orientation = cell.Orientation.Left
    }

    if (!animationID) {
      switch (this.Entity.Orientation) {
        case cell.Orientation.Up:
          animationID = this.Entity.Animations.get('idle_up')
          break
        case cell.Orientation.Right:
          animationID = this.Entity.Animations.get('idle_right')
          break
        case cell.Orientation.Down:
          animationID = this.Entity.Animations.get('idle_down')
          break
        case cell.Orientation.Left:
          animationID = this.Entity.Animations.get('idle_left')
          break
      }
    }

    if (animationID) {
      this.Entity.E.AnimationID = parse(animationID)

      const duplicateID = this.Entity.Animations.get(animationID)
      if (duplicateID) {
        this.Entity?.Body?.play(duplicateID, true)
      }
    }
  }

  updateBackground() {
    if (this.Background) {
      this.Background.tilePositionX = this.Entity.Body.body.x % 1024
      this.Background.tilePositionY = this.Entity.Body.body.y % 512
    }
  }

  update(time: number, deltaTime: number) {
    // update cursor position when moving keyboard
    this.input.activePointer.updateWorldPoint(this.cameras.main)

    this.updateBodyEntity()
    this.updateBackground()

    let o: cell.Orientation = cell.Orientation.None
    const x = this.Entity.Body.body.x
    const y = this.Entity.Body.body.y
    const up = this.Border.get(cell.Orientation.Up)!
    const right = this.Border.get(cell.Orientation.Right)!
    const down = this.Border.get(cell.Orientation.Down)!
    const left = this.Border.get(cell.Orientation.Left)!

    if (y < up) {
      if (x < left) {
        o = cell.Orientation.UpLeft
      } else if (x < right) {
        o = cell.Orientation.Up
      } else {
        o = cell.Orientation.UpRight
      }
    } else if (y < down) {
      if (x < left) {
        o = cell.Orientation.Left
      } else if (x < right) {
        o = cell.Orientation.None
      } else {
        o = cell.Orientation.Right
      }
    } else {
      if (x < left) {
        o = cell.Orientation.DownLeft
      } else if (x < right) {
        o = cell.Orientation.Down
      } else {
        o = cell.Orientation.DownRight
      }
    }

    if (o != cell.Orientation.None && this.Loading == undefined) {
      const id = this.Cells.get(o)?.Cell.ID
      this.Loading = id


      if (id) {
        this.Entity.E.CellID = id

        this.loadMap(o).then(() => {
          console.log('loaded cell')
        })
      } else {
        // don't load cell, out of world
      }
    }

    // Move entity
    this.Entity.E.X = Math.round(x)
    this.Entity.E.Y = Math.round(y)

    this.Entities.forEach((e: GraphicEntity) => {
      // self reconciliation
      if (ulid(e.E.ID) == ulid(this.Entity.E.ID)) {
        return
      }

      interpolateGraphicEntity(e)

      const x = e.Sprite.x + ((e.Direction.x - e.Sprite.x) * e.Interpolation)
      const y = e.Sprite.y + ((e.Direction.y - e.Sprite.y) * e.Interpolation)
      e?.Sprite.setX(x)
      e?.Sprite.setY(y)

      const animationID = e.Animations.get(ulid(e.E.AnimationID))
      if (animationID) {
        e?.Sprite.play(animationID, true)
      }
    })
  }

  // Orientation o uses preload surrounded cells
  // cell.Orientation.None loads everything
  async loadMap(o: cell.Orientation) {
    // call current pc cell
    return this.listCell((() => {
      const req = new celldto.ListCellReq()

      req.IDs = [this.Entity.E.CellID]

      return req
    })())
      .then((cells: celldto.ListCellResp) => {
        // load current pc cell
        if (cells.Cells.length != 1) {
          throw new Error('failed to load cell')
        }

        const c = cells.Cells[0]
        const contig = c.Contiguous

        let newCellIDs: Uint8Array[] = []
        let deletedCells: GraphicCell[] = []

        switch (o) {
          case cell.Orientation.None:
            // assign new cells to load
            newCellIDs.push(
              contig.get(cell.Orientation.Up)!,
              contig.get(cell.Orientation.UpRight)!,
              contig.get(cell.Orientation.Right)!,
              contig.get(cell.Orientation.DownRight)!,
              contig.get(cell.Orientation.Down)!,
              contig.get(cell.Orientation.DownLeft)!,
              contig.get(cell.Orientation.Left)!,
              contig.get(cell.Orientation.UpLeft)!,
            )

            // assign cells to delete
            // ...
            // create new blank graphic cell from loaded cell
            this.Cells.set(cell.Orientation.None, {
              Cell: c,
              Tilemap: this.Blank.Tilemap,
              Layers: this.Blank.Layers,
              Colliders: this.Blank.Colliders,
              Objects: this.Blank.Objects
            })

            this.Cells.delete(cell.Orientation.Up)
            this.Cells.delete(cell.Orientation.UpRight)
            this.Cells.delete(cell.Orientation.Right)
            this.Cells.delete(cell.Orientation.DownRight)
            this.Cells.delete(cell.Orientation.Down)
            this.Cells.delete(cell.Orientation.DownLeft)
            this.Cells.delete(cell.Orientation.Left)
            this.Cells.delete(cell.Orientation.UpLeft)
            break;
          case cell.Orientation.Up:
            // assign new cells to load
            newCellIDs.push(
              contig.get(cell.Orientation.UpLeft)!,
              contig.get(cell.Orientation.Up)!,
              contig.get(cell.Orientation.UpRight)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(cell.Orientation.DownLeft)!,
              this.Cells.get(cell.Orientation.Down)!,
              this.Cells.get(cell.Orientation.DownRight)!,
            )
            // shift preloaded
            this.Cells.set(cell.Orientation.DownLeft, this.Cells.get(cell.Orientation.Left)!)
            this.Cells.set(cell.Orientation.Down, this.Cells.get(cell.Orientation.None)!)
            this.Cells.set(cell.Orientation.DownRight, this.Cells.get(cell.Orientation.Right)!)
            this.Cells.set(cell.Orientation.Left, this.Cells.get(cell.Orientation.UpLeft)!)
            this.Cells.set(cell.Orientation.None, this.Cells.get(cell.Orientation.Up)!)
            this.Cells.set(cell.Orientation.Right, this.Cells.get(cell.Orientation.UpRight)!)
            this.Cells.delete(cell.Orientation.UpLeft)
            this.Cells.delete(cell.Orientation.Up)
            this.Cells.delete(cell.Orientation.UpRight)
            break;
          case cell.Orientation.UpRight:
            // assign new cells to load
            newCellIDs.push(
              contig.get(cell.Orientation.UpLeft)!,
              contig.get(cell.Orientation.Up)!,
              contig.get(cell.Orientation.UpRight)!,
              contig.get(cell.Orientation.Right)!,
              contig.get(cell.Orientation.DownRight)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(cell.Orientation.UpLeft)!,
              this.Cells.get(cell.Orientation.Left)!,
              this.Cells.get(cell.Orientation.DownLeft)!,
              this.Cells.get(cell.Orientation.Down)!,
              this.Cells.get(cell.Orientation.DownRight)!,
            )
            // shift preloaded
            this.Cells.set(cell.Orientation.Left, this.Cells.get(cell.Orientation.Up)!)
            this.Cells.set(cell.Orientation.DownLeft, this.Cells.get(cell.Orientation.None)!)
            this.Cells.set(cell.Orientation.Down, this.Cells.get(cell.Orientation.Right)!)
            this.Cells.set(cell.Orientation.None, this.Cells.get(cell.Orientation.UpRight)!)
            this.Cells.delete(cell.Orientation.UpLeft)
            this.Cells.delete(cell.Orientation.Up)
            this.Cells.delete(cell.Orientation.UpRight)
            this.Cells.delete(cell.Orientation.Right)
            this.Cells.delete(cell.Orientation.DownRight)
            break;
          case cell.Orientation.Right:
            // assign new cells to load
            newCellIDs.push(
              contig.get(cell.Orientation.UpRight)!,
              contig.get(cell.Orientation.Right)!,
              contig.get(cell.Orientation.DownRight)!
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(cell.Orientation.UpLeft)!,
              this.Cells.get(cell.Orientation.Left)!,
              this.Cells.get(cell.Orientation.DownLeft)!,
            )
            // shift preloaded
            this.Cells.set(cell.Orientation.UpLeft, this.Cells.get(cell.Orientation.Up)!)
            this.Cells.set(cell.Orientation.Left, this.Cells.get(cell.Orientation.None)!)
            this.Cells.set(cell.Orientation.DownLeft, this.Cells.get(cell.Orientation.Down)!)
            this.Cells.set(cell.Orientation.Up, this.Cells.get(cell.Orientation.UpRight)!)
            this.Cells.set(cell.Orientation.None, this.Cells.get(cell.Orientation.Right)!)
            this.Cells.set(cell.Orientation.Down, this.Cells.get(cell.Orientation.DownRight)!)
            this.Cells.delete(cell.Orientation.UpRight)
            this.Cells.delete(cell.Orientation.Right)
            this.Cells.delete(cell.Orientation.DownRight)
            break;
          case cell.Orientation.DownRight:
            // assign new cells to load
            newCellIDs.push(
              contig.get(cell.Orientation.DownLeft)!,
              contig.get(cell.Orientation.Down)!,
              contig.get(cell.Orientation.DownRight)!,
              contig.get(cell.Orientation.Right)!,
              contig.get(cell.Orientation.UpRight)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(cell.Orientation.UpLeft)!,
              this.Cells.get(cell.Orientation.Left)!,
              this.Cells.get(cell.Orientation.DownLeft)!,
              this.Cells.get(cell.Orientation.Up)!,
              this.Cells.get(cell.Orientation.UpRight)!,
            )
            // shift preloaded
            this.Cells.set(cell.Orientation.Left, this.Cells.get(cell.Orientation.Down)!)
            this.Cells.set(cell.Orientation.UpLeft, this.Cells.get(cell.Orientation.None)!)
            this.Cells.set(cell.Orientation.Up, this.Cells.get(cell.Orientation.Right)!)
            this.Cells.set(cell.Orientation.None, this.Cells.get(cell.Orientation.DownRight)!)
            this.Cells.delete(cell.Orientation.DownLeft)
            this.Cells.delete(cell.Orientation.Down)
            this.Cells.delete(cell.Orientation.DownRight)
            this.Cells.delete(cell.Orientation.Right)
            this.Cells.delete(cell.Orientation.UpRight)
            break;
          case cell.Orientation.Down:
            // assign new cells to load
            newCellIDs.push(
              contig.get(cell.Orientation.DownLeft)!,
              contig.get(cell.Orientation.Down)!,
              contig.get(cell.Orientation.DownRight)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(cell.Orientation.UpLeft)!,
              this.Cells.get(cell.Orientation.Up)!,
              this.Cells.get(cell.Orientation.UpRight)!,
            )
            // shift preloaded
            this.Cells.set(cell.Orientation.UpLeft, this.Cells.get(cell.Orientation.Left)!)
            this.Cells.set(cell.Orientation.Up, this.Cells.get(cell.Orientation.None)!)
            this.Cells.set(cell.Orientation.UpRight, this.Cells.get(cell.Orientation.Right)!)
            this.Cells.set(cell.Orientation.Left, this.Cells.get(cell.Orientation.DownLeft)!)
            this.Cells.set(cell.Orientation.None, this.Cells.get(cell.Orientation.Down)!)
            this.Cells.set(cell.Orientation.Right, this.Cells.get(cell.Orientation.DownRight)!)
            this.Cells.delete(cell.Orientation.DownLeft)
            this.Cells.delete(cell.Orientation.Down)
            this.Cells.delete(cell.Orientation.DownRight)
            break;
          case cell.Orientation.DownLeft:
            // assign new cells to load
            newCellIDs.push(
              contig.get(cell.Orientation.DownRight)!,
              contig.get(cell.Orientation.Down)!,
              contig.get(cell.Orientation.DownLeft)!,
              contig.get(cell.Orientation.Left)!,
              contig.get(cell.Orientation.UpLeft)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(cell.Orientation.UpLeft)!,
              this.Cells.get(cell.Orientation.Up)!,
              this.Cells.get(cell.Orientation.UpRight)!,
              this.Cells.get(cell.Orientation.Right)!,
              this.Cells.get(cell.Orientation.DownRight)!,
            )
            // shift preloaded
            this.Cells.set(cell.Orientation.Right, this.Cells.get(cell.Orientation.Down)!)
            this.Cells.set(cell.Orientation.UpRight, this.Cells.get(cell.Orientation.None)!)
            this.Cells.set(cell.Orientation.Up, this.Cells.get(cell.Orientation.Left)!)
            this.Cells.set(cell.Orientation.None, this.Cells.get(cell.Orientation.DownLeft)!)
            this.Cells.delete(cell.Orientation.DownRight)
            this.Cells.delete(cell.Orientation.Down)
            this.Cells.delete(cell.Orientation.DownLeft)
            this.Cells.delete(cell.Orientation.Left)
            this.Cells.delete(cell.Orientation.UpLeft)
            break;
          case cell.Orientation.Left:
            // assign new cells to load
            newCellIDs.push(
              contig.get(cell.Orientation.UpLeft)!,
              contig.get(cell.Orientation.Left)!,
              contig.get(cell.Orientation.DownLeft)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(cell.Orientation.UpRight)!,
              this.Cells.get(cell.Orientation.Right)!,
              this.Cells.get(cell.Orientation.DownRight)!,
            )
            // shift preloaded
            this.Cells.set(cell.Orientation.UpRight, this.Cells.get(cell.Orientation.Up)!)
            this.Cells.set(cell.Orientation.Right, this.Cells.get(cell.Orientation.None)!)
            this.Cells.set(cell.Orientation.DownRight, this.Cells.get(cell.Orientation.Down)!)
            this.Cells.set(cell.Orientation.Up, this.Cells.get(cell.Orientation.UpLeft)!)
            this.Cells.set(cell.Orientation.None, this.Cells.get(cell.Orientation.Left)!)
            this.Cells.set(cell.Orientation.Down, this.Cells.get(cell.Orientation.DownLeft)!)
            this.Cells.delete(cell.Orientation.UpLeft)
            this.Cells.delete(cell.Orientation.Left)
            this.Cells.delete(cell.Orientation.DownLeft)
            break;
          case cell.Orientation.UpLeft:
            // assign new cells to load
            newCellIDs.push(
              contig.get(cell.Orientation.UpRight)!,
              contig.get(cell.Orientation.Up)!,
              contig.get(cell.Orientation.UpLeft)!,
              contig.get(cell.Orientation.Left)!,
              contig.get(cell.Orientation.DownLeft)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(cell.Orientation.UpRight)!,
              this.Cells.get(cell.Orientation.Right)!,
              this.Cells.get(cell.Orientation.DownRight)!,
              this.Cells.get(cell.Orientation.Down)!,
              this.Cells.get(cell.Orientation.DownLeft)!,
            )
            // shift preloaded
            this.Cells.set(cell.Orientation.Right, this.Cells.get(cell.Orientation.Up)!)
            this.Cells.set(cell.Orientation.DownRight, this.Cells.get(cell.Orientation.None)!)
            this.Cells.set(cell.Orientation.Down, this.Cells.get(cell.Orientation.Left)!)
            this.Cells.set(cell.Orientation.None, this.Cells.get(cell.Orientation.UpLeft)!)
            this.Cells.delete(cell.Orientation.UpRight)
            this.Cells.delete(cell.Orientation.Up)
            this.Cells.delete(cell.Orientation.UpLeft)
            this.Cells.delete(cell.Orientation.Left)
            this.Cells.delete(cell.Orientation.DownLeft)
            break;
        }

        // update border after none cell is up to date
        const cn = this.Cells.get(cell.Orientation.None)!
        this.Border.set(cell.Orientation.Up, cn.Cell.Y * this.World.CellHeight)
        this.Border.set(cell.Orientation.Right, (cn.Cell.X + 1) * this.World.CellWidth)
        this.Border.set(cell.Orientation.Down, (cn.Cell.Y + 1) * this.World.CellHeight)
        this.Border.set(cell.Orientation.Left, cn.Cell.X * this.World.CellWidth)

        this.cleanCells(deletedCells).then(() => { console.log('finish to destroy unused tilemaps') })

        // reassign cellsbyid with new cells
        this.CellsByID.clear()
        this.Cells.forEach((v, k) => {
          if (v) {
            this.CellsByID.set(ulid(v.Cell.ID), k)
          }
        })

        return this.listCell((() => {
          const req = new celldto.ListCellReq()

          req.IDs = newCellIDs.filter((v) => (!(v == undefined)))

          return req
        })())
      })
      .then((cells: celldto.ListCellResp) => {

        // map new loaded cells
        const cellMap = new Map<string, cell.Cell>()
        cells.Cells.map((v) => {
          cellMap.set(ulid(v.ID), v)
        })
        if (o == cell.Orientation.None) {
          const c = this.Cells.get(cell.Orientation.None)!
          cellMap.set(ulid(c?.Cell.ID), c.Cell)
        }

        const contig = this.Cells.get(cell.Orientation.None)?.Cell.Contiguous!
        if (o == cell.Orientation.None) {
          contig.set(cell.Orientation.None, this.Cells.get(cell.Orientation.None)?.Cell.ID!)
        }

        // must fit above loaded cells array
        const loadedCells = new Array<cell.Orientation>()
        switch (o) {
          case cell.Orientation.None:
            loadedCells.push(
              cell.Orientation.None,
              cell.Orientation.Up,
              cell.Orientation.UpRight,
              cell.Orientation.Right,
              cell.Orientation.DownRight,
              cell.Orientation.Down,
              cell.Orientation.DownLeft,
              cell.Orientation.Left,
              cell.Orientation.UpLeft
            )
            break;
          case cell.Orientation.Up:
            loadedCells.push(
              cell.Orientation.Up,
              cell.Orientation.UpRight,
              cell.Orientation.UpLeft
            )
            break;
          case cell.Orientation.UpRight:
            loadedCells.push(
              cell.Orientation.Up,
              cell.Orientation.UpRight,
              cell.Orientation.Right,
              cell.Orientation.DownRight,
              cell.Orientation.UpLeft
            )
            break;
          case cell.Orientation.Right:
            loadedCells.push(
              cell.Orientation.UpRight,
              cell.Orientation.Right,
              cell.Orientation.DownRight,
            )
            break;
          case cell.Orientation.DownRight:
            loadedCells.push(
              cell.Orientation.UpRight,
              cell.Orientation.Right,
              cell.Orientation.DownRight,
              cell.Orientation.Down,
              cell.Orientation.DownLeft,
            )
            break;
          case cell.Orientation.Down:
            loadedCells.push(
              cell.Orientation.DownRight,
              cell.Orientation.Down,
              cell.Orientation.DownLeft,
            )
            break;
          case cell.Orientation.DownLeft:
            loadedCells.push(
              cell.Orientation.DownRight,
              cell.Orientation.Down,
              cell.Orientation.DownLeft,
              cell.Orientation.Left,
              cell.Orientation.UpLeft
            )
            break;
          case cell.Orientation.Left:
            loadedCells.push(
              cell.Orientation.DownLeft,
              cell.Orientation.Left,
              cell.Orientation.UpLeft
            )
            break;
          case cell.Orientation.UpLeft:
            loadedCells.push(
              cell.Orientation.Up,
              cell.Orientation.UpRight,
              cell.Orientation.DownLeft,
              cell.Orientation.Left,
              cell.Orientation.UpLeft
            )
            break;
        }

        let loaded = 0
        loadedCells.map((o) => {
          const resetLoad = () => {
            loaded++
            if (loaded == loadedCells.length) {
              console.log('loading unlock')
              this.Loading = undefined
            }
          }

          if (!contig.has(o)) {
            // world border
            console.log('contig has no:', o)
            resetLoad()
            return
          }

          const c = cellMap.get(ulid(contig.get(o)!))!
          if (!c) {
            // world border
            // TO INVESTIGATE, shouldn't happen ?
            console.log('cell not found:', o)
            return
          }

          const tm = ulid(c.Tilemap)
          this.CellLoader.tilemapTiledJSON(tm, 'json/assets/' + tm + '.json')

          this.Cells.set(o, {
            Cell: c,
            Tilemap: this.Blank.Tilemap,
            Layers: new Map(),
            Colliders: new Map(),
            Objects: new Map(),
          })
          this.CellsByID.set(ulid(c.ID), o)

          const loadTM = () => {
            console.log('load_tilemap ', o)

            // create new cell
            const map = this.make.tilemap({ key: tm, width: this.World.CellWidth, height: this.World.CellHeight })

            let sets = new Array<Phaser.Tilemaps.Tileset>()

            const loadTS = (tsName: string) => {
              const ts = map.addTilesetImage(tsName)

              if (!ts) {
                console.log('failed to add tileset image ' + tsName)
                return
              }

              sets.push(ts)

              console.log('loaded TS ', sets.length, '/', map.tilesets.length)
              if (sets.length < map.tilesets.length) {
                return
              }

              const cc = this.Cells.get(o)
              if (!cc) {
                // should never happen
                return
              }

              cc.Tilemap = map

              const x = c.X * this.World.CellWidth
              const y = c.Y * this.World.CellHeight

              map.layers.map((l) => {
                const layer = map.createLayer(l.name, sets, x, y)
                if (!layer) {
                  console.log('failed to create layer:', l.name, x, y)
                  return
                }

                // TODO: required or not ?
                layer.setPipeline('main')

                console.log('created layer:', l.name, x, y)
                cc.Layers.set(l.name, layer)
              })

              // add objects
              map.objects.map((os) => {
                const objects = map.createFromObjects(os.name, {}, false) as Phaser.GameObjects.Sprite[]

                // offset on layer position
                const group = this.physics.add.staticGroup(objects.map((o) => {
                  o.x += x
                  o.y += y

                  return o
                }))
                cc.Objects.set(os.name, group)

                const collider = this.physics.add.collider(this.Entity.Body, group)

                console.log('created object layer:', os.name)
                cc.Colliders.set(os.name, collider)
              })

              // Loading reset
              // Add safe concurrency for next 2 instructions
              resetLoad()
            }

            map.tilesets.map((ts) => {
              if (this.textures.exists(ts.name)) {
                loadTS(ts.name)
              } else {
                this.CellLoader.image(ts.name, 'img/assets/' + ts.name + '.png')
                console.log('add listener on ', 'filecomplete-image-' + ts.name)
                this.CellLoader.on('filecomplete-image-' + ts.name, () => {
                  console.log('image completed on ', o)
                  loadTS(ts.name)
                })
              }
            })
          }

          // Check if files already been loaded
          if (this.cache.tilemap.exists(tm)) {
            loadTM()
          } else {
            console.log('add listener on ', 'filecomplete-tilemapJSON-' + tm)
            this.CellLoader.on('filecomplete-tilemapJSON-' + tm, () => {
              console.log('json completed on ', o)
              loadTM()
            })
          }
        })

        this.CellLoader.on('complete', () => {
          console.log('reset loader')
          this.CellLoader.removeAllListeners()
          this.CellLoader.reset()
        })

        console.log('start global cell loading')
        this.CellLoader.start()
      })
  }

  async cleanCells(cells: GraphicCell[]) {
    cells.map((c) => {
      if (c) {
        console.log('clean cell', ulid(c.Cell.ID))
        c.Colliders.forEach((v, k) => {
          if (v.world) {
            v.destroy()
          }
        })

        c.Objects.forEach((v, k) => {
          if (v.world) {
            v.destroy(true, true)
          }
        })

        c.Tilemap.destroy()
      }
    })
  }

  async cleanEntities() {
    let ids: string[] = new Array()

    this.EntityLastTick.forEach((val, key) => {
      if (val == 0) {
        ids.push(key)
      } else {
        this.EntityLastTick.set(key, 0)
      }
    })

    console.log('cleaning entities', ids)

    ids.forEach((id) => {
      let e = this.Entities.get(id)!

      this.EntityLastTick.delete(id)

      // TODO: ensure it cleans everything correctly ?
      e.Sprite.destroy()

      // remove collisions
      e.Colliders.forEach((v, k) => {
        if (v.world) {
          v.destroy()
        }
      })

      e.Objects.forEach((v, k) => {
        if (v.world) {
          v.destroy(true, true)
        }
      })

      // TODO: clean animations if not used elsewhere ?

      this.Entities.delete(id)
    })
  }

  async displayEntities(message: entitydto.ListEntityResp) {
    // load all entities
    let entityIDs: Uint8Array[] = []

    message.Entities.forEach((entry: entity.E) => {
      const id = ulid(entry.ID)
      const meid = ulid(this.Entity.E.ID)

      // Set last tick to 0
      this.EntityLastTick.set(id, (this.EntityLastTick.get(id) || 0) + 1)

      if (this.Entities.has(id)) {
        // update state only
        const x = entry.X //+ (this.Cells.get(this.CellsByID.get(ulid(entry.CellID))!)?.Cell.X! * this.World.getCellwidth())
        const y = entry.Y //+ (this.Cells.get(this.CellsByID.get(ulid(entry.CellID))!)?.Cell.Y! * this.World.getCellheight())
        updateGraphicEntity(this.Entities.get(id)!, x, y)

        this.Entities.get(id)!.E = entry
      } else {
        // create associated sprite
        // receive own entity from server once and initialize
        if (id == meid) {
          this.Entity.Body.destroy()

          if (entry.Objects.length == 0) {
            console.log('player entity has no collision body. set default')
            this.Entity.Body = this.physics.add.sprite(entry.X, entry.Y, id).setSize(16, 16).setOffset(0, 0)
          } else {
            // pick first dynamic body from list to assign as main collision object
            const obj = entry.Objects.at(0)!
            this.Entity.Body = this.physics.add.sprite(entry.X, entry.Y, id).
              setSize(obj.Width, obj.Height).
              setOffset(obj.X, obj.Y)
          }

          console.log('set body from server info')
          this.Entity.E.X = entry.X
          this.Entity.E.Y = entry.Y

          // local player sprite loaded, start camera follow
          this.cameras.main.startFollow(this.Entity.Body)
          this.Entity.Body.setDepth(entitySpriteDepth)

          this.Entities.set(id, {
            E: entry,
            Animations: new Map(),
            Direction: new Phaser.Geom.Point(entry.X, entry.Y),
            Interpolation: 0.00,
            Sprite: this.Entity.Body,

            Objects: new Map(),
            Colliders: new Map(),
          })

          console.log("receive position from server:", this.Entity.E.X, this.Entity.E.Y)

          // load own abilities
          const req = new abilitydto.ListAbilityReq()
          req.EntityID = this.Entity.E.ID
          req.Size = 100
          this.listAbility(req)
            .then((resp: abilitydto.ListAbilityResp) => {
              resp.Abilities.forEach((ab: ability.A) => {
                const id = ulid(ab.ID)

                const ga: GraphicAbility = {
                  A: ab,
                  Effects: new Array()
                }

                // TODO: group targets in this.Entity.Ability
                // Order to redefine ?
                ab.Effects.forEach((e: ability.Effect, effectID: string) => {
                  const ge: GraphicEffect = {
                    ID: effectID,
                    Targets: new Array()
                  }

                  const targetsByGroup = new Map<string, GraphicTarget>()
                  e.Targets.forEach((target: ability.Target, targetID: string) => {
                    const groupID = ulid(target.GroupID)
                    if (!targetsByGroup.has(groupID)) {
                      targetsByGroup.set(groupID, {
                        Type: target.Type,
                        GroupID: groupID,
                        Targets: new Map()
                      })
                    }

                    targetsByGroup.get(groupID)?.Targets.set(targetID, target)
                  })

                  const targets = new Array<GraphicTarget>()
                  // specify target order here ?
                  targetsByGroup.forEach((v, k) => {
                    targets.push(v)
                  })

                  ga.Effects.push({
                    ID: effectID,
                    Targets: targets
                  })
                })

                this.Abilities.set(id, ab)
                this.Entity.Abilities.set(id, ga)
              })
            })
            .then(() => {
              this.Connected()
            })
            .catch((err) => { console.log(err) })
        } else if (this.Entities.has(meid)) {
          const eid = ulid(entry.ID)
          const sprite = this.add.sprite(entry.X, entry.Y, id)
          sprite.setDepth(entitySpriteDepth - 1)

          // over targeting at sprite creation
          sprite.on('over', () => {
            if (this.Targeting &&
              this.Targeting.Targets.size < this.CastTargets.size &&
              !this.CastTargets.has(eid) &&
              ((this.Targeting.Type == ability.TargetType.Foe && (entry.FactionID != this.Entity.E.FactionID)) ||
                (this.Targeting.Type == ability.TargetType.Ally && (entry.FactionID == this.Entity.E.FactionID)))) {
              const target = new cast.CastTarget()
              target.ID = entry.ID
              this.CastTargets.set(eid, target)
            }
          })

          let ge: GraphicEntity = {
            E: entry,
            Animations: new Map(),
            Direction: new Phaser.Geom.Point(entry.X, entry.Y),
            Interpolation: 0.00, // default speed
            Sprite: sprite,

            Objects: new Map(),
            Colliders: new Map(),
          }

          console.log('set entity: ', id, entry.X, entry.Y)

          // set collision objects
          // offset on layer position
          const objects = entry.Objects
          if (objects.length > 0) {
            const group = this.physics.add.staticGroup(objects.map((b) => {
              console.log('set entity collision: ', id, b.X + ge.E.X, b.Y + ge.E.Y)
              return this.physics.add.
                staticImage(
                  ge.E.X + b.X,
                  ge.E.Y + b.Y,
                  '').
                setSize(b.Width, b.Height).
                setVisible(false).
                setImmovable(true).
                setOffset(0, 0)
            }))

            const collider = this.physics.add.collider(this.Entity.Body, group)

            ge.Objects.set(id, group)
            ge.Colliders.set(id, collider)
            console.log('created entity collision:', id)
          }

          this.Entities.set(id, ge)
        } else {
          // this.Entity.Body not initialized yet, don't load anything until we have it (for collisions)
          return
        }

        // load entity animations
        entityIDs.push(entry.ID)
      }
    })

    if (entityIDs.length == 0) {
      return
    }

    // load all animations
    this.listAnimation((() => {
      const req = new animationdto.ListAnimationReq()

      req.setEntityidsList(entityIDs)
      req.setSize(1000) // TODO: less random

      return req
    })())
      .then((animations: animationdto.ListAnimationResp) => {
        // load all current animations
        animations.getAnimationsList().forEach((an: animation.Animation) => {
          const id = ulid(an.ID)
          const sheetID = ulid(an.SheetID)
          const duplicateID = ulid(an.DuplicateID)
          const entityID = ulid(an.EntityID)

          let anims = this.Entities.get(entityID)?.Animations!

          // loading self animations, change object + add named animations
          if (entityID == ulid(this.Entity.E.ID)) {
            anims = this.Entity.Animations

            console.log('set named animation:', an.Name, id)
            anims.set(an.Name, id)
          }

          // return if already loaded
          if (this.anims.exists(duplicateID)) {
            anims.set(id, duplicateID)

            return
          }

          const loadAnim = (an: animation.Animation) => {
            // Create animation
            const seq = an.Sequence.length > 0 ? { frames: an.Sequence } : { start: an.Start, end: an.End };
            const newAnim = this.anims.create({
              key: duplicateID,
              frames: this.anims.generateFrameNumbers(
                sheetID,
                seq,
              ),
              frameRate: an.Rate,
              repeat: an.Repeat,
              delay: an.Delay,
              duration: an.Duration,
              hideOnComplete: an.ShowAndHide,
              showOnStart: an.ShowAndHide,
            })
            if (!newAnim) {
              console.log('failed to load animation ' + duplicateID)

              return
            }

            console.log('set animation:', id, duplicateID)

            // Add animation to entity animations
            anims.set(id, duplicateID)
          }

          if (!this.SpriteSheets.get(sheetID)) {
            // load sprite sheet

            this.EntityLoader.spritesheet(sheetID, 'img/assets/' + sheetID + '.png', {
              frameWidth: an.FrameWidth,
              frameHeight: an.FrameHeight,
              startFrame: an.FrameStart,
              endFrame: an.FrameEnd,
              margin: an.FrameMargin,
              spacing: an.FrameSpacing,
            }).once('filecomplete-spritesheet-' + sheetID, () => {
              // Add spritesheet to loaded
              this.SpriteSheets.set(sheetID, 1)

              loadAnim(an)
            }).start()
          } else {
            loadAnim(an)
          }
        })
      })
  }

  // API Cell
  listCell(req: celldto.ListCellReq) {
    let md = new grpc.Metadata()
    md.set('token', this.registry.get('token'))

    const prom = new Promise<celldto.ListCellResp>((resolve, reject) => {
      this.APIClient.ListCell(req, md, (res: celldto.ListCellResp) => { resolve(res) })
    })

    return prom
  }

  // API Entity
  listEntity(req: entitydto.ListEntityReq) {
    let md = new grpc.Metadata()
    md.set('token', this.registry.get('token'))

    const prom = new Promise<entitydto.ListEntityResp>((resolve, reject) => {
      this.APIClient.ListEntity(req, md, (res: entitydto.ListEntityResp) => { resolve(res) })
    });

    return prom
  }

  // API Animation
  listAnimation(req: animationdto.ListAnimationReq) {
    let md = new grpc.Metadata()
    md.set('token', this.registry.get('token'))

    const prom = new Promise<animationdto.ListAnimationResp>((resolve, reject) => {
      this.APIClient.ListAnimation(req, md, (res: animationdto.ListAnimationResp) => { resolve(res) })
    });

    return prom
  }

  // API Ability
  listAbility(req: abilitydto.ListAbilityReq) {
    let md = new grpc.Metadata()
    md.set('token', this.registry.get('token'))

    const prom = new Promise<abilitydto.ListAbilityResp>((resolve, reject) => {
      this.APIClient.ListAbility(req, md, (res: abilitydto.ListAbilityResp) => { resolve(res) })
    });

    return prom
  }

  // API PC preferences
  getPCPreferences(req: pc.PC) {
    let md = new grpc.Metadata()
    md.set('token', this.registry.get('token'))

    const prom = new Promise<pcpreferences.PCPreferences>((resolve, reject) => {
      this.APIClient.GetPCPreferences(req, md, (res: pcpreferences.PCPreferences) => { resolve(res) })
    });

    return prom
  }

  updatePCPreferences(req: pcpreferences.PCPreferences) {
    let md = new grpc.Metadata()
    md.set('token', this.registry.get('token'))

    const prom = new Promise<pcpreferences.PCPreferences>((resolve, reject) => {
      this.APIClient.UpdatePCPreferences(req, md, (res: pcpreferences.PCPreferences) => { resolve(res) })
    });

    return prom
  }

  // API World
  listWorld(req: worlddto.ListWorldReq) {
    let md = new grpc.Metadata()
    md.set('token', this.registry.get('token'))

    const prom = new Promise<worlddto.ListWorldResp>((resolve, reject) => {
      this.APIClient.ListWorld(req, md, (res: worlddto.ListWorldResp) => { resolve(res) })
    });

    return prom
  }

  // RTC connection
  sendICE(): grpc.Client<grpc.ProtobufMessage, grpc.ProtobufMessage> {
    let client = grpc.client(Core.SendICE, {
      host: 'https://core.legacyfactory.com:8083',
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

  coreConnect(req: ConnectReq) {
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
}
