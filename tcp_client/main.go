package main

import (
	"context"
	"log"
	"test/test"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8810", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	grpcClient := test.NewTesterClient(conn)

	request := test.Request{
		A: "a",
	}

	response, err := grpcClient.MyTest(context.Background(), &request)
	if err != nil {
		panic(err)
	}
	log.Println(response.Recv())
}
