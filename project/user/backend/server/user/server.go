package user

import (
	"github.com/carefree/server/rpc"

	pb "github.com/carefree/api/user/v1"
)

// Server implement carefree user service.
type Server struct {
}

// NewServer returns an user service.
func NewServer() *Server {
	return &Server{}
}

// Register implement rpc service's Register method.
func (s *Server) Register(svr *rpc.Server) {
	pb.RegisterUserServiceServer(svr.GRPC, s)
}
