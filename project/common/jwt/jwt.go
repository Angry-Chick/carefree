package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	identityTokenDefaultTTL = 7 * 24 * time.Hour
)

// JWT 包含签名秘钥
type JWT struct {
	SigningKey []byte
}

// New 初始化一个 JWT 对象
func New() *JWT {
	return &JWT{SigningKey: []byte("carefree.com")}
}

// Claims 自定义有效载荷(这里采用自定义的Name和Email作为有效载荷的一部分)
type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
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

// IdentityTokenExpiry define a 7 day's expiryTime
func IdentityTokenExpiry() time.Time {
	return time.Now().Add(identityTokenDefaultTTL)
}
