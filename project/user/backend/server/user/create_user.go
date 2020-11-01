package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/user/v1"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
