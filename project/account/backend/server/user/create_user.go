package user

import (
	"context"

	"github.com/carefree/project/account/datamodel/user"
	"github.com/carefree/project/common/db"

	pb "github.com/carefree/api/project/account/user/v1"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	proc := &createProc{req: req}
	if err := proc.do(s.db); err != nil {
		return nil, err
	}
	return proc.resp, nil
}

type createProc struct {
	req  *pb.CreateUserRequest
	resp *pb.User
}

func (p *createProc) do(db *db.DB) error {
	users := user.New(db)
	var err error
	n := &pb.User{
		// 未来使用单独的 user id
		Name:     user.FullName(p.req.User.Username),
		Username: p.req.User.Username,
		Password: p.req.User.Password,
	}
	p.resp, err = users.Create(n)
	if err != nil {
		return err
	}
	return nil
}
