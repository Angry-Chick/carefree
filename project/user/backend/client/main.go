package main

import (
	"context"
	"log"

	"github.com/carefree/server/rpc"

	pb "github.com/carefree/api/user/v1"
)

const (
	address = "localhost:50051"
)

func main() {
	ctx := context.Background()
	uc, err := rpc.Dial(ctx, address)
	if err != nil {
		log.Fatal(err)
	}
	defer uc.Close()

	cli := pb.NewUserServiceClient(uc)

	r, err := cli.GetUser(ctx, &pb.GetUserRequest{Name: "李俊毅"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Name: %v", r.GetName())
}
