import { Scene } from 'phaser';
import * as grpc from 'grpc-web';

import { Mutex, MutexInterface } from 'async-mutex';

import { Buffer } from 'buffer'

import { Map as pbmap } from 'google-protobuf';
import { getCookie } from 'typescript-cookie'

import { ulid, parse } from '../lib/ulid'

import { APIClient } from 'cmd/api/grpc/ApiServiceClientPb';
import { CoreClient } from 'cmd/core/grpc/CoreServiceClientPb';

import { ConnectReq, ICECandidate, SDP } from 'pkg/rtc/dto/rtc_pb';

import * as Ability from 'pkg/ability/ability_pb';
import * as AbilityDTO from 'pkg/ability/dto/ability_pb';
import * as Cast from 'pkg/ability/cast_pb';

import * as Animation from 'pkg/entity/animation_pb';
import * as AnimationDTO from 'pkg/entity/dto/animation_pb';

import * as Entity from 'pkg/entity/entity_pb';
import * as EntityDTO from 'pkg/entity/dto/entity_pb';

import * as PC from 'pkg/entity/pc_pb';
import * as PCDTO from 'pkg/entity/dto/pc_pb';

import * as PCPreferences from 'pkg/entity/pc_preferences_pb';

import * as Cell from 'pkg/room/cell_pb';
import * as CellDTO from 'pkg/room/dto/cell_pb';

import * as World from 'pkg/room/world_pb';
import * as WorldDTO from 'pkg/room/dto/world_pb';

import { Empty } from "pkg/pbtypes/empty_pb";

import { Circle, Rect } from 'pkg/geometry/geometry_pb';

const logger = require('pino')()

const entitySpriteDepth = 42


type GraphicTarget = {
  Type: Ability.TargetType
  GroupID: string
  Targets: Map<string, Ability.Target>
}

type GraphicEffect = {
  ID: string
  Targets: Array<GraphicTarget>
}

type GraphicAbility = {
  A: Ability.A
  // Targets by effect
  // Targets array define order in which target groups are defined
  Effects: Array<GraphicEffect>
}

type BodyEntity = {
  E: Entity.E
  Body: Phaser.Types.Physics.Arcade.SpriteWithDynamicBody
  Animations: Map<string, string>
  Orientation: Cell.Orientation
  Abilities: Map<string, GraphicAbility>
}

type GraphicEntity = {
  E: Entity.E
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
  Cell: Cell.Cell
  Tilemap: Phaser.Tilemaps.Tilemap
  Objects: Map<string, Phaser.Physics.Arcade.StaticGroup>
  Colliders: Map<string, Phaser.Physics.Arcade.Collider>
  Layers: Map<string, Phaser.Tilemaps.TilemapLayer | null>
}

export class Game extends Scene {

  // Self
  PC: PC.PC;
  Entity: BodyEntity;
  EntityID: string;

  // Cells & maps
  Cells: Map<Cell.Orientation, GraphicCell>
  CellsByID: Map<string, Cell.Orientation>
  // Specific loader for cells
  CellLoader: Phaser.Loader.LoaderPlugin
  // Current world loaded
  World: World.World
  // Client cell loading
  Border: Map<Cell.Orientation, number>
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
  Abilities: Map<string, Ability.A>

  // Keys assignment
  Keys: Map<Action, Phaser.Input.Keyboard.Key>

  // Targets live variable
  Targeting: GraphicTarget | undefined
  CastTargets: Map<string, Cast.CastTarget>

  // Spritesheets already loaded
  SpriteSheets: Map<string, integer>

  // Static blank cell for filling
  Blank: GraphicCell

  // Start load pc mutex before map
  LoadMapMutex: Mutex

  MapLoaded: MutexInterface.Releaser

  SyncTimer: number

  SendChannel: RTCDataChannel

  APIClient: APIClient

  CoreClient: CoreClient

  MetadataAccess: grpc.Metadata
  MetadataSession: grpc.Metadata

  constructor(config: string | Phaser.Types.Scenes.SettingsConfig) {
    super(config);
  }

  init(pc: PCDTO.PC) {
    this.MetadataAccess = { 'token': getCookie('access')! }
    this.MetadataSession = { 'token': this.registry.get('token') }

    this.APIClient = new APIClient('https://api.legacyfactory.com:8082', null)
    this.CoreClient = new CoreClient('https://core.legacyfactory.com:8083', null)

    this.PC = pc.getPc() as PC.PC;
    const e = pc.getEntity() as Entity.E;
    this.Entity = {
      E: e,
      Body: this.physics.add.sprite(0, 0, ''),
      Animations: new Map(),
      Orientation: Cell.Orientation.DOWN,
      Abilities: new Map()
    }
    this.EntityID = ulid(this.Entity.E.getId_asU8())

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

    this.World = new World.World()

    this.LoadMapMutex = new Mutex()

    const m = this.make.tilemap()
    this.Blank = {
      Cell: new Cell.Cell,
      Tilemap: m,
      Layers: new Map(),
      Colliders: new Map(),
      Objects: new Map(),
    }
  }

  async createConnection() {
    this.MapLoaded = await this.LoadMapMutex.acquire()

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
      logger.info('channel send_entity opened')
      this.LoadMapMutex.waitForUnlock().then(() => {
        logger.info('ticker send_entity start')
        const id = ulid(this.Entity.E.getId_asU8())
        this.SyncTimer = window.setInterval(() => {

          const c = new Cast.Cast()

          const pos = new Circle()
          pos.setX(this.Entity.E.getX())
          pos.setY(this.Entity.E.getY())

          const target = new Cast.CastTarget()
          target.setCircle(pos)

          c.getTargetsMap().set(id, target)
          const now = Date.now()
          c.setAt(now)

          this.SendChannel.send(c.serializeBinary())

          logger.info('send entity position at ' + now)
        }, 500)
      })
    }
    this.SendChannel.onclose = () => { logger.info('channel send_entity closed') }
    this.SendChannel.onmessage = (m) => { logger.info('message received on send_entity:', m) }

    const re = local.createDataChannel('receive_entity')
    re.binaryType = 'arraybuffer' // !IMPORTANT OR MESSAGES WILL BE EMPTY
    re.onopen = () => { logger.info('channel receive_entity opened') }
    re.onclose = () => { logger.info('channel receive_entity closed') }
    re.onmessage = (m) => {
      logger.info('message received on receive_entity')
      const resp = EntityDTO.ListEntityResp.deserializeBinary(m.data)

      this.displayEntities(resp)
    }

