# game_03
---

Dev once setup:

```sh
$ cd cmd/dashboard && npm install
$ cd cmd/client && npm install
$ GO111MODULE=off go get github.com/gogo/protobuf/proto
$ go install github.com/gogo/protobuf/protoc-gen-gogoslick
$ cat scripts/add_localhost.hosts | sudo tee -a /etc/hosts > /dev/null
```

Dev setup:

```sh
$ docker-compose up -d # wait ~10 sec for scylla to boot
$ cat docker/scylla/keyspace.cql | docker exec -i game_03_scylla cqlsh
$ make admin && ./bin/game_03_admin config/admin/local.json
$ grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '"cql"' -plaintext localhost:4282 grpc.Admin/MigrateUp
$ make populate
$ make api && ./bin/game_03_api config/api/local.json
$ make auth && ./bin/game_03_auth config/auth/local.json
$ make client
$ make web_client && ./bin/game_03_web config/web_client/local.json
```

Upload assets (`make populate`):

```sh
$ ./scripts/upload_default_images.sh
$ ./scripts/create_default_templates.sh
$ ./scripts/create_default_tilesets.sh
$ ./scripts/create_default_animations.sh
$ ./scripts/create_default_abilities.sh
$ grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '' -plaintext localhost:4282 grpc.Admin/CreateWorld
```

Regenerate assets from `assets/external` to `assets/` (`make regen-assets`):

```sh
$ go run ./scripts/write_animation/main.go 'assets/animations'
$ go run ./scripts/write_tileset/main.go 'assets/external/Tilesets' 'assets/tilesets'
```

### TODO
---

- [p0] -> blocking game
- [p1] -> blocking feature
- [p2] -> annoying game
- [p3] -> annoying feature
- [p4] -> improvement

- [EPIC] Manage map creation
  + [ ] [p4] Fix collision values gaps in `.json` adding useless rectangles in min dissection
  + [ ] [p4] Add sky background (bonus: animated tiles) -> as parameter (different background possible in same map too ?)
  + [ ] [p4] Fix core crtl+c context error (grpc panic drain ?)
  + [ ] [p3] Change dashboard/rooms to display post login (event/state propag)
  + [ ] [p2] !!! Feature remove one-tile grass paths frustrating in generation
  + [ ] [p2] Fix interpolation speed in game.ts (corresponding actual player speed)
  + [ ] [p1] Console warnings (webGL mostly)
  + [ ] [p0] Audio
  + [ ] [p0] Remove `Math.round()` in `game.ts` and set entity.X entity.Y as float64 to fit Phaser.Body x/y
  + [ ] [p0] [HD] (re-)implement spawns -> change into point of interest per map to fill or not
    + [x] NPC entity have lower z index
    + [~]  NPC entity static/dynamic collision -> Add manual (in code ?) collisions for altar in room/app
    + [ ] implement spawn mechanic
	+ [ ] implement skill mechanic
	  + [ ] AbilityModifier + EffectModifier
	  + [ ] Triggers at component level
	  + [ ] Eval triggers first, then modify effects then apply effects
	  + [x] Ability store+cache + Entity Ability store+cache
	  + [ ] [HD] API routes for ability + ability entity
	  + [ ] [HD] Add faction implementation all along :(
  + [ ] [p0] Implement RTC both directions
	  + [x] Clean cancel of send_entity
	  + [ ] When ctrl+c -> cancel ctx clean
	  + [ ] RTC peer connectin not found 1/2 refresh (delete/recreate mechanism)
  + [ ] [BUG] [p0] Collision on entity don't happen (entity loaded before pc ?)

# FLOW
---

Note: `CreateTilesheet` is optional when using local files in `cmd/client/dist/`.

Instead you can use `scripts/create_default_tilesheets.sh` once from current directory to "load" local `scripts/assets` into local client.

### scratch setup

### Entity

- `CreateTemplate`
- `CreateTilesheet`
- `CreateAnimation`

### World

- `CreateTilesheet`
- `CreateWorld` (add a allow_entity_template list by name ?)

### Room

- `CreateRoom`

### PC

- `CreatePC` (using entity_template.name)
- `ConnectPC`


## Dev local port assignments

- `web_client` -> `8080`
- `web_dashboard` -> `8081`
- `api` -> `http:8082` `grpc:4280`
- `auth` -> `4281`
- `admin` -> `4282`
- `core` -> `http:8083` `grpc:4283`


### IDEAS
---
- When killing a boss, grant a X power materia/gear for final hit guild -> Boss belongs to guild, can't kill it anymore -> When another guilds kill it (final hit guaranteed), grant a X+1 power materia/gear for guild and so on...
- How to implement TP, lighthouse ?

- classes:
  + rage warrior, when conditions that don't follow gameplan (opponent debuff disappearing, debuff on self, distance too long), warrior gain rage (ability ? resource ? buff ?)
  + mage, long gameplan, need to consume buff on opponent to buff itself and gain big (ability ? resource ? buff ?)
  + thief, looking for specific timing/conditions (how to make that thief can't trigger himself those conditions, need to be opponent)

### Unrelated
---

`sh template`

```sh
#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail
if [[ "${TRACE-0}" == "1" ]]; then
    set -o xtrace
fi

if [[ "${1-}" =~ ^-*h(elp)?$ ]]; then
    echo 'Usage: ./script.sh arg-one arg-two

This is an awesome bash script to make your life better.

'
    exit
fi

cd "$(dirname "$0")"

main() {
    echo do awesome stuff
}

main "$@"
```
