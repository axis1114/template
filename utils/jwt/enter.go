package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type PayLoad struct {
	Username string `json:"username"` // 用户名
	Role     int    `json:"role"`     // 权限
	UserID   uint   `json:"user_id"`  // 用户id
}

var MySecret []byte

type CustomClaims struct {
	PayLoad
	jwt.StandardClaims
}
