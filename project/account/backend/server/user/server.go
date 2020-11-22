package user

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/account/datamodel/user"
	"github.com/carefree/project/common/db"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/project/account/user/v1"
)

// Server implement carefree user service.
type Server struct {
	db *db.DB
}

// NewServer returns an user service.
func NewServer(db *db.DB) *Server {
	return &Server{db: db}
}

// Register implement rpc service's Register method.
func (s *Server) Register(svr *rpc.Server) {
	pb.RegisterUserServicesServer(svr.GRPC, s)
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*empty.Empty, error) {
	rs := user.New(s.db)
	err := rs.Purge(req.Name)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	rs := user.New(s.db)
	n, err := rs.Get(req.Name)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
