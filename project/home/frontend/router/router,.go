package router

import (
	"context"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	dpb "github.com/carefree/api/project/home/v1/door"'
	npb "github.com/carefree/api/project/home/v1/namespace"
	upb "github.com/carefree/api/project/user/v1"
)

type serviceName string

const (
	DoorService serviceName = "door"
	UserService serviceName = "user"
)

type ServiceConn map[serviceName]*grpc.ClientConn

var DefaultServiceConn = make(ServiceConn)

func (s ServiceConn) RegisterService(sn serviceName, conn *grpc.ClientConn) { s[sn] = conn }

type Router struct {
	*gin.Engine
	doorNamespaceCli npb.NamespaceServiceClient
	doorDoorsCli     dpb.DoorServiceClient
	userUserCli      upb.UserServiceClient
}

func New(sc ServiceConn) *Router {
	r := Router{
		Engine:           gin.Default(),
		doorNamespaceCli: npb.NewNamespaceServiceClient(sc[DoorService]),
		doorDoorsCli:     dpb.NewDoorServiceClient(sc[DoorService]),
		userUserCli:      upb.NewUserServiceClient(sc[UserService]),
	}
	return &r
}

var root = Resolve("../")

var buildPath = "carefree/project/door/frontend/build"

func (r *Router) RegisterHandle(ctx context.Context) {
	r.LoadHTMLFiles(path.Join(root, buildPath, "index.html"))
	r.Static("/static", path.Join(root, buildPath, "static"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Door"})
	})
	r.handleDoorNamespace(ctx)
	r.handleUserUser(ctx)
}

func Resolve(p string) string {
	r, err := filepath.Abs(p)
	if err != nil {
		return p
	}
	return r
}

func (r *Router) handleDoorNamespace(ctx context.Context) {
	r.GET("/v1/namespaces/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := r.doorNamespaceCli.GetNamespace(ctx, &npb.GetNamespaceRequest{
			Name: path.Join("namespace", id),
		})
		if err != nil {
			c.JSON(301, err.Error())
		}
		c.JSON(200, resp)
	})
}

func (r *Router) handleUserUser(ctx context.Context) {
	r.GET("/v1/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := r.userUserCli.GetUser(ctx, &upb.GetUserRequest{
			Name: path.Join("users", id),
		})
		if err != nil {
			c.JSON(301, err.Error())
		}
		c.JSON(200, resp)
	})
	r.POST("/v1/users", func(c *gin.Context) {
		id := c.PostForm("id")
		password := c.PostForm("password")
		resp, err := r.userUserCli.CreateUser(ctx, &upb.CreateUserRequest{
			Id: id,
			User: &upb.User{
				Password: password,
			},
		})
		if err != nil {
			c.JSON(301, err.Error())
		}
		c.JSON(200, resp)
	})
	r.DELETE("/v1/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		resp, err := r.userUserCli.DeleteUser(ctx, &upb.DeleteUserRequest{
			Name: path.Join("users", id),
		})
		if err != nil {
			c.JSON(301, err.Error())
		}
		c.JSON(200, resp)
	})
}
