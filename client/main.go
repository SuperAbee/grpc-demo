package main

import (
	"context"
	"log"
	"test/test"

	"test/transport"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("server1", grpc.WithInsecure(), grpc.WithContextDialer(transport.CustomeDialer))
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
