package room

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/project/home/room/v1"
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
	pb.RegisterRoomServiceServer(svr.GRPC, s)
}

func (s *Server) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}

func (s *Server) UpdateRoom(ctx context.Context, req *pb.UpdateRoomRequest) (*pb.Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}

func (s *Server) DeleteRoom(ctx context.Context, req *pb.DeleteRoomRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}

func (s *Server) GetRoom(ctx context.Context, req *pb.GetRoomRequest) (*pb.Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
