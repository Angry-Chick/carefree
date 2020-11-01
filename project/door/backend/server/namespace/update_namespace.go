package namespace

import (
	"context"

	"github.com/carefree/project/common/db"
	"github.com/carefree/project/door/datamodel/namespace"

	pb "github.com/carefree/api/door/v1/namespace"
)

func (s *Server) UpdateNamespace(ctx context.Context, req *pb.UpdateNamespaceRequest) (*pb.Namespace, error) {
	proc := &updateProc{req: req}
	if err := s.db.Transaction(proc.do); err != nil {
		return nil, err
	}
	return proc.resp, nil
}

type updateProc struct {
	req  *pb.UpdateNamespaceRequest
	resp *pb.Namespace
}

func (p *updateProc) do(db *db.DB) error {
	namespaces := namespace.New(db)
	ns, err := namespaces.Get(p.req.Namespace.GetName())
	if err != nil {
		return err
	}
	ns.DisplayName = p.req.Namespace.DisplayName
	ns.Description = p.req.Namespace.Description
	p.resp, err = namespaces.Update(ns)
	if err != nil {
		return err
	}
	return nil
}
