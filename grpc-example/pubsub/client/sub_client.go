package main

import (
	"client/proto"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(
		context.Background(), &proto.String{Value: "golang:"},
	)
	if nil != err {
		log.Fatal(err)
	}

	go func() {
		for {
			reply, err := stream.Recv()
			if nil != err {
				if io.EOF == err {
					break
				}
				log.Fatal(err)
			}
			fmt.Println("sub1: ", reply.GetValue())
		}
	}()

	streamTopic, err := client.SubscribeTopic(
		context.Background(), &proto.String{Value: "golang:"},
	)
	if nil != err {
		log.Fatal(err)
	}

	go func() {
		for {
			reply, err := streamTopic.Recv()
			if nil != err {
				if io.EOF == err {
					break
				}
				log.Fatal(err)
			}
			fmt.Println("subTopic: ", reply.GetValue())
		}
	}()

	<-make(chan bool)
}
