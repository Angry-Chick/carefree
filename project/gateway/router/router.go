package router

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

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

func (r *Router) RegisterHandle(ctx context.Context) {
	r.POST("/:service/:method", func(c *gin.Context) {
		var req interface{}
		if err := c.BindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "err:%s", err)
		}
		data, err := json.Marshal(req)
		if err != nil {
			c.String(http.StatusBadRequest, "err:%s", err)
		}
		svr, err := r.generateService(c.Param("service"), data)
		if err != nil {
			c.String(http.StatusBadRequest, "err:%s", err)
		}
		resp, err := svr.handle(ctx, c.Param("method"))
		if err != nil {
			c.String(http.StatusBadRequest, "err:%s", err)
		}
		c.String(200, "%s", string(resp))
	})
}

type service interface {
	handle(ctx context.Context, method string) (resp []byte, err error)
}

func (r *Router) generateService(sn string, req []byte) (service, error) {
	switch sn {
	case "carefree.project.account.v1.Account":
		return &accountService{
			cli: r.accountCli,
			req: req,
		}, nil
	case "carefree.project.home.room.v1.SliceService":
		return &spaceSliceService{
			cli: r.portalSliceCli,
			req: req,
		}, nil
	case "carefree.project.portal.v1.PortalService":
		return &portalService{
			cli: r.portalCli,
			req: req,
		}, nil
	}
	return nil, errors.New("unknown service name")
}
