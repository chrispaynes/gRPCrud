SRC="api/proto/v1"
protoc --proto_path=$SRC --proto_path=vendor --go_out=plugins=grpc:pkg/api/v1 todo-service.proto