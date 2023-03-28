#!/usr/bin/env bash

# Character templates
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "BlueNinja", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "BlueSamurai", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Boy", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Cavegirl", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Cavegirl2", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Caveman", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Caveman2", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Child", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "DarkNinja", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "EggBoy", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "EggGirl", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Eskimo", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "EskimoNinja", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "GoldKnight", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "GrayNinja", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Greenman", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "GreenNinja", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Inspector", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Knight", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Lion", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "MaskedNinja", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "MaskFrog", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Master", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Monk", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Monk2", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Noble", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "OldMan", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "OldMan2", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "OldMan3", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "OldWoman", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Princess", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "RedNinja", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "RedNinja2", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "RedSamurai", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Samurai", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Skeleton", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Villager", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Villager2", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Villager3", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Villager4", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Woman", "Entity": {"Objects": [{"Width": 16, "Height": 16}]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate

# NPC templates
# 224x288
# grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Altar", "Entity": {
# 	"Radius": 10,
# 	"Objects": [
# 		{"X": 41, "Y": 38, "Width": 20, "Height": 63},
# 		{"X": 159, "Y": 38, "Width": 20, "Height": 63},
# 		{"X": 41, "Y": 138, "Width": 20, "Height": 63},
# 		{"X": 159, "Y": 138, "Width": 20, "Height": 63},
# 		{"X": 61, "Y": 62, "Width": 98, "Height": 3},
# 		{"X": 16, "Y": 110, "Width": 3, "Height": 70},
# 		{"X": 63, "Y": 218, "Width": 1, "Height": 63},
# 		{"X": 157, "Y": 218, "Width": 1, "Height": 63}
# ]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate

grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/admin.proto -d '{"Name": "Altar", "Entity": {
	"Radius": 10,
	"Objects": [
		{"X": -55, "Y": -90, "Width": 20, "Height": 63},
		{"X": 63, "Y": -90, "Width": 20, "Height": 63},
		{"X": -55, "Y": 10, "Width": 20, "Height": 63},
		{"X": 63, "Y": 10, "Width": 20, "Height": 63},
		{"X": -35, "Y": -66, "Width": 98, "Height": 3},
		{"X": -80, "Y": -18, "Width": 3, "Height": 70},
		{"X": -80, "Y": -32, "Width": 24, "Height": 24},
		{"X": -80, "Y": 52, "Width": 24, "Height": 24},
		{"X": 107, "Y": -18, "Width": 3, "Height": 70},
		{"X": 83, "Y": -32, "Width": 24, "Height": 24},
		{"X": 83, "Y": 52, "Width": 24, "Height": 24},
		{"X": -33, "Y": 90, "Width": 1, "Height": 63},
		{"X": 61, "Y": 90, "Width": 1, "Height": 63}
]}}' -plaintext localhost:4282 grpc.Admin/CreateTemplate
