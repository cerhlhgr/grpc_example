PHONY: generate-structs
generate-structs:
	mkdir -p pkg/pack_v1
	protoc --go_out=pkg/pack_v1  --go_opt=paths=import \
			--go-grpc_out=pkg/pack_v1 --go-grpc_opt=paths=import \
			api/pack_v1/pack.proto