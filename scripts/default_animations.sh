$ grpcurl -v -import-path ../../.. -proto cmd/admin/grpc/api.proto -d '' -plaintext localhost:8081 grpc.API/CreateAnimation
