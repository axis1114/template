package main

import (
	"fmt"
	"go.uber.org/zap"
	"template/core"
	"template/flags"
	"template/global"
	"template/routers"
	"template/utils"
	"template/utils/validator"
)

// @title Template
// @version 1.0
// @contact.name Axios
// @contact.email 1790146932@qq.com
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	core.InitConf()
	global.Log = core.InitLog()
	global.DB = core.InitGorm()
	global.Redis = core.InitRedis()
	global.Es = core.InitEs()
	global.Etcd = core.InitEtcd()
	flags.Newflags()
	err := validator.InitTrans("en")
	if err != nil {
		global.Log.Error("fail to init trans", zap.Error(err))
		return
	}
	utils.PrintSystem()
	router := routers.InitRouter()
	err = router.Run(fmt.Sprintf(":%d", global.Config.System.Port))
	if err != nil {
		global.Log.Error("fail to start server", zap.Error(err))
		return
	}
}
