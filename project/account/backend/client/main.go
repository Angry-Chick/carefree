package main

import (
	"context"
	"github.com/carefree/net/rpc"
	"github.com/carefree/project/account/datamodel/namespace"

	pb "github.com/carefree/api/project/account/admin/namespace/v1"
)

const (
	accountEndpoint = "localhost:9091"
)

func main() {
	ctx := context.Background()
	ac, err := rpc.Dial(ctx, accountEndpoint)
	if err != nil {
		panic(err)
	}
	acli := pb.NewNamespaceAdminClient(ac)
	_, err = acli.CreateNamespace(ctx, &pb.CreateNamespaceRequest{
		Id: "carefree",
		Namespace: &pb.Namespace{
			Name:        namespace.FullName("carefree"),
			DisplayName: "carefree namespace",
			Description: "carefree namespace",
		},
	})
	if err != nil {
		panic(err)
	}
}
