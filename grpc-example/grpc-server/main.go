package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"grpc-server/proto"
	"io"
	"io/ioutil"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

// Token 认证
func Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var user string
	var password string

	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}

	if user != "admin" || password != "admin" {
		return grpc.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}

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
	// server := grpc.NewServer()

	// Token 认证
	// var authInterceptor grpc.UnaryServerInterceptor
	// authInterceptor = func(
	// 	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	// ) (resp interface{}, err error) {
	// 	//拦截普通方法请求，验证 Token
	// 	err = Auth(ctx)
	// 	if err != nil {
	// 		return
	// 	}
	// 	// 继续处理请求
	// 	return handler(ctx, req)
	// }
	// server := grpc.NewServer(grpc.UnaryInterceptor(interceptor))

	// 证书认证-单向认证
	// creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server.key")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
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
		ClientAuth: tls.RequireAndVerifyClientCert,
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})

	// 增加验证器，分别是标准模式和流模式
	server := grpc.NewServer(grpc.Creds(creds),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// Token 认证
				// authInterceptor,
				grpc_validator.UnaryServerInterceptor(),
			),
		),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_validator.StreamServerInterceptor(),
			),
		),
	)
	// 注册 grpcurl 所需的 reflection 服务
	reflection.Register(server)
	// 注册业务服务
	proto.RegisterGreeterServer(server, &greeter{})

	fmt.Println("grpc server start ...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
