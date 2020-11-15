package account

import (
	"context"
	"path"

	"google.golang.org/grpc"

	pb "github.com/carefree/api/project/account/admin/user/v1"
)

// UserClient is an user client.
type UserClient struct {
	usrCli pb.UserAdminClient
}

// NewUserClient returns a client instance.
func NewUserClient(cc *grpc.ClientConn) *UserClient {
	return &UserClient{
		usrCli: pb.NewUserAdminClient(cc),
	}
}

// GetUser returns an user by user id.
func (cli *UserClient) GetUser(ctx context.Context, ns, userID string) (*pb.User, error) {
	return cli.usrCli.GetUser(ctx, &pb.GetUserRequest{Name: fullUserName(ns, userID)})
}

// CreateUser create a user.
func (cli *UserClient) CreateUser(ctx context.Context, namespace string, username string, password string) (*pb.User, error) {
	return cli.usrCli.CreateUser(ctx, &pb.CreateUserRequest{
		Namespace: namespace,
		Id:        username,
		User: &pb.User{
			Username: username,
			Password: password,
		},
	})
}

func fullUserName(namespace, userID string) string {
	return path.Join(namespace, "users", userID)
}
