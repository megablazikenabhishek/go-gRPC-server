package main

import (
	"context"
	"fmt"
	"net"
	proto "starter/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	fmt.Println("Starting server at port 8080...........")

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) ServerReply(c context.Context, req *proto.HelloRequest) (res *proto.HelloResponse, err error) {
	fmt.Println("Request: ", req.GetSomeString())
	return &proto.HelloResponse{
		Response: "Hello Sir",
	}, nil
}
