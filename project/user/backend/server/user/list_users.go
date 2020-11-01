package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/user/v1"
)

func (s *Server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
