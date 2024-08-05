package jwt

import (
	"gin_gorm/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

func GetToken(payload PayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		payload,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), //过期时间
			Issuer:    global.Config.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}
