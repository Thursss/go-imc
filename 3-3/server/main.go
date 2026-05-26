package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type (
	GetUserRes struct {
		Id string `json:"id"`
	}
	GetUserReq struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}
)

type UserServer struct{}

func (s *UserServer) GetUser(res GetUserRes, req *GetUserReq) error {
	if u, ok := users[res.Id]; ok {
		*req = GetUserReq{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}
		return nil
	}

	return errors.New("没有找到用户")
}

func main() {
	userServer := new(UserServer)

	rpc.Register(userServer)
	lister, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("监听失败", err)
	}

	log.Println("服务启动成功")

	for {
		conn, err := lister.Accept()
		if err != nil {
			log.Fatal("接收客户端连接失败")
			continue
		}

		go rpc.ServeConn(conn)
	}
}
