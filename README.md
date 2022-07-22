# game_03

Development setup:

Initial prep:

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

Initial data:

```sh
$ ./scripts/create_default_animations.sh
```

### TODO

- Manage collision
- Manage map creation
- Add default animations as JSON in some file somewhere + add default animations with at createPC
