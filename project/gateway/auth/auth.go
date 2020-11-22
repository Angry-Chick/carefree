package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/carefree/project/common/jwt"
	"github.com/gin-gonic/gin"

	tpb "github.com/carefree/api/project/type/accesstoken"
)

type key int

const (
	keyBody key = iota
)

func JwtAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/carefree.project.account.v1.Account/BasicAuth" {
			c.Next()
			return
		}
		if c.Request.URL.Path == "/carefree.project.portal.v1.PortalService/SignUp" {
			c.Next()
			return
		}
		tkStr := c.GetHeader("Authorization")
		if tkStr == "" {
			c.JSON(http.StatusUnauthorized, "missing token")
			c.Abort()
			return
		}
		token := &tpb.Token{}
		if err := json.Unmarshal([]byte(tkStr), token); err != nil {
			c.JSON(http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}
		now := time.Now().UnixNano()
		if now > token.GetExpiry() {
			c.JSON(http.StatusUnauthorized, "token has expired")
			c.Abort()
			return
		}
		tk, err := jwt.New().ParseToken(token.Opaque)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "token is invalid")
			c.Abort()
			return
		}
		if c.Request.URL.Path == "/getUser" {
			c.JSON(http.StatusOK, tk.User)
			return
		}
		ctx := context.WithValue(c, keyBody, tk.User)
		c.Request.WithContext(ctx)
		c.Next()
	}
}