    local.oniceconnectionstatechange = e => logger.info('peer connection state change', local.iceConnectionState)

    local.onnegotiationneeded = (e) => local.createOffer()
      .then(d => {
        local.setLocalDescription(d)

        const req = new ConnectReq()
        const sdp = new SDP()
        sdp.setEncoded(Buffer.from(JSON.stringify(d), 'binary').toString('base64'))
        req.setSdp(sdp)
        req.setPcid(this.PC.getId_asU8())
        req.setWorldid(this.PC.getWorldid_asU8())

        this.CoreClient.connect(req, this.MetadataAccess)
          .then((resp) => {
            const answer = Buffer.from(resp.getEncoded(), 'base64').toString('ascii')
            local.setRemoteDescription(JSON.parse(answer))

            logger.info('connect success')
          })
          .then(() => {
            // receive ICE
            const req = new Empty()
            const stream = this.CoreClient.receiveICE(req, this.MetadataAccess)
            stream.on("data",
              (candidate) => {
                const ic = Buffer.from(candidate.getEncoded(), 'base64').toString('ascii')

                logger.info('ice candidate received from signal', ic)

                local.addIceCandidate(JSON.parse(ic))
              })

            // send ICE
            local.onicecandidate = (ic) => {
              if (!ic.candidate) {
                return
              }

              logger.info('ice candidate received', ic.candidate)

              const req = new ICECandidate()
              req.setEncoded(Buffer.from(JSON.stringify(ic.candidate)).toString('base64'))
              this.CoreClient.sendICE(req, this.MetadataAccess)

              logger.info('ice candidate sent to signal', ic.candidate)
            }
          })
          .then(() => {
            logger.info('ICE trickle setup')
          })
          .catch((err) => { logger.info('failed to connect rtc', err) });
      })
      .catch((err) => { logger.info('failed to setup negotiation', err) });
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

    this.Loading = this.Entity.E.getCellid_asU8()

