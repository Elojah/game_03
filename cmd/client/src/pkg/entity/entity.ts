import { E } from "pkg/entity/entity_pb"
import { Orientation } from "pkg/room/cell_pb"
import { GAbility } from "../ability/ability"

export type GEntity = {
  E: E
  Animations: Map<string, string>
  Direction: Phaser.Geom.Point
  Interpolation: number
  Sprite: Phaser.GameObjects.Sprite

  Objects: Map<string, Phaser.Physics.Arcade.StaticGroup>
  Colliders: Map<string, Phaser.Physics.Arcade.Collider>
}


export type BodyEntity = {
  E: E
  Body: Phaser.Types.Physics.Arcade.SpriteWithDynamicBody
  Animations: Map<string, string>
  Orientation: Orientation
  Abilities: Map<string, GAbility>
}
