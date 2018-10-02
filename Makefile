.PHONY: cleanHostIP hostIP protogen restartGRPCClient restartGRPCServer unbindPort9090

cleanHostIP:
	@sh ./scripts/clean-host-ip.sh

protogen:
	@sh ./scripts/protoc-gen.sh

restartGRPCClient:
	@go run cmd/client-grpc/main.go

restartGRPCServer:
	@make unbindPort9090 && go run cmd/server/main.go

unbindPort9090:
	kill -9 $$(lsof -i :9090 | grep main | awk '{ print $$2}' | xargs) ||:

hostIP:
	@sh ./scripts/gen-host-ip.sh

clean:
	&sudo chmod -R 777 /home/afterlab/go/src/github.com/chrispaynes/gRPCrud/docker/pgdata ||: \
	&& docker rm grpc-client -f ||: \
	&& docker system prune -a ||: \
	&& docker-compose up grpc-client