package main

import (
	"context"
	"github.com/carefree/net/rpc"
	"github.com/carefree/project/home/datamodel/namespace"

	pb "github.com/carefree/api/project/home/admin/namespace/v1"
)

const (
	homeEndpoint = "localhost:9090"
)

func main() {
	ctx := context.Background()
	hc, err := rpc.Dial(ctx, homeEndpoint)
	if err != nil {
		panic(err)
	}
	acli := pb.NewNamespaceAdminClient(hc)
	_, err = acli.CreateNamespace(ctx, &pb.CreateNamespaceRequest{
		Id: "carefree",
		Namespace: &pb.Namespace{
			Name:          namespace.FullName("carefree"),
			DisplayName:   "carefree namespace",
			Description:   "carefree namespace",
			UserNamespace: "carefree",
		},
	})
	if err != nil {
		panic(err)
	}
}
