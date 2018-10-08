#!/bin/bash
SRC="api/proto/v1"
OUT=pkg/api/v1

if [ ! -d "$OUT" ]
    then mkdir -p "$OUT"
fi

protoc --proto_path=$SRC --proto_path=vendor --go_out=plugins=grpc:$OUT todo-service.proto
