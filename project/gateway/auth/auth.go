package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/carefree/project/common/jwt"
	"github.com/gin-gonic/gin"

	tpb "github.com/carefree/api/project/type/accesstoken"
)

func JwtAuthentication(c *gin.Context) {
	whitelist := []string{
		"/carefree.project.account.v1.Account/BasicAuth",
		"/carefree.project.portal.v1.PortalService/SignUp",
	}
	if isWhitelist(c.Request.URL.Path, whitelist) {
		c.Next()
		return
	}
	tkStr := c.GetHeader("Authorization")
	if tkStr == "" {
		c.String(http.StatusUnauthorized, "%s", "missing token")
		c.Abort()
		return
	}
	token := &tpb.Token{}
	if err := json.Unmarshal([]byte(tkStr), token); err != nil {
		c.String(http.StatusUnauthorized, "%s", "invalid token")
		c.Abort()
		return
	}
	now := time.Now().Unix()
	if now > token.GetExpiry() {
		c.String(http.StatusUnauthorized, "%s", "token has expired")
		c.Abort()
		return
	}
	tk, err := jwt.New().ParseToken(token.Opaque)
	if err != nil {
		c.String(http.StatusUnauthorized, "%s", "token is invalid")
		c.Abort()
		return
	}
	if c.Request.URL.Path == "/getUserFromToken" {
		c.String(http.StatusOK, "%s", tk.User)
	}
	c.Set("user", tk.User)
	c.Next()
}

func isWhitelist(path string, whitelist []string) bool {
	for _, v := range whitelist {
		if path == v {
			return true
		}
	}
	return false
}
