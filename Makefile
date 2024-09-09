PHONY: generate-structs
generate-structs:
	mkdir -p pkg/pack_v1
	protoc --go_out=pkg/pack_v1  --go_opt=paths=import \
			--go-grpc_out=pkg/pack_v1 --go-grpc_opt=paths=import \
			api/pack_v1/pack.proto
	mv pkg/pack_v1/github.com/cerhlhgr/grpc_example/pkg/pack_v1/* pkg/pack_v1/
	rm -rf pkg/pack_v1/github.com