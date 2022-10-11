# game_03
---

Dev setup:

```sh
$ cd cmd/browser && npm install
$ GO111MODULE=off go get github.com/gogo/protobuf/proto
$ go install github.com/gogo/protobuf/protoc-gen-gogoslick
```

Setup:

```sh
$ docker-compose up -d # wait ~10 sec for scylla to boot
$ cat scripts/keyspace.sh | docker exec -i game_03_scylla cqlsh
$ make admin && ./bin/game_03_admin config/admin/local.json
$ grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '"cql"' -plaintext localhost:8083 grpc.Admin/MigrateUp
$ make api && ./bin/game_03_api config/api/local.json
$ make auth && ./bin/game_03_auth config/auth/local.json
$ make browser
$ make web && ./bin/game_03_web config/web/local.json
```

Data:

```sh
$ ./scripts/create_default_templates.sh
$ ./scripts/create_default_tilesheets.sh
$ grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '' -plaintext localhost:8083 grpc.Admin/CreateWorld
```

### TODO
---

- Manage collision
- Manage map creation
- Add default animations as JSON in some file somewhere + add default animations with at createPC
- Remove `Math.round()` in `game.ts` and set entity.X entity.Y as float64 to fit Phaser.Body x/y
- Manage twitch loadFollow error when timeout (ez to reproduce...)


# FLOW
---

Note: `CreateTilesheet` is optional when using local files in `cmd/browser/dist/`.

Instead you can use `scripts/create_default_tilesheets.sh` once from current directory to "load" local `scripts/assets` into local browser.

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
