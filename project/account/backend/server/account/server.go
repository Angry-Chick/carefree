package account

import (
	"context"
	"time"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/account/datamodel/user"
	"github.com/carefree/project/common/db"
	"github.com/carefree/project/common/jwt"
	jg "github.com/dgrijalva/jwt-go"
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
	pb.RegisterAccountServiceServer(svr.GRPC, s)
}

func (s *Server) BasicAuth(ctx context.Context, req *pb.BasicAuthRequest) (*atpb.Token, error) {
	proc := &basicAuthProc{req: req}
	if err := s.db.Transaction(proc.do); err != nil {
		return nil, err
	}
	return proc.resp, nil
}

type basicAuthProc struct {
	req  *pb.BasicAuthRequest
	resp *atpb.Token
}

func (p *basicAuthProc) do(db *db.DB) error {
	users := user.New(db)

	// TODO(ljy):初期暂时先实现只通过用户名方式, id 也暂时设置为用户名。
	// 未来可以通过邮箱和手机号的形式登录
	// 用户名和用户 ID 为不同的值，通过关联表通过用户名来找到 user id
	un := user.FullName(p.req.GetUsername())
	u, err := users.Get(un)
	if err != nil {
		return status.Error(codes.NotFound, "user not found")
	}
	if u.Password != p.req.Password {
		return status.Error(codes.InvalidArgument, "error password")
	}
	expiresAt := jwt.IdentityTokenExpiry().Unix()
	claims := jwt.Claims{
		User: p.req.GetUsername(),
		StandardClaims: jg.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt,
			Issuer:    "carefree.account",
		},
	}
	tk, err := jwt.New().GenerateToken(&claims)
	if err != nil {
		return err
	}
	p.resp = &atpb.Token{
		Opaque: tk,
		Type:   "Bearer",
		Expiry: expiresAt,
	}
	return nil
}
