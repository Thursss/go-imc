package main

import (
	"context"
	"errors"
	"log"
	"net"
	"z/user/proto/user"

	"google.golang.org/grpc"
)

type UserServer struct {
}

func (s *UserServer) GetUser(ctx context.Context, res *user.GetUserRes) (*user.GetUserResp, error) {
	if u, ok := users[res.Id]; ok {
		return &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}, nil
	}

	return nil, errors.New("")
}

func main() {
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("失败")
	}

	s := grpc.NewServer()
	user.RegisterUserServer(s, new(UserServer))

	log.Println("已启动")

	s.Serve(listen)
}
