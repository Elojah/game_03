import { A, Target, TargetType } from "pkg/ability/ability_pb"

type GTarget = {
  Type: TargetType
  GroupID: string
  Targets: Map<string, Target>
}


type Effect = {
  ID: string
  Targets: Array<GTarget>
}

export type Ability = {
  A: A
  // Targets by effect
  // Targets array define order in which target groups are defined
  Effects: Array<Effect>
}
