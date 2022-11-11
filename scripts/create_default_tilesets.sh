#!/usr/bin/env bash
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d @ -plaintext localhost:4282 grpc.Admin/CreateTileset <assets/tilesets/RuinsWall.json
