package service

import (
	"GolandCode/grpc-demo/protobuf"
	"context"
)

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}

func (userService *userService) Login(ctx context.Context, req *protobuf.LoginRequest) (*protobuf.LoginResponse, error) {
	if req.Username == "admin" && req.Password == "123456" {
		resp := &protobuf.LoginResponse{
			Code:    10000,
			Message: "登录成功",
		}
		return resp, nil
	}

	resp := &protobuf.LoginResponse{
		Code:    10001,
		Message: "登录失败",
	}
	return resp, nil
}
