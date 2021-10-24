package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"runtime"
	"server/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type greeter struct {
}

func (*greeter) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	data := make(chan *proto.HelloReply, 1)
	go handle(ctx, req, data)
	select {
	case res := <-data:
		return res, nil
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
}

func handle(ctx context.Context, req *proto.HelloRequest, data chan<- *proto.HelloReply) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		runtime.Goexit() //超时后退出该Go协程
	case <-time.After(4 * time.Second): // 模拟耗时操作
		res := proto.HelloReply{
			Message: "hello " + req.Name,
		}
		// //修改数据库前进行超时判断
		// if ctx.Err() == context.Canceled{
		// 	...
		// 	//如果已经超时，则退出
		// }
		data <- &res
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
