package user

import (
	"context"

	pb "github.com/carefree/api/user/v1"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return &pb.User{Name: req.Name}, nil
}
