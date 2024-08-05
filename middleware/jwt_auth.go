package middleware

import (
	"gin_gorm/models/res"
	ctype "gin_gorm/models/types"
	"gin_gorm/service/redis_ser"
	"gin_gorm/utils/jwt"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
