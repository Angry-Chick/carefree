package space

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/project/portal/space/v1"
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
	pb.RegisterSpaceServiceServer(svr.GRPC, s)
}

func (s *Server) CreateSpace(ctx context.Context, req *pb.CreateSpaceRequest) (*pb.Space, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpace not implemented")
}
func (s *Server) DeleteSpace(ctx context.Context, req *pb.DeleteSpaceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSpace not implemented")
}
func (s *Server) GetSpace(ctx context.Context, req *pb.GetSpaceRequest) (*pb.Space, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpace not implemented")
}
func (s *Server) UpdateSpace(ctx context.Context, req *pb.UpdateSpaceRequest) (*pb.Space, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSpace not implemented")
}
