package main

import (
	"context"
	"fmt"
	"grpc-client/proto"
	"log"

	"google.golang.org/grpc"
)

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 简单调用
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	auth := Authentication{
		User:     "admin",
		Password: "admin",
	}

	// Token 认证
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
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
