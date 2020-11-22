package slice

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/project/portal/slice/v1"
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
	pb.RegisterSliceServiceServer(svr.GRPC, s)
}

func (s *Server) CreateSlice(ctx context.Context, req *pb.CreateSliceRequest) (*pb.Slice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSlice not implemented")
}

func (s *Server) UpdateSlice(ctx context.Context, req *pb.UpdateSliceRequest) (*pb.Slice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSlice not implemented")
}

func (s *Server) DeleteSlice(ctx context.Context, req *pb.DeleteSliceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSlice not implemented")
}

func (s *Server) GetSlice(ctx context.Context, req *pb.GetSliceRequest) (*pb.Slice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlice not implemented")
}
