package user

import (
	"context"

	"github.com/carefree/project/account/datamodel/user"
	"github.com/carefree/project/common/db"
	"github.com/carefree/project/common/objectid"

	pb "github.com/carefree/api/project/account/admin/user/v1"
)

func (s *Server) CreateNamespace(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
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
		Name:        user.FullName(p.req.Namespace, id),
		DisplayName: p.req.User.DisplayName,
	}
	p.resp, err = users.Create(n)
	if err != nil {
		return err
	}
	return nil
}
