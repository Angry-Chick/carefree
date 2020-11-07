package user

import (
	"context"

	"github.com/carefree/project/common/db"
	"github.com/carefree/project/common/objectid"
	"github.com/carefree/project/user/datamodel/user"

	pb "github.com/carefree/api/user/v1"
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
	id := p.req.Id
	if id == "" {
		if id, err = objectid.Base64(); err != nil {
			return err
		}
	} else if err := user.CheckID(id); err != nil {
		return err
	}
	n := &pb.User{
		Name:        user.FullName(id),
		DisplayName: p.req.User.DisplayName,
		Password:    p.req.User.Password,
	}
	p.resp, err = users.Create(n)
	if err != nil {
		return err
	}
	return nil
}
