package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "starter/proto"
)

func main() {
	// Setting up the connection
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := proto.NewExampleClient(conn)

	fmt.Println("Client Connection Success.............", client)

	time.Sleep(time.Second * 2)
	req := &proto.HelloRequest{
		SomeString: "Hello Abhishek Sir",
	}
	fmt.Println("Sending Server Reply..........")
	client.ServerReply(context.TODO(), req)
}
