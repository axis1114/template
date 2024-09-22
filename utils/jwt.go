package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"template/global"
	"time"
)

type PayLoad struct {
	Account string `json:"Account"` // 账号
	Role    int    `json:"role"`    // 权限
	UserID  uint   `json:"user_id"` // 用户id
}

var MySecret []byte

type CustomClaims struct {
	PayLoad
	jwt.StandardClaims
}

func GetToken(payload PayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		payload,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24 * time.Duration(global.Config.Jwt.Expires))), //过期时间
			Issuer:    global.Config.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error("parse token error:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
