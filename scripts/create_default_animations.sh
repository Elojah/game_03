grpcurl -v -import-path ../../.. -proto cmd/api/grpc/api.proto -d @ -plaintext localhost:4200 grpc.API/CreateAnimation < scripts/animation/idle.json
grpcurl -v -import-path ../../.. -proto cmd/api/grpc/api.proto -d @ -plaintext localhost:4200 grpc.API/CreateAnimation < scripts/animation/up.json
grpcurl -v -import-path ../../.. -proto cmd/api/grpc/api.proto -d @ -plaintext localhost:4200 grpc.API/CreateAnimation < scripts/animation/down.json
grpcurl -v -import-path ../../.. -proto cmd/api/grpc/api.proto -d @ -plaintext localhost:4200 grpc.API/CreateAnimation < scripts/animation/left.json
grpcurl -v -import-path ../../.. -proto cmd/api/grpc/api.proto -d @ -plaintext localhost:4200 grpc.API/CreateAnimation < scripts/animation/right.json
