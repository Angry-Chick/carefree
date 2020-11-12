package home

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/project/home/admin/home/v1"
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
	pb.RegisterHomeAdminServer(svr.GRPC, s)
}

func (s *Server) CreateHome(ctx context.Context, req *pb.CreateHomeRequest) (*pb.Home, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHome not implemented")
}
func (s *Server) DeleteHome(ctx context.Context, req *pb.DeleteHomeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHome not implemented")
}
func (s *Server) GetHome(ctx context.Context, req *pb.GetHomeRequest) (*pb.Home, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHome not implemented")
}
func (s *Server) UpdateHome(ctx context.Context, req *pb.UpdateHomeRequest) (*pb.Home, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHome not implemented")
}
