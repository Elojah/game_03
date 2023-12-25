import { A, Target, TargetType } from "pkg/ability/ability_pb"

export type GTarget = {
  Type: TargetType
  GroupID: string
  Targets: Map<string, Target>
}

export type GEffect = {
  ID: string
  Targets: Array<GTarget>
}

export type GAbility = {
  A: A
  // Targets by effect
  // Targets array define order in which target groups are defined
  Effects: Array<GEffect>
}
