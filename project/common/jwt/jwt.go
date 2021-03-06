package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	identityTokenDefaultTTL = 7 * 24 * time.Hour
)

// JWT 结构包含了 sign jwt 需要的参数。
type JWT struct {
	SigningKey []byte
}

// New 初始化一个 JWT 对象
// TODO(ljy): 暂时先使用默认的 singing key，之后从环境变量中读取
func New() *JWT {
	return &JWT{SigningKey: []byte("carefree.com")}
}

// Claims 自定义有效载荷
type Claims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

// GenerateToken 生成一个 token
func (j *JWT) GenerateToken(claims *Claims) (string, error) {
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tk.SignedString(j.SigningKey)
}

// ParseToken 解析 Token
func (j *JWT) ParseToken(ts string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(ts, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	c, ok := t.Claims.(*Claims)
	if !(ok && t.Valid) {
		return nil, fmt.Errorf("无效的 token")
	}
	return c, nil
}

// IdentityTokenExpiry 默认使用 7 天的有效期.
func IdentityTokenExpiry() time.Time {
	return time.Now().Add(identityTokenDefaultTTL)
}
