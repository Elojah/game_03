# game_03
---

Dev setup:

```sh
$ cd cmd/client && npm install
$ GO111MODULE=off go get github.com/gogo/protobuf/proto
$ go install github.com/gogo/protobuf/protoc-gen-gogoslick
$ cat scripts/add_localhost.hosts | sudo tee -a /etc/hosts > /dev/null
```

Setup:

```sh
$ docker-compose up -d # wait ~10 sec for scylla to boot
$ cat scripts/keyspace.cql | docker exec -i game_03_scylla cqlsh
$ make admin && ./bin/game_03_admin config/admin/local.json
$ grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '"cql"' -plaintext localhost:4282 grpc.Admin/MigrateUp
$ make api && ./bin/game_03_api config/api/local.json
$ make auth && ./bin/game_03_auth config/auth/local.json
$ make client
$ make web_client && ./bin/game_03_web config/web_client/local.json
```

Upload data:

```sh
$ ./scripts/upload_default_images.sh
$ ./scripts/create_default_templates.sh
$ ./scripts/create_default_tilesets.sh
$ ./scripts/create_default_animations.sh
$ grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '' -plaintext localhost:4282 grpc.Admin/CreateWorld
```

Regenerate data from `assets/`:

```sh
$ go run ./scripts/write_animation/main.go 'assets/animations'
$ go run ./scripts/write_tileset/main.go 'assets/external/Tilesets' 'assets/tilesets'
```

### TODO
---

- Manage map creation
- Remove `Math.round()` in `game.ts` and set entity.X entity.Y as float64 to fit Phaser.Body x/y


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

- `web` -> `8080`
- `web_dashboard` -> `8081`
- `api` -> `http:8082` `grpc:4280`
- `auth` -> `4281`
- `admin` -> `4282`



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
