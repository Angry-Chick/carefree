package account

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/account/datamodel/credential"
	"github.com/carefree/project/account/datamodel/user"
	"github.com/carefree/project/common/db"
	"github.com/carefree/project/common/jwt"
	"github.com/carefree/project/common/objectid"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	upb "github.com/carefree/api/project/account/admin/user/v1"
	acpb "github.com/carefree/api/project/account/type"
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
	credentials := credential.New(db)
	users := user.New(db)

	// TODO(ljy):暂时先实现只通过用户名方式, id 也暂时设置为用户名
	userName := user.FullName(p.req.GetNamespace(), p.req.GetUsername())
	u, err := users.Get(userName)
	if err != nil {
		return err
	}
	if u.Password != p.req.Password {
		return status.Error(codes.InvalidArgument, "invalid password")
	}
	expiryTime, err := ptypes.TimestampProto(jwt.IdentityTokenExpiry())
	if err != nil {
		return err
	}
	id, err := p.newCredentialID()
	if err != nil {
		return err
	}
	bc := &upb.BasicAuthCredential{
		User:     userName,
		Password: u.Password,
		Login: &upb.BasicAuthCredential_Username{
			Username: u.Username,
		},
	}
	cred := acpb.Credential{
		Name:       credential.FullName(id),
		ExpireTime: expiryTime,
		Payload: &acpb.Credential_BasicAuth{
			BasicAuth: bc,
		},
	}
	credentials.Create(&cred)
	p.resp = &atpb.Token{
		Opaque: id,
		Expiry: expiryTime,
		Type:   "Bearer",
	}
	return nil
}

func (p *basicAuthProc) newCredentialID() (string, error) {
	return objectid.Base64()
}
