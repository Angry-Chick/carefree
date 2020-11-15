package account

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/carefree/project/home/datamodel/namespace"
	"github.com/carefree/project/home/integration/account"
	"github.com/carefree/project/home/datamodel/user"

	pb "github.com/carefree/api/project/home/v1"
	upb "github.com/carefree/api/project/home/admin/user/v1"
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
	pb.RegisterAccountServer(svr.GRPC, s)
}

func (s *Server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	namespaces := namespace.New(s.db)
	ns, err := namespaces.Get(req.Namespace)
	if err != nil {
		return nil, err
	}
	users := user.New(s.db)
	if _, err := users.Create(&upb.User{
		Name:        user.FullName(req.Namespace, req.UserName),
		DisplayName: req.UserName,
	}); err != nil {
		return nil, err
	}
	if _, err = s.userCli.CreateUser(ctx, ns.UserNamespace, req.UserName, req.Password); err != nil {
		return nil, err
	}
	return &pb.SignUpResponse{}, nil
}
