package main

import (
	"fmt"
	"gin_gorm/core"
	"gin_gorm/flags"
	"gin_gorm/global"
	"gin_gorm/routers"
	"gin_gorm/utils"
	"gin_gorm/utils/validator"
	"go.uber.org/zap"
)

// @title Template
// @version 1.0
// @contact.name Axios
// @contact.email 1790146932@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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
