package main

import (
	"context"
	"log"
	"z/user/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("失败")
	}
	defer client.Close()

	c := user.NewUserClient(client)
	res, err := c.GetUser(context.Background(), &user.GetUserRes{
		Id: "1",
	})

	if err != nil {
		log.Fatalf("失败")
	}

	log.Println(res)
}
