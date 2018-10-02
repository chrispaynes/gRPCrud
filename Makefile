.PHONY: clean cleanHostIP hostIP protogen restartGRPCClient restartGRPCServer unbindPort9090

cleanHostIP:
	@sh ./scripts/clean-host-ip.sh

protogen:
	@sh ./scripts/protoc-gen.sh

GRPCClient:
	@sh ./scripts/start-grpc-client.sh

GRPCServer:
	@sh ./scripts/start-grpc-server.sh

unbindPort9090:
	kill -9 $$(lsof -i :9090 | grep main | awk '{ print $$2}' | xargs) ||:

hostIP:
	@sh ./scripts/gen-host-ip.sh

clean:
	@sudo chmod -R 777 /home/afterlab/go/src/github.com/chrispaynes/gRPCrud/docker/pgdata ||: \
	&& docker rm grpc-client -f ||: \
	&& docker system prune -a ||: \
	&& docker-compose up grpc-client