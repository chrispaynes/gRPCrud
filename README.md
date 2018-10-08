# gRPCrud
Golang Client/Server gRPC CRUD App

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Requirements](#requirements)
- [Quickstart](#quickstart)
- [Architecture](#architecture)
- [Todo](#todo)
- [Credits](#credits)

## Overview

Basic Go Client/Server CRUD app using gRPC.

## Features

- [gRPC](https://grpc.io/docs/)
- [Protobuf Format and Protoc](https://github.com/protocolbuffers/protobuf)
- PostgreSQL
- [Dep](https://github.com/golang/dep)

## Requirements

- [Docker](https://www.docker.com/)
- [Docker-Compose](https://docs.docker.com/compose/install/)
- [Protocol Buffer Compiler (Protoc)](https://developers.google.com/protocol-buffers/docs/downloads)
    - [Binary Releases](https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.1)
    - [Installation Instructions](https://github.com/google/protobuf)

## Configuration

Runtime environment variables that need to be in your Shell.

| Parameter        | Example           | Description  |
| ------------- |-------------| -----|
| GOPATH | $HOME/$USER/go | |
| GOROOT | /usr/local/go | |
| GOBIN | $HOME/$USER/go/bin | |

---

Postgres configuration variables that need to be defined in `./docker/postgres/postgres.env`.

| Parameter        | Example           | Description  |
| ------------- |-------------| -----|
| POSTGRES_USER | read_write_user | |
| POSTGRES_PASSWORD| securepassword | |
| POSTGRES_DB | grpcrud | |

---

Postgres configuration variables that need to be defined in `./configs/app.toml`.
- Replace `0.0.0.0` with your Host IP using `make hostIP` or by manually editing the tile

| Parameter        | Example           | Description  |
| ------------- |-------------| -----|
| GRPC_PORT | "0.0.0.0.9090" | |
| POSTGRES_HOST | "0.0.0.0:5432" | |
| POSTGRES_DATABASE | "grpcrud" | |
| POSTGRES_USER_USERNAME | "read_write_user" | |
| POSTGRES_USER_PASSWORD | "securepassword" | |
| SERVER | "0.0.0.0:9090" | |

## Quickstart

#### 1. Download and Install The Protobuf Compiler (Protoc)
- Place the binary somewhere in your `$PATH`

```bash
$ which protoc

  /usr/local/bin/protoc
```

##### Copy The "Include" Files (Optional)
- If you intend to use the included common types (such as "google/protobuf/timestamp.proto"), copy the contents of the `include` directory into '/usr/local/include/'.
```bash
/usr/local/include
└── google
    └── protobuf
        ├── any.proto
        ├── api.proto
        ├── compiler
        │   └── plugin.proto
        ├── descriptor.proto
        ├── duration.proto
        ├── empty.proto
        ├── field_mask.proto
        ├── source_context.proto
        ├── struct.proto
        ├── timestamp.proto
        ├── type.proto
        └── wrappers.proto
```

---

#### 2. Compile the Proto Files

```bash
$ make protogen

./pkg
|── api
    └── v1
        └── todo-service.pb.go

```

---

#### 3. Start the Client, Server and Postgres Containers

```bash
$ docker-compose up
```

## Architecture
```

├── api
│   └── proto
│       └── v1
│           └── todo-service.proto  - V1 Proto Description
├── cmd
│   ├── client-grpc
│   │   └── main.go                 - gRPC client entrypoint
│   └── server
│       └── main.go                 - Server entrypoint
├── configs
│   └── app.toml                    - Client/Server environment variables
├── docker
│   └── postgres
│       ├── init.sql
│       ├── postgres.dockerfile
│       └── postgres.env            - PostgreSQL environment variables
├── docker-compose.yaml
├── Dockerfile
├── Gopkg.lock                      - Project's Dependency Graph Snapshot (Autogenerated by "dep ensure" and "dep init")
├── Gopkg.toml                      - Rule declarations to govern dep's management behavior
├── Makefile
├── pkg
│   ├── api
│   │   └── v1
│   │       └── todo-service.pb.go  - Compiled Protobuf TodoService Package (Autogenerated by protoc-gen-go)
│   ├── protocol
│   │   └── grpc
│   │       └── server.go           - Functionality to run the gRPC server
│   └── service
│       └── v1
│           └── todo-service.go     - CRUD functionality using compiled protobuf API and PB definitions
├── README.md
├── scripts
│   ├── clean-host-ip.sh            - Remove Host IP Address from app.toml configs
│   ├── gen-host-ip.sh              - Inject Host IP Address into app.toml configs
│   ├── protoc-gen.sh               - Compile Proto Files
│   ├── start-grpc-client.sh
│   └── start-grpc-server.sh
└── vendor
    ├── github.com
    │   ├── golang
    │   │   └── protobuf
    │   │       └── ptypes
    │   └── jmoiron
    │       └── sqlx
    ├── golang.org
    │   └── x
    │       └── net
    └── google.golang.org
        ├── genproto
        │   ├── googleapis
        │   │   └── rpc
        └── grpc

```

## Todo
- Add Remainder of CRUD Functions
- Add REST Server and Swagger Files
- Add Mock Data Creation to the PSQL Init Script

## Credits
- [Aleksandr Sokolovskii](https://github.com/amsokol) - [go-grpc-http-rest-microservice-tutorial](https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-daebb36a97e9)
