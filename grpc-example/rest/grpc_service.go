package main

import (
	"context"
	"net"

	"rest/proto"

	"google.golang.org/grpc"
)

type RestServiceImpl struct{}

func (r *RestServiceImpl) Get(ctx context.Context, message *proto.StringMessage) (*proto.StringMessage, error) {
	return &proto.StringMessage{Value: "Get hi:" + message.Value + "#"}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *proto.StringMessage) (*proto.StringMessage, error) {
	return &proto.StringMessage{Value: "Post hi:" + message.Value + "@"}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	proto.RegisterRestServiceServer(grpcServer, new(RestServiceImpl))
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer.Serve(lis)
}
