import { APIClient } from 'cmd/api/grpc/ApiServiceClientPb';
import { ListAbilityReq, ListAbilityResp } from "pkg/ability/dto/ability_pb"
import { Metadata } from 'grpc-web';
import { A, Effect, Target } from 'pkg/ability/ability_pb';
import { ulid } from '../../lib/ulid';
import { GAbility, GEffect, GTarget } from './ability';

const logger = require('pino')({
  name: 'ability_self',
  level: 'info',
})

export class Self {
  private abilities: Map<string, GAbility>

  constructor(
    private id: Uint8Array,
    private APIClient: APIClient,
    private MetadataSession: Metadata,
  ) {
    const req = new ListAbilityReq()
    req.setEntityid(id)
    req.setSize(100) // paginate if more ?
    this.APIClient.listAbility(req, this.MetadataSession)
      .then((resp: ListAbilityResp) => {
        resp.getAbilitiesList().forEach((ab: A) => {
          const id = ulid(ab.getId_asU8())

          const ga: GAbility = {
            A: ab,
            Effects: new Array()
          }

          // TODO: group targets in this.self.Ability
          // Order to redefine ?
          ab.getEffectsMap().forEach((e: Effect, effectID: string) => {
            const targetsByGroup = new Map<string, GTarget>()
            e.getTargetsMap().forEach((target: Target, targetID: string) => {
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

            const targets = new Array<GTarget>()
            // specify target order here ?
            targetsByGroup.forEach((v, k) => {
              targets.push(v)
            })

            ga.Effects.push({
              ID: effectID,
              Targets: targets,
            })
          })

          // this.Abilities.set(id, ab)
          this.abilities.set(id, ga)
        })
      })
      .then(() => {
        logger.info("self abilities initialized")
        // this.MapLoaded()
      })
      .catch((err) => { logger.info(err) })
  }
}
