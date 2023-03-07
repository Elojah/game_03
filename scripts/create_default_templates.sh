#!/usr/bin/env bash

# Character templates
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "BlueNinja"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "BlueSamurai"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Boy"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Cavegirl"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Cavegirl2"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Caveman"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Caveman2"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Child"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "DarkNinja"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "EggBoy"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "EggGirl"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Eskimo"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "EskimoNinja"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "GoldKnight"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "GrayNinja"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Greenman"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "GreenNinja"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Inspector"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Knight"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Lion"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "MaskedNinja"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "MaskFrog"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Master"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Monk"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Monk2"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Noble"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "OldMan"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "OldMan2"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "OldMan3"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "OldWoman"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Princess"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "RedNinja"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "RedNinja2"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "RedSamurai"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Samurai"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Skeleton"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Villager"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Villager2"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Villager3"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Villager4"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Woman"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate

# NPC templates
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Altar"}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
