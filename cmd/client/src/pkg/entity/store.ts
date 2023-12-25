import { E } from "pkg/entity/entity_pb";
import { ulid } from "../../lib/ulid";
import { Entity, BodyEntity } from "./entity";
import * as EntityDTO from 'pkg/entity/dto/entity_pb';

const logger = require('pino')({
  name: 'entity_store',
  level: 'info',
  prettyPrint: { colorize: true }
})

export class Store {
  constructor(
    private self: BodyEntity,
    private physics: Phaser.Physics.Arcade.ArcadePhysics,
    private cameras: Phaser.Cameras.Scene2D.CameraManager,
  ) { }

  private tick: Map<string, number>

  // cleaning delay
  private delay: number

  private store: Map<string, Entity> = new Map();

  public get(id: string): Entity | undefined {
    return this.store.get(id);
  }

  public set(id: string, entity: E): void {
    const last = this.store.get(id);

    if (last) {
      // only update position
      if (last.Direction.x != entity.getX() || last.Direction.y != entity.getY()) {
        last.Interpolation = 0
      }
      last.Direction.x = entity.getX()
      last.Direction.y = entity.getY()

      this.store.set(id, last)

      return
    }


  }

  public update(): void {
    this.store.forEach((e: Entity) => {
      // self reconciliation
      if (ulid(e.E.getId_asU8()) == ulid(this.self.E.getId_asU8())) {
        return
      }

      if (e.Interpolation < 1) {
        e.Interpolation += 0.05
      }

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

  public async initSelf(entity: E): Promise<void> {
    this.self.Body.destroy()

    const id = ulid(entity.getId_asU8())

    if (entity.getObjectsList().length == 0) {
      logger.info('player entity has no collision body. set default')
      this.self.Body = this.physics.add.sprite(entity.getX(), entity.getY(), id).setSize(16, 16).setOffset(0, 0)
    } else {
      // pick first dynamic body from list to assign as main collision object
      const obj = entity.getObjectsList().at(0)!
      this.self.Body = this.physics.add.sprite(entity.getX(), entity.getY(), id).
        setSize(obj.getWidth(), obj.getHeight()).
        setOffset(obj.getX(), obj.getY())
    }

    logger.info('set body from server info')
    this.self.E.setX(entity.getX())
    this.self.E.setY(entity.getY())

    // local player sprite loaded, start camera follow
    this.cameras.main.startFollow(this.self.Body)
    this.self.Body.setDepth(entitySpriteDepth)

    this.Entities.set(id, {
      E: entity,
      Animations: new Map(),
      Direction: new Phaser.Geom.Point(entity.getX(), entity.getY()),
      Interpolation: 0.00,
      Sprite: this.self.Body,

      Objects: new Map(),
      Colliders: new Map(),
    })

    logger.info("receive position from server:", this.self.E.getX(), this.self.E.getY())

    // load own abilities
    const req = new AbilityDTO.ListAbilityReq()
    req.setEntityid(this.self.E.getId_asU8())
    req.setSize(100)
    this.APIClient.listAbility(req, this.MetadataSession)
      .then((resp: AbilityDTO.ListAbilityResp) => {
        resp.getAbilitiesList().forEach((ab: Ability.A) => {
          const id = ulid(ab.getId_asU8())

          const ga: GraphicAbility = {
            A: ab,
            Effects: new Array()
          }

          // TODO: group targets in this.self.Ability
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
          this.self.Abilities.set(id, ga)
        })
      })
      .then(() => {
        logger.info("connection established")
        this.MapLoaded()
      })
      .catch((err) => { logger.info(err) })
  }

  public async synchronize(message: EntityDTO.ListEntityResp): Promise<void> {
    // load all entities
    let entityIDs: Uint8Array[] = []

    message.getEntitiesList().forEach((entry: E) => {
      const id = ulid(entry.getId_asU8())
      const selfID = ulid(this.self.E.getId_asU8())

      // Set last tick to 0
      this.tick.set(id, (this.tick.get(id) || 0) + 1)

      this.set(id, entry)

      // TMP TO MOVE INTO .set
      if (this.store.has(id)) {
        // update state only
        const x = entry.getX() //+ (this.Cells.get(this.CellsByID.get(ulid(entry.getCellid_asU8()))!)?.Cell.getX()! * this.World.getCellwidth())
        const y = entry.getY() //+ (this.Cells.get(this.CellsByID.get(ulid(entry.getCellid_asU8()))!)?.Cell.getY()! * this.World.getCellheight())
        updateGraphicEntity(this.Entities.get(id)!, x, y)

        this.Entities.get(id)!.E = entry
      } else {
        // create associated sprite
        // receive own entity from server once and initialize
        if (id == selfID) {
          this.self.Body.destroy()

          if (entry.getObjectsList().length == 0) {
            logger.info('player entity has no collision body. set default')
            this.self.Body = this.physics.add.sprite(entry.getX(), entry.getY(), id).setSize(16, 16).setOffset(0, 0)
          } else {
            // pick first dynamic body from list to assign as main collision object
            const obj = entry.getObjectsList().at(0)!
            this.self.Body = this.physics.add.sprite(entry.getX(), entry.getY(), id).
              setSize(obj.getWidth(), obj.getHeight()).
              setOffset(obj.getX(), obj.getY())
          }

          logger.info('set body from server info')
          this.self.E.setX(entry.getX())
          this.self.E.setY(entry.getY())

          // local player sprite loaded, start camera follow
          this.cameras.main.startFollow(this.self.Body)
          this.self.Body.setDepth(entitySpriteDepth)

          this.Entities.set(id, {
            E: entry,
            Animations: new Map(),
            Direction: new Phaser.Geom.Point(entry.getX(), entry.getY()),
            Interpolation: 0.00,
            Sprite: this.self.Body,

            Objects: new Map(),
            Colliders: new Map(),
          })

          logger.info("receive position from server:", this.self.E.getX(), this.self.E.getY())

          // load own abilities
          const req = new AbilityDTO.ListAbilityReq()
          req.setEntityid(this.self.E.getId_asU8())
          req.setSize(100)
          this.APIClient.listAbility(req, this.MetadataSession)
            .then((resp: AbilityDTO.ListAbilityResp) => {
              resp.getAbilitiesList().forEach((ab: Ability.A) => {
                const id = ulid(ab.getId_asU8())

                const ga: GraphicAbility = {
                  A: ab,
                  Effects: new Array()
                }

                // TODO: group targets in this.self.Ability
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
                this.self.Abilities.set(id, ga)
              })
            })
            .then(() => {
              logger.info("connection established")
              this.MapLoaded()
            })
            .catch((err) => { logger.info(err) })
        } else if (this.Entities.has(selfID)) {
          const eid = ulid(entry.getId_asU8())
          const sprite = this.add.sprite(entry.getX(), entry.getY(), id)
          sprite.setDepth(entitySpriteDepth - 1)

          // over targeting at sprite creation
          sprite.on('over', () => {
            if (this.Targeting &&
              this.Targeting.Targets.size < this.CastTargets.size &&
              !this.CastTargets.has(eid) &&
              ((this.Targeting.Type == Ability.TargetType.FOE && (entry.getFactionid() != this.self.E.getFactionid())) ||
                (this.Targeting.Type == Ability.TargetType.ALLY && (entry.getFactionid() == this.self.E.getFactionid())))) {
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

            const collider = this.physics.add.collider(this.self.Body, group)

            ge.Objects.set(id, group)
            ge.Colliders.set(id, collider)
            logger.info('created entity collision:', id)
          }

          this.Entities.set(id, ge)
        } else {
          // this.self.Body not initialized yet, don't load anything until we have it (for collisions)
          return
        }

        // load entity animations
        entityIDs.push(entry.getId_asU8())
      }
    })
  }
