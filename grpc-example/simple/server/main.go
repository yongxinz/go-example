package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type greeter struct {
}

func (*greeter) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println(req)
	reply := &proto.HelloReply{Message: "hello"}
	return reply, nil
}

func (*greeter) SayHelloStream(stream proto.Greeter_SayHelloStreamServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		fmt.Println("Recv: " + args.Name)
		reply := &proto.HelloReply{Message: "hi " + args.Name}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 简单调用
	server := grpc.NewServer()
	// 注册 grpcurl 所需的 reflection 服务
	reflection.Register(server)
	// 注册业务服务
	proto.RegisterGreeterServer(server, &greeter{})

	fmt.Println("grpc server start ...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
