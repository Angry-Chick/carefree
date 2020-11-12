package account

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/carefree/api/project/account/v1"
	atpb "github.com/carefree/api/project/type/accesstoken"
)

// Server implement carefree account service.
type Server struct {
	db *db.DB
}

// NewServer returns an account service.
func NewServer(db *db.DB) *Server {
	return &Server{db: db}
}

// Register implement rpc service's Register method.
func (s *Server) Register(svr *rpc.Server) {
	pb.RegisterAccountServer(svr.GRPC, s)
}

func (s *Server) BasicAuth(context.Context, *pb.BasicAuthRequest) (*atpb.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BasicAuth not implemented")
}
