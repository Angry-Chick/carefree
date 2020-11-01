package user

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/user/v1"
)

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
