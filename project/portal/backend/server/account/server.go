package account

import (
	"context"

	"github.com/carefree/net/rpc"
	"github.com/carefree/project/common/db"
	"github.com/carefree/project/common/objectid"
	"github.com/carefree/project/portal/datamodel/slice"
	"github.com/carefree/project/portal/datamodel/space"
	"github.com/carefree/project/portal/datamodel/user"
	"github.com/carefree/project/portal/integration/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pspb "github.com/carefree/api/project/portal/slice/v1"
	spb "github.com/carefree/api/project/portal/space/v1"
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

// SignUp 会为用户创建一个 portal 服务下的 user 资源，以及在 account 服务下创建一个用户，同时还会为用户初始化一个默认 space
// TODO(ljy)
func (s *Server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	proc := &signUpProc{req: req}
	if _, err := s.userCli.CreateUser(ctx, req.Username, req.Password); err != nil {
		return nil, err
	}
	if err := s.db.Transaction(proc.do); err != nil {
		return nil, err
	}
	return proc.resp, nil
}

type signUpProc struct {
	req  *pb.SignUpRequest
	resp *pb.SignUpResponse
}

func (p *signUpProc) do(db *db.DB) error {
	spaces := space.New(db)
	defaultID := objectid.MustNew().Base64()
	sp, err := spaces.Create(&spb.Space{
		Name:        space.FullName(defaultID),
		DisplayName: "default space",
		MySlice:     []string{defaultID},
	})
	if err != nil {
		return status.Errorf(codes.Unknown, "failed create space, err:%v", err)
	}
	slices := slice.New(db)
	_, err = slices.Create(&pspb.Slice{
		Name: slice.FullName(sp.Name, defaultID),
	})
	if err != nil {
		return status.Errorf(codes.Unknown, "failed create slice, err:%v", err)
	}
	users := user.New(db)
	_, err = users.Create(&upb.User{
		Name:        user.FullName(p.req.Username),
		DisplayName: p.req.Username,
		MySpaces:    []string{defaultID},
	})
	if err != nil {
		return err
	}
	p.resp = &pb.SignUpResponse{}
	return nil
}
