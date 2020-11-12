package namespace

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/account/datamodel/namespace"
	"github.com/carefree/project/common/db"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/project/account/admin/namespace/v1"
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
	pb.RegisterNamespaceAdminServer(svr.GRPC, s)
}

func (s *Server) DeleteNamespace(ctx context.Context, req *pb.DeleteNamespaceRequest) (*empty.Empty, error) {
	rs := namespace.New(s.db)
	err := rs.Purge(req.Name)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetNamespace(ctx context.Context, req *pb.GetNamespaceRequest) (*pb.Namespace, error) {
	rs := namespace.New(s.db)
	n, err := rs.Get(req.Name)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (s *Server) UpdateNamespace(ctx context.Context, req *pb.UpdateNamespaceRequest) (*pb.Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNamespace not implemented")
}
