package main

import (
	"client/proto"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	// 简单调用
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(3*time.Second)))
	defer cancel()

	client := proto.NewGreeterClient(conn)
	// 简单调用
	reply, err := client.SayHello(ctx, &proto.HelloRequest{Name: "zzz"})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalln("client.SayHello err: deadline")
			}
		}

		log.Fatalf("client.SayHello err: %v", err)
	}
	fmt.Println(reply.Message)
}
