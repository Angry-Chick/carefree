package router

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	aupb "github.com/carefree/api/project/account/user/v1"
	papb "github.com/carefree/api/project/account/v1"
	pspb "github.com/carefree/api/project/portal/slice/v1"
	spb "github.com/carefree/api/project/portal/space/v1"
	upb "github.com/carefree/api/project/portal/user/v1"
	pvpb "github.com/carefree/api/project/portal/v1"
)

type serviceName string

const (
	AccountService serviceName = "account"
	PortalService  serviceName = "portal"
)

type ServiceConn map[serviceName]*grpc.ClientConn

var DefaultServiceConn = make(ServiceConn)

func (s ServiceConn) RegisterService(sn serviceName, conn *grpc.ClientConn) { s[sn] = conn }

type Router struct {
	*gin.Engine

	portalCli      pvpb.PortalServiceClient
	portalUserCli  upb.UserServiceClient
	portalSliceCli pspb.SliceServiceClient
	portalSpaceCli spb.SpaceServiceClient

	accountUserCli aupb.UserServicesClient
	accountCli     papb.AccountServiceClient
}

func New(sc ServiceConn) *Router {
	r := Router{
		Engine: gin.Default(),

		portalCli:      pvpb.NewPortalServiceClient(sc[PortalService]),
		portalUserCli:  upb.NewUserServiceClient(sc[PortalService]),
		portalSliceCli: pspb.NewSliceServiceClient(sc[PortalService]),
		portalSpaceCli: spb.NewSpaceServiceClient(sc[PortalService]),

		accountUserCli: aupb.NewUserServicesClient(sc[AccountService]),
		accountCli:     papb.NewAccountServiceClient(sc[AccountService]),
	}
	return &r
}

func (r *Router) RegisterHandle() {
	r.POST("/:service/:method", func(c *gin.Context) {
		var req interface{}
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		data, err := json.Marshal(req)
		if err != nil {
			return
		}
		resp, err := r.handle(c, c.Param("service"), c.Param("method"), data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(200, resp)
	})
}

type service interface {
	handle(ctx context.Context) (proto.Message, error)
}

func (r *Router) handle(ctx context.Context, sn string, mn string, req []byte) (proto.Message, error) {
	var sv service
	switch sn {
	case "carefree.project.account.v1.Account":
		sv = &accountService{cli: r.accountCli, req: req, method: mn}
	case "carefree.project.portal.slice.v1.SliceService":
		sv = &portalSliceService{cli: r.portalSliceCli, req: req, method: mn}
	case "carefree.project.portal.v1.PortalService":
		sv = &portalService{cli: r.portalCli, req: req, method: mn}
	case "carefree.project.portal.space.v1.SpaceService":
		sv = &portalSpaceService{cli: r.portalSpaceCli, req: req, method: mn}
	case "carefree.project.portal.user.v1.UserService":
		sv = &portalUserService{cli: r.portalUserCli, req: req, method: mn}
	default:
		return nil, errors.New("unknown service name")
	}
	return sv.handle(ctx)
}