    // Load world & cell sync
    // Connect for entity send/receive
    this.createConnection()
      .then(() => {
        return this.LoadMapMutex.waitForUnlock()
      })
      .then(() => {
        const req = new WorldDTO.ListWorldReq()

        req.setIdsList([this.PC.getWorldid_asU8()])
        req.setSize(1)
        return this.APIClient.listWorld(req, this.MetadataSession)
      })
      .then((ws: WorldDTO.ListWorldResp) => {
        if (ws.getWorldsList().length != 1) {
          logger.info('failed to load world')
        } else {
          this.World = ws.getWorldsList()[0]
        }

        return this.loadMap(Cell.Orientation.NONE)
      }).then(() => {
        // post first loading launch
        // this.Cursors = this.input.keyboard!.createCursorKeys();
        this.createKeys()
        this.applyPCPreferences()

        setInterval(() => {
          this.cleanEntities()
        }, this.cleanEntitiesDelay)
      })
      .catch((err) => {
        logger.info('setup error', err)
      })
  }

  async prepareAbility(key: Phaser.Input.Keyboard.Key, a: Action) {

    const eid = Game.hotkeyIcons.get(a)

    if (!eid) {
      // should never happen
      return
    }

    const abilityIDStr = document.getElementById(eid)?.dataset['abilityID']
    if (!abilityIDStr) {
      // no ability in hotkey
      return
    }

    // Replace with this.casting
    if (this.Targeting) {
      // cancel targeting
      return
      // this.Targeting = undefined
      // this.CastTargets = new Map()
    }

    const abilityID = parse(abilityIDStr)

    const c = new Cast.Cast()

    c.setAbilityid(abilityID)

    // overrided, used to bypass parse error
    c.setId(abilityID)

    const la = this.Entity.Abilities.get(abilityIDStr)

    if (!la) {
      // ability not initialized locally
      return
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
          return
        }

        // don't launch ability if we have zero targets
        if (this.CastTargets.size == 0) {
          return
        }

        t.Targets.forEach((target, key) => {
          if (this.CastTargets.size == 0) {
            return
          }

          // assign randomly and depile current cast targets
          const ct = this.CastTargets.entries().next().value
          this.CastTargets.delete(ct[0])
          c.getTargetsMap().set(key, ct[1])
        })
      }
    }

    this.launchAbility(c)

    return
  }

  launchAbility(c: Cast.Cast) {
    const abilityID = ulid(c.getAbilityid_asU8())

    logger.info('launch ability ', abilityID)

    const ab = this.Abilities.get(abilityID)
    if (!ab) {
      logger.info('ability not found', abilityID)
      return
    }

    const la = this.Entity.Abilities.get(abilityID)
    if (!la) {
      return
    }

    const now = Date.now()
    c.setAt(now)

    // Play cast animation on entity
    // this.Entity.Body.body.setVelocity(0)
    // this.Entity.E.setAnimationid(parse(animationID))
    // const duplicateID = this.Entity.Animations.get(animationID)
    // if (duplicateID) {
    // 	this.Entity?.Body?.play(duplicateID, true)
    // }


    // play animation on targets
    for (const e of la.Effects) {
      for (const t of e.Targets) {
        switch (t.Type) {
          case Ability.TargetType.NONETARGET: { break }
          case Ability.TargetType.SELF: { break }
          case Ability.TargetType.ALLY: {
            t.Targets.forEach((_, id) => {
              const target = c.getTargetsMap().get(id)
              if (!target) {
                return
              }
              const e = this.Entities.get(ulid(target.getId_asU8()))
              if (!e) {
                return
              }

              this.playCastAnimation(ulid(ab.getAnimation_asU8()!), e.E.getX(), e.E.getY())
            })
            break
          }
          case Ability.TargetType.FOE: {
            t.Targets.forEach((_, id) => {
              const target = c.getTargetsMap().get(id)
              if (!target) {
                return
              }
              const e = this.Entities.get(ulid(target.getId_asU8()))
              if (!e) {
                return
              }

              this.playCastAnimation(ulid(ab.getAnimation_asU8()!), e.E.getX(), e.E.getY())
            })
            break
          }
          case Ability.TargetType.RECT: {
            t.Targets.forEach((_, id) => {
              const target = c.getTargetsMap().get(id)
              if (!target) {
                return
              }

              const x = target.getRect()?.getX()!
              const y = target.getRect()?.getY()!

              this.playCastAnimation(ulid(ab.getAnimation_asU8()!), x, y)
            })
            break
          }
          case Ability.TargetType.CIRCLE: {
            t.Targets.forEach((_, id) => {
              const target = c.getTargetsMap().get(id)
              if (!target) {
                logger.info('target not found', id)
                return
              }

              const x = target.getCircle()?.getX()!
              const y = target.getCircle()?.getY()!

              this.playCastAnimation(ulid(ab.getAnimation_asU8()!), x, y)
            })
            break
          }
        }
      }
    }

    this.SendChannel.send(c.serializeBinary())

    logger.info('send ability cast at ' + now)
  }

  playCastAnimation(animationID: string, x: number, y: number) {
    const animID = this.Entity.Animations.get(animationID)
    if (!animID) {
      logger.info('ability animation not found', animationID)
      return
    }

    const tmpanim = this.add.sprite(x, y, '').setVisible(false)

    logger.info('play ability animation', animationID, x, y)
    tmpanim.play(animID)

    tmpanim.on('animationcomplete', () => {
      tmpanim.destroy()
    })
  }

  async targetSelect(gt: GraphicTarget, key: Phaser.Input.Keyboard.Key): Promise<boolean> {
    // check if already targeting
    // if (this.Targeting) {
    //   return false
    // }

    logger.info('select target on', key.keyCode)

    this.CastTargets = new Map()

    // change cursor
    this.input.manager.canvas.style.cursor = 'crosshair'

    switch (gt.Type) {
      case Ability.TargetType.NONETARGET: { break }
      case Ability.TargetType.SELF: {
        const ct = new Cast.CastTarget()
        ct.setId(this.EntityID)
        this.CastTargets.set(this.EntityID, ct)

        // instantly returns instead of waiting confirmation
        this.input.manager.canvas.style.cursor = 'default'

        return true
      }
      // allies and foes are already on over targeting at creation
      // and enabled using this.Targeting set above
      case Ability.TargetType.ALLY: {
        this.Targeting = gt
        break
      }
      case Ability.TargetType.FOE: {
        this.Targeting = gt
        break
      }
      case Ability.TargetType.RECT: {
        let n = 0
        const target = gt?.Targets.entries().next().value

        this.input.on('pointerdown', (pointer: Phaser.Input.Pointer) => {
          const ct = new Cast.CastTarget()
          const rect = new Rect()
          rect.setX(Math.round(pointer.worldX))
          rect.setY(Math.round(pointer.worldY))
          rect.setWidth(target[1].getWidth())
          rect.setHeight(target[1].getHeight())
          ct.setRect(rect)
          // ct.setCellid()
          this.CastTargets.set((n++).toString(), ct)
          if (n >= gt?.Targets.size!) {
            key.emit('down')

            return
          }
        })
        break
      }
      case Ability.TargetType.CIRCLE: {
        let n = 0
        const target = gt?.Targets.entries().next().value

        this.input.on('pointerdown', (pointer: Phaser.Input.Pointer) => {
          // add new circle
          const ct = new Cast.CastTarget()
          const circle = new Circle()
          circle.setX(Math.round(pointer.worldX))
          circle.setY(Math.round(pointer.worldY))
          circle.setRadius(target[1].getRadius())
          ct.setCircle(circle)
          // ct.setCellid()
          this.CastTargets.set((n++).toString(), ct)
          if (n >= gt?.Targets.size!) {
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
        this.input.manager.canvas.style.cursor = 'default'
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

      document.getElementById('window-entity-name')!.innerHTML += this.Entity.E.getName()
      document.getElementById('window-entity-damage')!.innerHTML += this.Entity.E.getStats()!.getDamage().toString()
      document.getElementById('window-entity-defense')!.innerHTML += this.Entity.E.getStats()!.getDefense().toString()
      document.getElementById('window-entity-move-speed')!.innerHTML += this.Entity.E.getStats()!.getMovespeed().toString()
      document.getElementById('window-entity-cast-speed')!.innerHTML += this.Entity.E.getStats()!.getCastspeed().toString()
      document.getElementById('window-entity-cooldown-reduction')!.innerHTML += this.Entity.E.getStats()!.getCooldownreduction().toString()
      document.getElementById('window-entity-max-hp')!.innerHTML += this.Entity.E.getStats()!.getMaxhp().toString()
      document.getElementById('window-entity-max-mp')!.innerHTML += this.Entity.E.getStats()!.getMaxmp().toString()

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
        const iconID = ulid(ab.getIcon_asU8())

        const tmp = document.createElement('template')
        tmp.innerHTML = line.replace('{{icon}}', iconID).replace('{{name}}', ab.getName()).replace('{{id}}', id)
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
    this.APIClient.getPCPreferences(this.PC, this.MetadataSession).
      then((result: PCPreferences.PCPreferences) => {
        logger.info("pc preferences found")

        result.getAbilityhotbarsMap().forEach((abilityID: Uint8Array, hotkey: string) => {
          logger.info("assign from preferences: ", ulid(abilityID), hotkey)
          const target = document.getElementById(hotkey)
          if (!target) {
            logger.info("pc preference target not found", hotkey)
            return
          }

          const icon = document.createElement('img')

          const id = ulid(abilityID)
          const iconID = ulid(this.Abilities.get(id)?.getIcon_asU8()!)
          logger.info("assign icon:", iconID)
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
        logger.info(error)
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
      this.Entity.Orientation = Cell.Orientation.UP
    } else if (this.Keys.get(Action.Down)?.isDown) {
      animationID = this.Entity.Animations.get('walk_down')
      this.Entity.Body.body.setVelocityY(speed)
      this.Entity.Orientation = Cell.Orientation.DOWN
    }

    if (this.Keys.get(Action.Right)?.isDown) {
      animationID = this.Entity.Animations.get('walk_right')
      this.Entity.Body.body.setVelocityX(speed)
      this.Entity.Orientation = Cell.Orientation.RIGHT
    } else if (this.Keys.get(Action.Left)?.isDown) {
      animationID = this.Entity.Animations.get('walk_left')
      this.Entity.Body.body.setVelocityX(-speed)
      this.Entity.Orientation = Cell.Orientation.LEFT
    }

    if (!animationID) {
      switch (this.Entity.Orientation) {
        case Cell.Orientation.UP:
          animationID = this.Entity.Animations.get('idle_up')
          break
        case Cell.Orientation.RIGHT:
          animationID = this.Entity.Animations.get('idle_right')
          break
        case Cell.Orientation.DOWN:
          animationID = this.Entity.Animations.get('idle_down')
          break
        case Cell.Orientation.LEFT:
          animationID = this.Entity.Animations.get('idle_left')
          break
      }
    }

    if (animationID) {
      this.Entity.E.setAnimationid(parse(animationID))

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

    let o: Cell.Orientation = Cell.Orientation.NONE
    const x = this.Entity.Body.body.x
    const y = this.Entity.Body.body.y
    const up = this.Border.get(Cell.Orientation.UP)!
    const right = this.Border.get(Cell.Orientation.RIGHT)!
    const down = this.Border.get(Cell.Orientation.DOWN)!
    const left = this.Border.get(Cell.Orientation.LEFT)!

    if (y < up) {
      if (x < left) {
        o = Cell.Orientation.UPLEFT
      } else if (x < right) {
        o = Cell.Orientation.UP
      } else {
        o = Cell.Orientation.UPRIGHT
      }
    } else if (y < down) {
      if (x < left) {
        o = Cell.Orientation.LEFT
      } else if (x < right) {
        o = Cell.Orientation.NONE
      } else {
        o = Cell.Orientation.RIGHT
      }
    } else {
      if (x < left) {
        o = Cell.Orientation.DOWNLEFT
      } else if (x < right) {
        o = Cell.Orientation.DOWN
      } else {
        o = Cell.Orientation.DOWNRIGHT
      }
    }

    if (o != Cell.Orientation.NONE && this.Loading == undefined) {
      const id = this.Cells.get(o)?.Cell.getId_asU8()
      this.Loading = id


      if (id) {
        this.Entity.E.setCellid(id)

        this.loadMap(o).then(() => {
          logger.info('loaded cell')
        })
      } else {
        // don't load cell, out of world
      }
    }

    // Move entity
    this.Entity.E.setX(Math.round(x))
    this.Entity.E.setY(Math.round(y))

    this.Entities.forEach((e: GraphicEntity) => {
      // self reconciliation
      if (ulid(e.E.getId_asU8()) == ulid(this.Entity.E.getId_asU8())) {
        return
      }

      interpolateGraphicEntity(e)

      const x = e.Sprite.x + ((e.Direction.x - e.Sprite.x) * e.Interpolation)
      const y = e.Sprite.y + ((e.Direction.y - e.Sprite.y) * e.Interpolation)
      e?.Sprite.setX(x)
      e?.Sprite.setY(y)

      const animationID = e.Animations.get(ulid(e.E.getAnimationid_asU8()))
      if (animationID) {
        e?.Sprite.play(animationID, true)
      }
    })
  }

  // Orientation o uses preload surrounded cells
  // Cell.Orientation.NONE loads everything
  async loadMap(o: Cell.Orientation) {
    // call current pc cell
    const req = new CellDTO.ListCellReq()

    req.setIdsList([this.Entity.E.getCellid_asU8()])
    return this.APIClient.listCell(req, this.MetadataSession)
      .then((cells: CellDTO.ListCellResp) => {
        // load current pc cell
        if (cells.getCellsList().length != 1) {
          throw new Error('failed to load cell')
        }

        const c = cells.getCellsList()[0]
        const contig = c.getContiguousMap() as pbmap<number, Uint8Array>

        let newCellIDs: Uint8Array[] = []
        let deletedCells: GraphicCell[] = []

        switch (o) {
          case Cell.Orientation.NONE:
            // assign new cells to load
            newCellIDs.push(
              contig.get(Cell.Orientation.UP)!,
              contig.get(Cell.Orientation.UPRIGHT)!,
              contig.get(Cell.Orientation.RIGHT)!,
              contig.get(Cell.Orientation.DOWNRIGHT)!,
              contig.get(Cell.Orientation.DOWN)!,
              contig.get(Cell.Orientation.DOWNLEFT)!,
              contig.get(Cell.Orientation.LEFT)!,
              contig.get(Cell.Orientation.UPLEFT)!,
            )

            // assign cells to delete
            // ...
            // create new blank graphic cell from loaded cell
            this.Cells.set(Cell.Orientation.NONE, {
              Cell: c,
              Tilemap: this.Blank.Tilemap,
              Layers: this.Blank.Layers,
              Colliders: this.Blank.Colliders,
              Objects: this.Blank.Objects
            })

            this.Cells.delete(Cell.Orientation.UP)
            this.Cells.delete(Cell.Orientation.UPRIGHT)
            this.Cells.delete(Cell.Orientation.RIGHT)
            this.Cells.delete(Cell.Orientation.DOWNRIGHT)
            this.Cells.delete(Cell.Orientation.DOWN)
            this.Cells.delete(Cell.Orientation.DOWNLEFT)
            this.Cells.delete(Cell.Orientation.LEFT)
            this.Cells.delete(Cell.Orientation.UPLEFT)
            break;
          case Cell.Orientation.UP:
            // assign new cells to load
            newCellIDs.push(
              contig.get(Cell.Orientation.UPLEFT)!,
              contig.get(Cell.Orientation.UP)!,
              contig.get(Cell.Orientation.UPRIGHT)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(Cell.Orientation.DOWNLEFT)!,
              this.Cells.get(Cell.Orientation.DOWN)!,
              this.Cells.get(Cell.Orientation.DOWNRIGHT)!,
            )
            // shift preloaded
            this.Cells.set(Cell.Orientation.DOWNLEFT, this.Cells.get(Cell.Orientation.LEFT)!)
            this.Cells.set(Cell.Orientation.DOWN, this.Cells.get(Cell.Orientation.NONE)!)
            this.Cells.set(Cell.Orientation.DOWNRIGHT, this.Cells.get(Cell.Orientation.RIGHT)!)
            this.Cells.set(Cell.Orientation.LEFT, this.Cells.get(Cell.Orientation.UPLEFT)!)
            this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.UP)!)
            this.Cells.set(Cell.Orientation.RIGHT, this.Cells.get(Cell.Orientation.UPRIGHT)!)
            this.Cells.delete(Cell.Orientation.UPLEFT)
            this.Cells.delete(Cell.Orientation.UP)
            this.Cells.delete(Cell.Orientation.UPRIGHT)
            break;
          case Cell.Orientation.UPRIGHT:
            // assign new cells to load
            newCellIDs.push(
              contig.get(Cell.Orientation.UPLEFT)!,
              contig.get(Cell.Orientation.UP)!,
              contig.get(Cell.Orientation.UPRIGHT)!,
              contig.get(Cell.Orientation.RIGHT)!,
              contig.get(Cell.Orientation.DOWNRIGHT)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(Cell.Orientation.UPLEFT)!,
              this.Cells.get(Cell.Orientation.LEFT)!,
              this.Cells.get(Cell.Orientation.DOWNLEFT)!,
              this.Cells.get(Cell.Orientation.DOWN)!,
              this.Cells.get(Cell.Orientation.DOWNRIGHT)!,
            )
            // shift preloaded
            this.Cells.set(Cell.Orientation.LEFT, this.Cells.get(Cell.Orientation.UP)!)
            this.Cells.set(Cell.Orientation.DOWNLEFT, this.Cells.get(Cell.Orientation.NONE)!)
            this.Cells.set(Cell.Orientation.DOWN, this.Cells.get(Cell.Orientation.RIGHT)!)
            this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.UPRIGHT)!)
            this.Cells.delete(Cell.Orientation.UPLEFT)
            this.Cells.delete(Cell.Orientation.UP)
            this.Cells.delete(Cell.Orientation.UPRIGHT)
            this.Cells.delete(Cell.Orientation.RIGHT)
            this.Cells.delete(Cell.Orientation.DOWNRIGHT)
            break;
          case Cell.Orientation.RIGHT:
            // assign new cells to load
            newCellIDs.push(
              contig.get(Cell.Orientation.UPRIGHT)!,
              contig.get(Cell.Orientation.RIGHT)!,
              contig.get(Cell.Orientation.DOWNRIGHT)!
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(Cell.Orientation.UPLEFT)!,
              this.Cells.get(Cell.Orientation.LEFT)!,
              this.Cells.get(Cell.Orientation.DOWNLEFT)!,
            )
            // shift preloaded
            this.Cells.set(Cell.Orientation.UPLEFT, this.Cells.get(Cell.Orientation.UP)!)
            this.Cells.set(Cell.Orientation.LEFT, this.Cells.get(Cell.Orientation.NONE)!)
            this.Cells.set(Cell.Orientation.DOWNLEFT, this.Cells.get(Cell.Orientation.DOWN)!)
            this.Cells.set(Cell.Orientation.UP, this.Cells.get(Cell.Orientation.UPRIGHT)!)
            this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.RIGHT)!)
            this.Cells.set(Cell.Orientation.DOWN, this.Cells.get(Cell.Orientation.DOWNRIGHT)!)
            this.Cells.delete(Cell.Orientation.UPRIGHT)
            this.Cells.delete(Cell.Orientation.RIGHT)
            this.Cells.delete(Cell.Orientation.DOWNRIGHT)
            break;
          case Cell.Orientation.DOWNRIGHT:
            // assign new cells to load
            newCellIDs.push(
              contig.get(Cell.Orientation.DOWNLEFT)!,
              contig.get(Cell.Orientation.DOWN)!,
              contig.get(Cell.Orientation.DOWNRIGHT)!,
              contig.get(Cell.Orientation.RIGHT)!,
              contig.get(Cell.Orientation.UPRIGHT)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(Cell.Orientation.UPLEFT)!,
              this.Cells.get(Cell.Orientation.LEFT)!,
              this.Cells.get(Cell.Orientation.DOWNLEFT)!,
              this.Cells.get(Cell.Orientation.UP)!,
              this.Cells.get(Cell.Orientation.UPRIGHT)!,
            )
            // shift preloaded
            this.Cells.set(Cell.Orientation.LEFT, this.Cells.get(Cell.Orientation.DOWN)!)
            this.Cells.set(Cell.Orientation.UPLEFT, this.Cells.get(Cell.Orientation.NONE)!)
            this.Cells.set(Cell.Orientation.UP, this.Cells.get(Cell.Orientation.RIGHT)!)
            this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.DOWNRIGHT)!)
            this.Cells.delete(Cell.Orientation.DOWNLEFT)
            this.Cells.delete(Cell.Orientation.DOWN)
            this.Cells.delete(Cell.Orientation.DOWNRIGHT)
            this.Cells.delete(Cell.Orientation.RIGHT)
            this.Cells.delete(Cell.Orientation.UPRIGHT)
            break;
          case Cell.Orientation.DOWN:
            // assign new cells to load
            newCellIDs.push(
              contig.get(Cell.Orientation.DOWNLEFT)!,
              contig.get(Cell.Orientation.DOWN)!,
              contig.get(Cell.Orientation.DOWNRIGHT)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(Cell.Orientation.UPLEFT)!,
              this.Cells.get(Cell.Orientation.UP)!,
              this.Cells.get(Cell.Orientation.UPRIGHT)!,
            )
            // shift preloaded
            this.Cells.set(Cell.Orientation.UPLEFT, this.Cells.get(Cell.Orientation.LEFT)!)
            this.Cells.set(Cell.Orientation.UP, this.Cells.get(Cell.Orientation.NONE)!)
            this.Cells.set(Cell.Orientation.UPRIGHT, this.Cells.get(Cell.Orientation.RIGHT)!)
            this.Cells.set(Cell.Orientation.LEFT, this.Cells.get(Cell.Orientation.DOWNLEFT)!)
            this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.DOWN)!)
            this.Cells.set(Cell.Orientation.RIGHT, this.Cells.get(Cell.Orientation.DOWNRIGHT)!)
            this.Cells.delete(Cell.Orientation.DOWNLEFT)
            this.Cells.delete(Cell.Orientation.DOWN)
            this.Cells.delete(Cell.Orientation.DOWNRIGHT)
            break;
          case Cell.Orientation.DOWNLEFT:
            // assign new cells to load
            newCellIDs.push(
              contig.get(Cell.Orientation.DOWNRIGHT)!,
              contig.get(Cell.Orientation.DOWN)!,
              contig.get(Cell.Orientation.DOWNLEFT)!,
              contig.get(Cell.Orientation.LEFT)!,
              contig.get(Cell.Orientation.UPLEFT)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(Cell.Orientation.UPLEFT)!,
              this.Cells.get(Cell.Orientation.UP)!,
              this.Cells.get(Cell.Orientation.UPRIGHT)!,
              this.Cells.get(Cell.Orientation.RIGHT)!,
              this.Cells.get(Cell.Orientation.DOWNRIGHT)!,
            )
            // shift preloaded
            this.Cells.set(Cell.Orientation.RIGHT, this.Cells.get(Cell.Orientation.DOWN)!)
            this.Cells.set(Cell.Orientation.UPRIGHT, this.Cells.get(Cell.Orientation.NONE)!)
            this.Cells.set(Cell.Orientation.UP, this.Cells.get(Cell.Orientation.LEFT)!)
            this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.DOWNLEFT)!)
            this.Cells.delete(Cell.Orientation.DOWNRIGHT)
            this.Cells.delete(Cell.Orientation.DOWN)
            this.Cells.delete(Cell.Orientation.DOWNLEFT)
            this.Cells.delete(Cell.Orientation.LEFT)
            this.Cells.delete(Cell.Orientation.UPLEFT)
            break;
          case Cell.Orientation.LEFT:
            // assign new cells to load
            newCellIDs.push(
              contig.get(Cell.Orientation.UPLEFT)!,
              contig.get(Cell.Orientation.LEFT)!,
              contig.get(Cell.Orientation.DOWNLEFT)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(Cell.Orientation.UPRIGHT)!,
              this.Cells.get(Cell.Orientation.RIGHT)!,
              this.Cells.get(Cell.Orientation.DOWNRIGHT)!,
            )
            // shift preloaded
            this.Cells.set(Cell.Orientation.UPRIGHT, this.Cells.get(Cell.Orientation.UP)!)
            this.Cells.set(Cell.Orientation.RIGHT, this.Cells.get(Cell.Orientation.NONE)!)
            this.Cells.set(Cell.Orientation.DOWNRIGHT, this.Cells.get(Cell.Orientation.DOWN)!)
            this.Cells.set(Cell.Orientation.UP, this.Cells.get(Cell.Orientation.UPLEFT)!)
            this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.LEFT)!)
            this.Cells.set(Cell.Orientation.DOWN, this.Cells.get(Cell.Orientation.DOWNLEFT)!)
            this.Cells.delete(Cell.Orientation.UPLEFT)
            this.Cells.delete(Cell.Orientation.LEFT)
            this.Cells.delete(Cell.Orientation.DOWNLEFT)
            break;
          case Cell.Orientation.UPLEFT:
            // assign new cells to load
            newCellIDs.push(
              contig.get(Cell.Orientation.UPRIGHT)!,
              contig.get(Cell.Orientation.UP)!,
              contig.get(Cell.Orientation.UPLEFT)!,
              contig.get(Cell.Orientation.LEFT)!,
              contig.get(Cell.Orientation.DOWNLEFT)!,
            )

            // assign cells to delete
            deletedCells.push(
              this.Cells.get(Cell.Orientation.UPRIGHT)!,
              this.Cells.get(Cell.Orientation.RIGHT)!,
              this.Cells.get(Cell.Orientation.DOWNRIGHT)!,
              this.Cells.get(Cell.Orientation.DOWN)!,
              this.Cells.get(Cell.Orientation.DOWNLEFT)!,
            )
            // shift preloaded
            this.Cells.set(Cell.Orientation.RIGHT, this.Cells.get(Cell.Orientation.UP)!)
            this.Cells.set(Cell.Orientation.DOWNRIGHT, this.Cells.get(Cell.Orientation.NONE)!)
            this.Cells.set(Cell.Orientation.DOWN, this.Cells.get(Cell.Orientation.LEFT)!)
            this.Cells.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.UPLEFT)!)
            this.Cells.delete(Cell.Orientation.UPRIGHT)
            this.Cells.delete(Cell.Orientation.UP)
            this.Cells.delete(Cell.Orientation.UPLEFT)
            this.Cells.delete(Cell.Orientation.LEFT)
            this.Cells.delete(Cell.Orientation.DOWNLEFT)
            break;
        }

        // update border after none cell is up to date
        const cn = this.Cells.get(Cell.Orientation.NONE)!
        this.Border.set(Cell.Orientation.UP, cn.Cell.getY() * this.World.getCellheight())
        this.Border.set(Cell.Orientation.RIGHT, (cn.Cell.getX() + 1) * this.World.getCellwidth())
        this.Border.set(Cell.Orientation.DOWN, (cn.Cell.getY() + 1) * this.World.getCellheight())
        this.Border.set(Cell.Orientation.LEFT, cn.Cell.getX() * this.World.getCellwidth())

        this.cleanCells(deletedCells).then(() => { logger.info('finish to destroy unused tilemaps') })

        // reassign cellsbyid with new cells
        this.CellsByID.clear()
        this.Cells.forEach((v, k) => {
          if (v) {
            this.CellsByID.set(ulid(v.Cell.getId_asU8()), k)
          }
        })

        const req = new CellDTO.ListCellReq()
        req.setIdsList(newCellIDs.filter((v) => (!(v == undefined))))
        return this.APIClient.listCell(req, this.MetadataSession)
      })
      .then((cells: CellDTO.ListCellResp) => {

        // map new loaded cells
        const cellMap = new Map<string, Cell.Cell>()
        cells.getCellsList().map((v) => {
          cellMap.set(ulid(v.getId_asU8()), v)
        })
        if (o == Cell.Orientation.NONE) {
          const c = this.Cells.get(Cell.Orientation.NONE)!
          cellMap.set(ulid(c?.Cell.getId_asU8()), c.Cell)
        }

        const contig = this.Cells.get(Cell.Orientation.NONE)?.Cell.getContiguousMap() as pbmap<number, Uint8Array>
        if (o == Cell.Orientation.NONE) {
          contig.set(Cell.Orientation.NONE, this.Cells.get(Cell.Orientation.NONE)?.Cell.getId_asU8()!)
        }

        // must fit above loaded cells array
        const loadedCells = new Array<Cell.Orientation>()
        switch (o) {
          case Cell.Orientation.NONE:
            loadedCells.push(
              Cell.Orientation.NONE,
              Cell.Orientation.UP,
              Cell.Orientation.UPRIGHT,
              Cell.Orientation.RIGHT,
              Cell.Orientation.DOWNRIGHT,
              Cell.Orientation.DOWN,
              Cell.Orientation.DOWNLEFT,
              Cell.Orientation.LEFT,
              Cell.Orientation.UPLEFT
            )
            break;
          case Cell.Orientation.UP:
            loadedCells.push(
              Cell.Orientation.UP,
              Cell.Orientation.UPRIGHT,
              Cell.Orientation.UPLEFT
            )
            break;
          case Cell.Orientation.UPRIGHT:
            loadedCells.push(
              Cell.Orientation.UP,
              Cell.Orientation.UPRIGHT,
              Cell.Orientation.RIGHT,
              Cell.Orientation.DOWNRIGHT,
              Cell.Orientation.UPLEFT
            )
            break;
          case Cell.Orientation.RIGHT:
            loadedCells.push(
              Cell.Orientation.UPRIGHT,
              Cell.Orientation.RIGHT,
              Cell.Orientation.DOWNRIGHT,
            )
            break;
          case Cell.Orientation.DOWNRIGHT:
            loadedCells.push(
              Cell.Orientation.UPRIGHT,
              Cell.Orientation.RIGHT,
              Cell.Orientation.DOWNRIGHT,
              Cell.Orientation.DOWN,
              Cell.Orientation.DOWNLEFT,
            )
            break;
          case Cell.Orientation.DOWN:
            loadedCells.push(
              Cell.Orientation.DOWNRIGHT,
              Cell.Orientation.DOWN,
              Cell.Orientation.DOWNLEFT,
            )
            break;
          case Cell.Orientation.DOWNLEFT:
            loadedCells.push(
              Cell.Orientation.DOWNRIGHT,
              Cell.Orientation.DOWN,
              Cell.Orientation.DOWNLEFT,
              Cell.Orientation.LEFT,
              Cell.Orientation.UPLEFT
            )
            break;
          case Cell.Orientation.LEFT:
            loadedCells.push(
              Cell.Orientation.DOWNLEFT,
              Cell.Orientation.LEFT,
              Cell.Orientation.UPLEFT
            )
            break;
          case Cell.Orientation.UPLEFT:
            loadedCells.push(
              Cell.Orientation.UP,
              Cell.Orientation.UPRIGHT,
              Cell.Orientation.DOWNLEFT,
              Cell.Orientation.LEFT,
              Cell.Orientation.UPLEFT
            )
            break;
        }

        let loaded = 0
        loadedCells.map((o) => {
          const resetLoad = () => {
            loaded++
            if (loaded == loadedCells.length) {
              logger.info('loading unlock')
              this.Loading = undefined
            }
          }

          if (!contig.has(o)) {
            // world border
            logger.info('contig has no:', o)
            resetLoad()
            return
          }

          const c = cellMap.get(ulid(contig.get(o)!))!
          if (!c) {
            // world border
            // TO INVESTIGATE, shouldn't happen ?
            logger.info('cell not found:', o)
            return
          }

          const tm = ulid(c.getTilemap_asU8())
          this.CellLoader.tilemapTiledJSON(tm, 'json/assets/' + tm + '.json')

          this.Cells.set(o, {
            Cell: c,
            Tilemap: this.Blank.Tilemap,
            Layers: new Map(),
            Colliders: new Map(),
            Objects: new Map(),
          })
          this.CellsByID.set(ulid(c.getId_asU8()), o)

          const loadTM = () => {
            logger.info('load_tilemap ', o)

            // create new cell
            const map = this.make.tilemap({ key: tm, width: this.World.getCellwidth(), height: this.World.getCellheight() })

            let sets = new Array<Phaser.Tilemaps.Tileset>()

            const loadTS = (tsName: string) => {
              const ts = map.addTilesetImage(tsName)

              if (!ts) {
                logger.info('failed to add tileset image ' + tsName)
                return
              }

              sets.push(ts)

              logger.info('loaded TS ', sets.length, '/', map.tilesets.length)
              if (sets.length < map.tilesets.length) {
                return
              }

              const cc = this.Cells.get(o)
              if (!cc) {
                // should never happen
                return
              }

              cc.Tilemap = map

              const x = c.getX() * this.World.getCellwidth()
              const y = c.getY() * this.World.getCellheight()

              map.layers.map((l) => {
                const layer = map.createLayer(l.name, sets, x, y)
                if (!layer) {
                  logger.info('failed to create layer:', l.name, x, y)
                  return
                }

                // TODO: required or not ?
                layer.setPipeline('main')

                logger.info('created layer:', l.name, x, y)
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

                logger.info('created object layer:', os.name)
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
                logger.info('add listener on ', 'filecomplete-image-' + ts.name)
                this.CellLoader.on('filecomplete-image-' + ts.name, () => {
                  logger.info('image completed on ', o)
                  loadTS(ts.name)
                })
              }
            })
          }

          // Check if files already been loaded
          if (this.cache.tilemap.exists(tm)) {
            loadTM()
          } else {
            logger.info('add listener on ', 'filecomplete-tilemapJSON-' + tm)
            this.CellLoader.on('filecomplete-tilemapJSON-' + tm, () => {
              logger.info('json completed on ', o)
              loadTM()
            })
          }
        })

        this.CellLoader.on('complete', () => {
          logger.info('reset loader')
          this.CellLoader.removeAllListeners()
          this.CellLoader.reset()
        })

        logger.info('start global cell loading')
        this.CellLoader.start()
      })
  }

  async cleanCells(cells: GraphicCell[]) {
    cells.map((c) => {
      if (c) {
        logger.info('clean cell', ulid(c.Cell.getId_asU8()))
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

    logger.info('cleaning entities', ids)

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

  async displayEntities(message: EntityDTO.ListEntityResp) {
    // load all entities
    let entityIDs: Uint8Array[] = []

    message.getEntitiesList().forEach((entry: Entity.E) => {
      const id = ulid(entry.getId_asU8())
      const meid = ulid(this.Entity.E.getId_asU8())

      // Set last tick to 0
      this.EntityLastTick.set(id, (this.EntityLastTick.get(id) || 0) + 1)

      if (this.Entities.has(id)) {
        // update state only
        const x = entry.getX() //+ (this.Cells.get(this.CellsByID.get(ulid(entry.getCellid_asU8()))!)?.Cell.getX()! * this.World.getCellwidth())
        const y = entry.getY() //+ (this.Cells.get(this.CellsByID.get(ulid(entry.getCellid_asU8()))!)?.Cell.getY()! * this.World.getCellheight())
        updateGraphicEntity(this.Entities.get(id)!, x, y)

        this.Entities.get(id)!.E = entry
      } else {
        // create associated sprite
        // receive own entity from server once and initialize
        if (id == meid) {
          this.Entity.Body.destroy()

          if (entry.getObjectsList().length == 0) {
            logger.info('player entity has no collision body. set default')
            this.Entity.Body = this.physics.add.sprite(entry.getX(), entry.getY(), id).setSize(16, 16).setOffset(0, 0)
          } else {
            // pick first dynamic body from list to assign as main collision object
            const obj = entry.getObjectsList().at(0)!
            this.Entity.Body = this.physics.add.sprite(entry.getX(), entry.getY(), id).
              setSize(obj.getWidth(), obj.getHeight()).
              setOffset(obj.getX(), obj.getY())
          }

          logger.info('set body from server info')
          this.Entity.E.setX(entry.getX())
          this.Entity.E.setY(entry.getY())

          // local player sprite loaded, start camera follow
          this.cameras.main.startFollow(this.Entity.Body)
          this.Entity.Body.setDepth(entitySpriteDepth)

          this.Entities.set(id, {
            E: entry,
            Animations: new Map(),
            Direction: new Phaser.Geom.Point(entry.getX(), entry.getY()),
            Interpolation: 0.00,
            Sprite: this.Entity.Body,

            Objects: new Map(),
            Colliders: new Map(),
          })

          logger.info("receive position from server:", this.Entity.E.getX(), this.Entity.E.getY())

          // load own abilities
          const req = new AbilityDTO.ListAbilityReq()
          req.setEntityid(this.Entity.E.getId_asU8())
          req.setSize(100)
          this.APIClient.listAbility(req, this.MetadataSession)
            .then((resp: AbilityDTO.ListAbilityResp) => {
              resp.getAbilitiesList().forEach((ab: Ability.A) => {
                const id = ulid(ab.getId_asU8())

                const ga: GraphicAbility = {
                  A: ab,
                  Effects: new Array()
                }

                // TODO: group targets in this.Entity.Ability
                // Order to redefine ?
                ab.getEffectsMap().forEach((e: Ability.Effect, effectID: string) => {
                  const ge: GraphicEffect = {
                    ID: effectID,
                    Targets: new Array()
                  }

                  const targetsByGroup = new Map<string, GraphicTarget>()
                  e.getTargetsMap().forEach((target: Ability.Target, targetID: string) => {
                    const groupID = ulid(target.getGroupid_asU8())
                    if (!targetsByGroup.has(groupID)) {
                      targetsByGroup.set(groupID, {
                        Type: target.getType(),
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
              logger.info("connection established")
              this.MapLoaded()
            })
            .catch((err) => { logger.info(err) })
        } else if (this.Entities.has(meid)) {
          const eid = ulid(entry.getId_asU8())
          const sprite = this.add.sprite(entry.getX(), entry.getY(), id)
          sprite.setDepth(entitySpriteDepth - 1)

          // over targeting at sprite creation
          sprite.on('over', () => {
            if (this.Targeting &&
              this.Targeting.Targets.size < this.CastTargets.size &&
              !this.CastTargets.has(eid) &&
              ((this.Targeting.Type == Ability.TargetType.FOE && (entry.getFactionid() != this.Entity.E.getFactionid())) ||
                (this.Targeting.Type == Ability.TargetType.ALLY && (entry.getFactionid() == this.Entity.E.getFactionid())))) {
              const target = new Cast.CastTarget()
              target.setId(eid)
              this.CastTargets.set(eid, target)
            }
          })

          let ge: GraphicEntity = {
            E: entry,
            Animations: new Map(),
            Direction: new Phaser.Geom.Point(entry.getX(), entry.getY()),
            Interpolation: 0.00, // default speed
            Sprite: sprite,

            Objects: new Map(),
            Colliders: new Map(),
          }

          logger.info('set entity: ', id, entry.getX(), entry.getY())

          // set collision objects
          // offset on layer position
          const objects = entry.getObjectsList()
          if (objects.length > 0) {
            const group = this.physics.add.staticGroup(objects.map((b) => {
              logger.info('set entity collision: ', id, b.getX() + ge.E.getX(), b.getY() + ge.E.getY())
              return this.physics.add.
                staticImage(
                  ge.E.getX() + b.getX(),
                  ge.E.getY() + b.getY(),
                  '').
                setSize(b.getWidth(), b.getHeight()).
                setVisible(false).
                setImmovable(true).
                setOffset(0, 0)
            }))

            const collider = this.physics.add.collider(this.Entity.Body, group)

            ge.Objects.set(id, group)
            ge.Colliders.set(id, collider)
            logger.info('created entity collision:', id)
          }

          this.Entities.set(id, ge)
        } else {
          // this.Entity.Body not initialized yet, don't load anything until we have it (for collisions)
          return
        }

        // load entity animations
        entityIDs.push(entry.getId_asU8())
      }
    })

    if (entityIDs.length == 0) {
      return
    }

    // load all animations
    const req = new AnimationDTO.ListAnimationReq()

    req.setEntityidsList(entityIDs)
    req.setSize(1000) // TODO: less random
    this.APIClient.listAnimation(req, this.MetadataSession)
      .then((animations: AnimationDTO.ListAnimationResp) => {
        // load all current animations
        animations.getAnimationsList().forEach((an: Animation.Animation) => {
          const id = ulid(an.getId_asU8())
          const sheetID = ulid(an.getSheetid_asU8())
          const duplicateID = ulid(an.getDuplicateid_asU8())
          const entityID = ulid(an.getEntityid_asU8())

          let anims = this.Entities.get(entityID)?.Animations!

          // loading self animations, change object + add named animations
          if (entityID == ulid(this.Entity.E.getId_asU8())) {
            anims = this.Entity.Animations

            logger.info('set named animation:', an.getName(), id)
            anims.set(an.getName(), id)
          }

          // return if already loaded
          if (this.anims.exists(duplicateID)) {
            anims.set(id, duplicateID)

            return
          }

          const loadAnim = (an: Animation.Animation) => {
            // Create animation
            const seq = an.getSequenceList().length > 0 ? { frames: an.getSequenceList() } : { start: an.getStart(), end: an.getEnd() };
            const newAnim = this.anims.create({
              key: duplicateID,
              frames: this.anims.generateFrameNumbers(
                sheetID,
                seq,
              ),
              frameRate: an.getRate(),
              repeat: an.getRepeat(),
              delay: an.getDelay(),
              duration: an.getDuration(),
              hideOnComplete: an.getShowandhide(),
              showOnStart: an.getShowandhide(),
            })
            if (!newAnim) {
              logger.info('failed to load animation ' + duplicateID)

              return
            }

            logger.info('set animation:', id, duplicateID)

            // Add animation to entity animations
            anims.set(id, duplicateID)
          }

          if (!this.SpriteSheets.get(sheetID)) {
            // load sprite sheet

            this.EntityLoader.spritesheet(sheetID, 'img/assets/' + sheetID + '.png', {
              frameWidth: an.getFramewidth(),
              frameHeight: an.getFrameheight(),
              startFrame: an.getFramestart(),
              endFrame: an.getFrameend(),
              margin: an.getFramemargin(),
              spacing: an.getFramespacing(),
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
}
