package user

import (
	"context"

	pb "github.com/carefree/api/user/v1"
	"github.com/carefree/project/user/datamodel/user"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	rs := user.New(s.db)
	n, err := rs.Get(req.Name)
	if err != nil {
		return nil, err
	}
	return n, nil
}
