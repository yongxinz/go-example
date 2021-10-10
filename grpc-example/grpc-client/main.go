package main

import (
	"context"
	"fmt"
	"grpc-client/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "zhangsan"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.Message)
}
