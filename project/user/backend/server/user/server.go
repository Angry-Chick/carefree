package user

import (
	"github.com/carefree/project/common/db"
	"github.com/carefree/server/rpc"

	pb "github.com/carefree/api/user/v1"
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
	pb.RegisterUserServiceServer(svr.GRPC, s)
}
