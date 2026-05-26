package main

import (
	"log"
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

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("建立连接失败", err)
	}
	defer client.Close()

	var (
		req = &GetUserReq{
			Id: "4",
		}
		res GetUserReq
	)

	err = client.Call("UserServer.GetUser", req, &res)
	if err != nil {
		log.Fatal("请求失败", err)
	}

	log.Println(res)
}
