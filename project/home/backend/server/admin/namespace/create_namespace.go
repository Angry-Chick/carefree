package namespace

import (
	"context"

	"github.com/carefree/project/home/datamodel/namespace"
	"github.com/carefree/project/common/db"
	"github.com/carefree/project/common/objectid"

	pb "github.com/carefree/api/project/home/admin/namespace/v1"
)

func (s *Server) CreateNamespace(ctx context.Context, req *pb.CreateNamespaceRequest) (*pb.Namespace, error) {
	proc := &createProc{req: req}
	if err := proc.do(s.db); err != nil {
		return nil, err
	}
	return proc.resp, nil
}

type createProc struct {
	req  *pb.CreateNamespaceRequest
	resp *pb.Namespace
}

func (p *createProc) do(db *db.DB) error {
	namespaces := namespace.New(db)
	var err error
	id := p.req.Id
	if id == "" {
		if id, err = objectid.Base64(); err != nil {
			return err
		}
	} else if err := namespace.CheckID(id); err != nil {
		return err
	}
	n := &pb.Namespace{
		Name:          namespace.FullName(id),
		DisplayName:   p.req.Namespace.DisplayName,
		UserNamespace: p.req.Namespace.UserNamespace,
	}
	p.resp, err = namespaces.Create(n)
	if err != nil {
		return err
	}
	return nil
}
