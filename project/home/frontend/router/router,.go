package router

import (
	"context"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	npb "github.com/carefree/api/project/account/admin/namespace/v1"
	aupb "github.com/carefree/api/project/account/admin/user/v1"
	hpb "github.com/carefree/api/project/home/admin/home/v1"
	upb "github.com/carefree/api/project/home/admin/user/v1"
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
	accountNamespaceCli npb.NamespaceAdminClient
	accountUserCli      aupb.UserAdminClient
}

func New(sc ServiceConn) *Router {
	r := Router{
		Engine:              gin.Default(),
		homeCli:             hpb.NewHomeAdminClient(sc[HomeService]),
		homeUserCli:         upb.NewUserAdminClient(sc[HomeService]),
		accountNamespaceCli: npb.NewNamespaceAdminClient(sc[AccountService]),
		accountUserCli:      aupb.NewUserAdminClient(sc[AccountService]),
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
	// r.handleHome(ctx)
	// r.handleAccountNamespace(ctx)
	r.handleAccountUser(ctx)
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
	r.POST("/v1/namespaces/:nid/users", func(c *gin.Context) {
		id := c.PostForm("id")
		nid := c.Param("nid")
		password := c.PostForm("password")
		userName := c.PostForm("username")
		resp, err := r.accountUserCli.CreateUser(ctx, &aupb.CreateUserRequest{
			Namespace: path.Join("namespaces", nid),
			Id:        id,
			User: &aupb.User{
				Username: userName,
				Password: password,
			},
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

func Resolve(p string) string {
	r, err := filepath.Abs(p)
	if err != nil {
		return p
	}
	return r
}

// func (r *Router) handleHome(ctx context.Context) {
// 	r.GET("/v1/namespaces/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		resp, err := r.doorNamespaceCli.GetNamespace(ctx, &npb.GetNamespaceRequest{
// 			Name: path.Join("namespace", id),
// 		})
// 		if err != nil {
// 			c.JSON(301, err.Error())
// 		}
// 		c.JSON(200, resp)
// 	})
// }

// func (r *Router) handleAccountNamespace(ctx context.Context) {
// 	r.GET("/v1/namespaces/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		resp, err := r.userUserCli.GetUser(ctx, &upb.GetUserRequest{
// 			Name: path.Join("users", id),
// 		})
// 		if err != nil {
// 			c.JSON(301, err.Error())
// 		}
// 		c.JSON(200, resp)
// 	})
// 	r.POST("/v1/namespaces", func(c *gin.Context) {
// 		id := c.PostForm("id")
// 		password := c.PostForm("password")
// 		resp, err := r.userUserCli.CreateUser(ctx, &upb.CreateUserRequest{
// 			Id: id,
// 			User: &upb.User{
// 				Password: password,
// 			},
// 		})
// 		if err != nil {
// 			c.JSON(301, err.Error())
// 		}
// 		c.JSON(200, resp)
// 	})
// 	r.DELETE("/v1/users/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		resp, err := r.userUserCli.DeleteUser(ctx, &upb.DeleteUserRequest{
// 			Name: path.Join("users", id),
// 		})
// 		if err != nil {
// 			c.JSON(301, err.Error())
// 		}
// 		c.JSON(200, resp)
// 	})
// }
