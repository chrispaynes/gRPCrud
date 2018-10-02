package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chrispaynes/gRPCrud/pkg/protocol/grpc"
	"github.com/chrispaynes/gRPCrud/pkg/service/v1"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var log = logrus.New()

func main() {
	log.Formatter = &logrus.JSONFormatter{}

	if err := runServer(); err != nil {
		log.Errorf("%v\n", err)
		os.Exit(1)
	}
}

// RunServer runs gRPC server and HTTP gateway
func runServer() error {
	ctx := context.Background()

	viper.SetConfigName("app")
	viper.AddConfigPath("$GOPATH/src/github.com/chrispaynes/gRPCrud/configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config file: %s", err.Error())
	}

	grpcPort := viper.GetString("GRPC_PORT")

	if len(grpcPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", grpcPort)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		viper.GetString("POSTGRES_USER_USERNAME"),
		viper.GetString("POSTGRES_USER_PASSWORD"),
		viper.GetString("POSTGRES_HOST"),
		viper.GetString("POSTGRES_DATABASE"),
	)

	log.Info("CONNECTED TO POSTGRES")
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	log.Info("Creating new TODO service server")
	v1API := v1.NewTodoServiceServer(db)

	log.Info("grpc.RunServer ")

	return grpc.RunServer(ctx, v1API, grpcPort)
}
