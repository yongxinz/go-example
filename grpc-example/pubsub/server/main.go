package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"server/proto"
	"strings"
	"time"

	"github.com/moby/moby/pkg/pubsub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func (p *PubsubService) Publish(ctx context.Context, arg *proto.String) (*proto.String, error) {
	p.pub.Publish(arg.GetValue())
	return &proto.String{}, nil
}

func (p *PubsubService) SubscribeTopic(arg *proto.String, stream proto.PubsubService_SubscribeTopicServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&proto.String{Value: v.(string)}); nil != err {
			return err
		}
	}
	return nil
}

func (p *PubsubService) Subscribe(arg *proto.String, stream proto.PubsubService_SubscribeServer) error {
	ch := p.pub.Subscribe()

	for v := range ch {
		if err := stream.Send(&proto.String{Value: v.(string)}); nil != err {
			return err
		}
	}
	return nil
}

func NewPubsubService() *PubsubService {
	return &PubsubService{pub: pubsub.NewPublisher(100*time.Millisecond, 10)}
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
	proto.RegisterPubsubServiceServer(server, NewPubsubService())

	fmt.Println("grpc server start ...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
