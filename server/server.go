package main

import (
	"GoProj1.0/models"
	"GoProj1.0/protofiles"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserInfoService struct{}

var u = UserInfoService{}

func (u *UserInfoService) GetUserInfo(ctx context.Context, req *protofiles.UserRequest) (resp *protofiles.UserResponse, err error) {
	username := req.Username

	user, e := models.UserGetByUsername(username)

	if e != nil {
		log.Fatal("读取数据出错", err.Error())
		return resp, e
	}
	resp = &protofiles.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
	}

	err = nil
	return resp, err
}

func main() {
	port := ":2333"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("listen error:", err.Error())
	}
	s := grpc.NewServer()

	protofiles.RegisterUserInfoServiceServer(s, &u)
	s.Serve(lis)
}
