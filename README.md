# game_03

(WIP)

```sh
$ docker-compose up -d
$ sh scripts/keyspace.sh // not really...
$ make admin && ./bin/game_03_admin config/admin/local.json
$ grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '"cql"' -plaintext localhost:8083 grpc.Admin/MigrateUp
$ make api && ./bin/game_03_api config/api/local.json
$ make auth && ./bin/game_03_auth config/auth/local.json
$ make browser
$ make web && ./bin/game_03_web config/web/local.json
```

### TODO
- adapt db model to add `Tileset` to PC
- Fetch PC tileset at game load
