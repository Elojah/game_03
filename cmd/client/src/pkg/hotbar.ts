import * as Ability from 'pkg/ability/ability_pb';
import * as Cast from 'pkg/ability/cast_pb';
import { ulid, parse } from '../lib/ulid'

const logger = require('pino')()

export enum Action {
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
  Type: Ability.TargetType
  GroupID: string
  Targets: Map<string, Ability.Target>
}

export class Hotbar {
  Input: Phaser.Input.InputPlugin

  // Abilities for entity
  Abilities: Map<string, Ability.A>

  Keys: Map<Action, Phaser.Input.Keyboard.Key>

  CastTargets: Map<string, Cast.CastTarget>
  Casting: string | undefined
  Targeting: GraphicTarget | undefined

  constructor() { }

  static readonly hotkeyIcons = new Map<Action, string>([
    [Action.Hotkey00, 'hotkey-0-0-icon'],
    [Action.Hotkey01, 'hotkey-0-1-icon'],
    [Action.Hotkey02, 'hotkey-0-2-icon'],
    [Action.Hotkey03, 'hotkey-0-3-icon'],
    [Action.Hotkey04, 'hotkey-0-4-icon'],
    [Action.Hotkey05, 'hotkey-0-5-icon'],
    [Action.Hotkey10, 'hotkey-1-0-icon'],
    [Action.Hotkey11, 'hotkey-1-1-icon'],
    [Action.Hotkey12, 'hotkey-1-2-icon'],
    [Action.Hotkey13, 'hotkey-1-3-icon'],
    [Action.Hotkey14, 'hotkey-1-4-icon'],
    [Action.Hotkey15, 'hotkey-1-5-icon'],
  ])

  public init() {
    this.Keys = new Map()

    this.Keys.set(Action.Up, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.W)!)
    this.Keys.set(Action.Right, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.D)!)
    this.Keys.set(Action.Down, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.S)!)
    this.Keys.set(Action.Left, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.A)!)
    this.Keys.set(Action.Hotkey00, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.ONE)!)
    this.Keys.set(Action.Hotkey01, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.TWO)!)
    this.Keys.set(Action.Hotkey02, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.THREE)!)
    this.Keys.set(Action.Hotkey03, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.FOUR)!)
    this.Keys.set(Action.Hotkey04, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.FIVE)!)
    this.Keys.set(Action.Hotkey05, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.SIX)!)
    this.Keys.set(Action.Hotkey10, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.SEVEN)!)
    this.Keys.set(Action.Hotkey11, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.EIGHT)!)
    this.Keys.set(Action.Hotkey12, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.NINE)!)
    this.Keys.set(Action.Hotkey13, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.ZERO)!)
    this.Keys.set(Action.Hotkey14, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.COMMA)!)
    this.Keys.set(Action.Hotkey15, this.Input.keyboard?.addKey(Phaser.Input.Keyboard.KeyCodes.PERIOD)!)

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

    const eid = Hotbar.hotkeyIcons.get(a)

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
    this.Input.manager.canvas.style.cursor = 'crosshair'

    switch (gt.Type) {
      case Ability.TargetType.NONETARGET: { break }
      case Ability.TargetType.SELF: {
        const ct = new Cast.CastTarget()
        ct.setId(this.EntityID)
        this.CastTargets.set(this.EntityID, ct)

        // instantly returns instead of waiting confirmation
        this.Input.manager.canvas.style.cursor = 'default'

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

        this.Input.on('pointerdown', (pointer: Phaser.Input.Pointer) => {
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

        this.Input.on('pointerdown', (pointer: Phaser.Input.Pointer) => {
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
        this.Input.removeListener('pointerdown')
        this.Targeting = undefined
        this.Input.manager.canvas.style.cursor = 'default'
        accept(true)
      })
    })
  }


}
