package main

import (
	"context"
	"fmt"
	"log"

	"github.com/carefree/server/rpc"

	pb "github.com/carefree/api/door/v1/namespace"
)

const (
	address = "localhost:9090"
)

func main() {
	ctx := context.Background()
	rn, err := rpc.Dial(ctx, address)
	if err != nil {
		log.Fatal(err)
	}
	defer rn.Close()

	cli := pb.NewNamespaceServiceClient(rn)
	n, err := cli.CreateNamespace(ctx, &pb.CreateNamespaceRequest{
		Id: "test5",
		Namespace: &pb.Namespace{
			DisplayName: "test display name",
			Description: "test description",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	// gn, err := cli.GetNamespace(ctx, &pb.GetNamespaceRequest{
	// 	Name: namespace.FullName("test3"),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(gn)

	// gnn, err := cli.UpdateNamespace(ctx, &pb.UpdateNamespaceRequest{
	// Namespace: &pb.Namespace{

	// 		Name:        namespace.FullName("test1"),
	// 		DisplayName: "hahahahah",
	// 		Description: "jjjjjjj",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(gnn)
}
