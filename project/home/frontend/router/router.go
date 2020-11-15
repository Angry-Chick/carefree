package router

import (
	"context"
	"net/http"
	"path"
	"path/filepath"

	npb "github.com/carefree/api/project/account/admin/namespace/v1"
	aupb "github.com/carefree/api/project/account/admin/user/v1"
	papb "github.com/carefree/api/project/account/v1"
	hpb "github.com/carefree/api/project/home/admin/home/v1"
	upb "github.com/carefree/api/project/home/admin/user/v1"
	hvpb "github.com/carefree/api/project/home/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type serviceName string

const (
	AccountService serviceName = "account"
	HomeService    serviceName = "home"
)

type ServiceConn map[serviceName]*grpc.ClientConn

var DefaultServiceConn = make(ServiceConn)

func (s ServiceConn) RegisterService(sn serviceName, conn *grpc.ClientConn) { s[sn] = conn }

type Router struct {
	*gin.Engine
	homeCli             hpb.HomeAdminClient
	homeUserCli         upb.UserAdminClient
	homeAccountCli      hvpb.AccountClient
	accountNamespaceCli npb.NamespaceAdminClient
	accountUserCli      aupb.UserAdminClient
	accountCli          papb.AccountClient
}

func New(sc ServiceConn) *Router {
	r := Router{
		Engine:              gin.Default(),
		homeCli:             hpb.NewHomeAdminClient(sc[HomeService]),
		homeUserCli:         upb.NewUserAdminClient(sc[HomeService]),
		homeAccountCli:      hvpb.NewAccountClient(sc[HomeService]),
		accountNamespaceCli: npb.NewNamespaceAdminClient(sc[AccountService]),
		accountUserCli:      aupb.NewUserAdminClient(sc[AccountService]),
		accountCli:          papb.NewAccountClient(sc[AccountService]),
	}
	return &r
}

var root = Resolve("../")

var buildPath = "carefree/project/home/frontend/build"

func (r *Router) RegisterHandle(ctx context.Context) {
	r.LoadHTMLFiles(path.Join(root, buildPath, "index.html"))
	r.Static("/static", path.Join(root, buildPath, "static"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Door"})
	})
	r.handleHomeAccount(ctx)
	r.handleAccountLogin(ctx)
	r.handleAccountUser(ctx)
}

func (r *Router) handleAccountLogin(ctx context.Context) {
	r.POST("/v1/namespace/:nid/login", func(c *gin.Context) {
		nid := c.Param("nid")
		username := c.PostForm("username")
		password := c.PostForm("password")
		tk, err := r.accountCli.BasicAuth(ctx, &papb.BasicAuthRequest{
			Namespace: path.Join("namespaces", nid),
			Login: &papb.BasicAuthRequest_Username{
				Username: username,
			},
			Password: password,
		})
		if err != nil {
			c.JSON(301, err.Error())
		}
		c.JSON(200, tk)
	})
}

func (r *Router) handleAccountUser(ctx context.Context) {
	r.GET("/v1/namespaces/:nid/users/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		nid := c.Param("nid")
		resp, err := r.accountUserCli.GetUser(ctx, &aupb.GetUserRequest{
			Name: path.Join("namespaces", nid, "users", uid),
		})
		if err != nil {
			c.JSON(301, err.Error())
		}
		c.JSON(200, resp)
	})
	r.DELETE("/v1/namespace/:nid/users/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		nid := c.Param("nid")
		resp, err := r.accountUserCli.DeleteUser(ctx, &aupb.DeleteUserRequest{
			Name: path.Join("namespace", nid, "users", uid),
		})
		if err != nil {
			c.JSON(301, err.Error())
		}
		c.JSON(200, resp)
	})
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Router) handleHomeAccount(ctx context.Context) {
	r.POST("/v1/namespaces/:nid/register", func(c *gin.Context) {
		nid := c.Param("nid")
		u := user{}
		c.BindJSON(&u)
		resp, err := r.homeAccountCli.SignUp(ctx, &hvpb.SignUpRequest{
			Namespace: path.Join("namespaces", nid),
			UserName:  u.Username,
			Password:  u.Password,
		})
		if err != nil {
			c.JSON(301, err.Error())
		}
		c.JSON(200, resp)
	})
}

func Resolve(p string) string {
	r, err := filepath.Abs(p)
	if err != nil {
		return p
	}
	return r
}
