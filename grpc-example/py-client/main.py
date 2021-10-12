import grpc

import helloworld_pb2
import helloworld_pb2_grpc


def main():
    channel = grpc.insecure_channel("127.0.0.1:50051")
    stub = helloworld_pb2_grpc.GreeterStub(channel)
    response = stub.SayHello(helloworld_pb2.HelloRequest(name="zhangsan"))
    print(response.message)


if __name__ == '__main__':
    main()
