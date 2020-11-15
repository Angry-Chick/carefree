package namespace

import (
	"context"

	"github.com/carefree/project/home/datamodel/namespace"
	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/carefree/api/project/home/admin/namespace/v1"
)

// Server implement carefree Door service.
type Server struct {
	db *db.DB
}

// NewServer returns an Door service.
func NewServer(db *db.DB) *Server {
	return &Server{db: db}
}

// Register implement rpc service's Register method.
func (s *Server) Register(svr *rpc.Server) {
	pb.RegisterNamespaceAdminServer(svr.GRPC, s)
}

func (s *Server) DeleteNamespace(ctx context.Context, req *pb.DeleteNamespaceRequest) (*empty.Empty, error) {
	ns := namespace.New(s.db)
	err := ns.Purge(req.Name)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetNamespace(ctx context.Context, req *pb.GetNamespaceRequest) (*pb.Namespace, error) {
	panic("implement me")
}

func (s *Server) UpdateNamespace(ctx context.Context, req *pb.UpdateNamespaceRequest) (*pb.Namespace, error) {
	panic("implement me")
}
