package main

import (
	"GolandCode/grpc-demo/protobuf"
	"GolandCode/grpc-demo/service"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:3456")
	if err != nil {
		panic(err)
	}
	fmt.Println("Now begin listen port 3456")
	grpcServer := grpc.NewServer()

	var userService service.UserService
	userService = service.NewUserService()

	protobuf.RegisterUserServiceServer(grpcServer, userService)

	fmt.Println("gRPC is running...")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("gRPC server err:%s\n", err)
	}
}
