package main

import (
	"GoProj1.0/protofiles"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":2333", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dial error:", err.Error())
	}
	defer conn.Close()

	client := protofiles.NewUserInfoServiceClient(conn)

	req := new(protofiles.UserRequest)
	req.Username = "micheal"

	resp, err := client.GetUserInfo(context.Background(), req)

	if err != nil {
		log.Fatal("resp error:", err.Error())
	}
	fmt.Printf("recevied: %v\n", resp)

}
