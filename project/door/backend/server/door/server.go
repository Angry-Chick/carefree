package door

import (
	"context"

	"github.com/carefree/project/common/db"
	"github.com/carefree/server/rpc"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/door/v1/door"
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
	pb.RegisterDoorServiceServer(svr.GRPC, s)
}

func (s *Server) DeleteDoor(ctx context.Context, req *pb.DeleteDoorRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDoor not implemented")
}
func (s *Server) GetDoor(ctx context.Context, req *pb.GetDoorRequest) (*pb.Door, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoor not implemented")
}
func (s *Server) ListDoors(ctx context.Context, req *pb.ListDoorsRequest) (*pb.ListDoorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDoors not implemented")
}
