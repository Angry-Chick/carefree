package namespace

import (
	"context"

	"github.com/carefree/project/common/db"
	"github.com/carefree/project/common/objectid"
	"github.com/carefree/project/door/datamodel/namespace"

	pb "github.com/carefree/api/door/v1/namespace"
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
		Name:        namespace.FullName(id),
		DisplayName: p.req.Namespace.DisplayName,
		Description: p.req.Namespace.Description,
	}
	p.resp, err = namespaces.Create(n)
	if err != nil {
		return err
	}
	return nil
}
