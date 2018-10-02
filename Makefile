.PHONY: protogen restartGRPCClient restartGRPCServer unbindPort9090

protogen:
	@sh ./scripts/protoc-gen.sh

restartGRPCClient:
	@go run cmd/client-grpc/main.go

restartGRPCServer:
	@make unbindPort9090 && go run cmd/server/main.go

unbindPort9090:
	kill -9 $$(lsof -i :9090 | grep main | awk '{ print $$2}' | xargs) ||: