SRC="api/proto/v1"
protoc -I=$SRC --go_out=pkg/api/v1 $SRC/todo-service.proto