package main

import (
	"GolandCode/grpc-demo/protobuf"
	"GolandCode/grpc-demo/service"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {

	creds, err := credentials.NewServerTLSFromFile("/Users/bn/GoProject/src/GolandCode/cert/server.pem", "/Users/bn/GoProject/src/GolandCode/cert/server.key")
	if err != nil {
		log.Fatal("generate certificate is error ", err)
	}

	var authInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = Auth(ctx)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(authInterceptor))
	protobuf.RegisterUserServiceServer(grpcServer, service.NewUserService())

	fmt.Println("gRPC is running...")

	fmt.Println("Now begin listen port 3456")
	listener, err := net.Listen("tcp", "127.0.0.1:3456")
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("gRPC server err:%s\n", err)
	}
}

func Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing verify info")
	}
	var username, password string
	if value, ok := md["username"]; ok {
		username = value[0]
	}
	if value, ok := md["password"]; ok {
		password = value[0]
	}
	if username != "admin" || password != "123456" {
		return status.Errorf(codes.Unauthenticated, "username and password is error")
	}
	return nil
}
