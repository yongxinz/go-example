package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"grpc-client/proto"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	// auth := Authentication{
	// 	User:     "admin",
	// 	Password: "admin",
	// }

	// Token 认证
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 证书认证-单向认证
	// creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "example.grpcdev.cn")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "www.example.grpcdev.cn",
		RootCAs:    certPool,
	})

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
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
