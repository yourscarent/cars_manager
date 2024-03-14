proto:
	protoc \
	--proto_path=pkg/proto \
	--go-grpc_out=./pkg/proto/gen \
	--go_out=./pkg/proto/gen \
	--go-grpc_opt=paths=source_relative \
	--go_opt=paths=source_relative \
	manager.proto

container:
	docker-compose -f docker-compose-local.yaml up

app:
	go run cmd/main.go -c cmd/config.local.yaml