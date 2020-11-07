package namespace

import (
	"context"

	"github.com/carefree/project/common/db"
	"github.com/carefree/project/door/datamodel/namespace"
	"github.com/carefree/net/rpc"
	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/carefree/api/door/v1/namespace"
)

// Server implement carefree namespace service.
type Server struct {
	db *db.DB
}

// NewServer returns an namespace service.
func NewServer(db *db.DB) *Server {
	return &Server{db: db}
}

// Register implement rpc service's Register method.
func (s *Server) Register(svr *rpc.Server) {
	pb.RegisterNamespaceServiceServer(svr.GRPC, s)
}

func (s *Server) DeleteNamespace(ctx context.Context, req *pb.DeleteNamespaceRequest) (*empty.Empty, error) {
	ns := namespace.New(s.db)
	if err := ns.Delete(req.Name); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetNamespace(ctx context.Context, req *pb.GetNamespaceRequest) (*pb.Namespace, error) {
	ns := namespace.New(s.db)
	n, err := ns.Get(req.Name)
	if err != nil {
		return nil, err
	}
	return n, nil
}
