package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
	"template/core"
	_ "template/docs"
	"template/global"
	"template/utils"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	//设置gin模式
	gin.SetMode(global.Config.System.Env)
	router := gin.New()
	router.Use(core.GinLogger(), core.GinRecovery(true))
	router.Use(utils.Cors())
	//将指定目录下的文件提供给客户端
	//"uploads" 是URL路径前缀，http.Dir("uploads")是实际文件系统中存储文件的目录
	router.StaticFS("upload", http.Dir("upload"))
	//注册swagger路由
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//创建路由组
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置api
	routerGroupApp.CaptchaRouter()
	return router
}
