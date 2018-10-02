package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/chrispaynes/gRPCrud/pkg/api/v1"
)

const apiVersion = "v1"

var log = logrus.New()

func main() {
	log.Formatter = &logrus.JSONFormatter{}

	conn, err := connectToServer()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := v1.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	pfx := t.Format(time.RFC3339Nano)

	createRequest := v1.CreateRequest{
		Api: apiVersion,
		Todo: &v1.Todo{
			Title:       "title (" + pfx + ")",
			Description: "description (" + pfx + ")",
			Reminder:    reminder,
		},
	}

	createResponse, err := c.Create(ctx, &createRequest)

	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	log.Printf("Create result: <%+v>\n\n", createResponse)

	// readAllRequest := v1.ReadAllRequest{
	// 	Api: apiVersion,
	// }

	// readAllResponse, err := c.ReadAll(ctx, &readAllRequest)

	// if err != nil {
	// 	log.Fatalf("ReadAll failed: %v", err)
	// }

	// log.Printf("ReadAll result: <%+v>\n\n", readAllResponse)
}

func connectToServer() (*grpc.ClientConn, error) {
	viper.SetConfigName("app")
	viper.AddConfigPath("$GOPATH/src/github.com/chrispaynes/gRPCrud/configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err.Error())
	}

	target := viper.GetString("server")

	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	return conn, nil
}
