package main

import (
	"log"
	"net"
	"test/test"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct{}

// func (g *GrpcServer) MyTest(ctx context.Context, req *test.Request) (*test.Response, error) {
// 	rsp := test.Response{
// 		B: req.A + "  rsp",
// 	}
// 	log.Printf("req %s, rsp %s", req.A, rsp.B)

// 	return &rsp, nil
// }

func (g *GrpcServer) MyTest(req *test.Request, ts test.Tester_MyTestServer) error {
	rsp := test.Response{
		B: req.A + "  rsp",
	}
	log.Printf("req %s, rsp %s", req.A, rsp.B)

	return ts.Send(&rsp)
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8810")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	test.RegisterTesterServer(grpcServer, &GrpcServer{})
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
