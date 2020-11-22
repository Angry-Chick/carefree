package account

import (
	"context"
	"path"

	"google.golang.org/grpc"

	pb "github.com/carefree/api/project/account/user/v1"
)

// UserClient is an user client.
type UserClient struct {
	usrCli pb.UserServicesClient
}

// NewUserClient returns a client instance.
func NewUserClient(cc *grpc.ClientConn) *UserClient {
	return &UserClient{
		usrCli: pb.NewUserServicesClient(cc),
	}
}

// GetUser returns an user by user id.
func (cli *UserClient) GetUser(ctx context.Context, userID string) (*pb.User, error) {
	return cli.usrCli.GetUser(ctx, &pb.GetUserRequest{Name: path.Join("users", userID)})
}

// CreateUser create a user.
func (cli *UserClient) CreateUser(ctx context.Context, username string, password string) (*pb.User, error) {
	return cli.usrCli.CreateUser(ctx, &pb.CreateUserRequest{
		User: &pb.User{
			Username: username,
			Password: password,
		},
	})
}
