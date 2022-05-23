package main

import (
	"GolandCode/grpc-demo/client/auth"
	"GolandCode/grpc-demo/protobuf"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile("/Users/bn/GoProject/src/GolandCode/cert/server.pem", "xjy.bn.com")
	if err != nil {
		log.Fatal("get cert client by certificate is error ", err)
	}
	token := &auth.Authentication{
		Username: "admin",
		Password: "123456",
	}
	// 1. 打开gRPC服务端链接
	conn, err := grpc.Dial("127.0.0.1:3456", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(token))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// 2. 创建gRPC客户端
	client := protobuf.NewUserServiceClient(conn)

	// 3. 构造请求参数
	req := &protobuf.LoginRequest{
		Username: "admin",
		Password: "123456",
	}

	// 4. 调用服务端提供的服务
	response, err := client.Login(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Login Response: ", response)
}
