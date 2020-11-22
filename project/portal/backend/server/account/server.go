package account

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/carefree/project/portal/datamodel/user"
	"github.com/carefree/project/portal/integration/account"

	upb "github.com/carefree/api/project/portal/user/v1"
	pb "github.com/carefree/api/project/portal/v1"
)

// Server implement carefree Door service.
type Server struct {
	userCli *account.UserClient
	db      *db.DB
}

// NewServer returns an Door service.
func NewServer(db *db.DB, userCli *account.UserClient) *Server {
	return &Server{db: db, userCli: userCli}
}

// Register implement rpc service's Register method.
func (s *Server) Register(svr *rpc.Server) {
	pb.RegisterPortalServiceServer(svr.GRPC, s)
}

func (s *Server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	users := user.New(s.db)
	if _, err := users.Create(&upb.User{
		Name:        user.FullName(req.Username),
		DisplayName: req.Username,
	}); err != nil {
		return nil, err
	}
	if _, err := s.userCli.CreateUser(ctx, req.Username, req.Password); err != nil {
		return nil, err
	}
	return &pb.SignUpResponse{}, nil
}
