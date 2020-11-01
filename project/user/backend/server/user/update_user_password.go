package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/user/v1"
)

func (s *Server) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.UpdateUserPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPassword not implemented")
}
