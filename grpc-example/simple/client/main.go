package main

import (
	"client/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// 简单调用
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	// 简单调用
	reply, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "zzz"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.Message)

	// 流处理
	// stream, err := client.SayHelloStream(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 发送消息
	// go func() {
	// 	for {
	// 		if err := stream.Send(&proto.HelloRequest{Name: "zhangsan"}); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	// // 接收消息
	// for {
	// 	reply, err := stream.Recv()
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(reply.Message)
	// }
}
