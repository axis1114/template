package main

import (
	"fmt"
	"gin_gorm/core"
	"gin_gorm/flags"
	"gin_gorm/global"
	"gin_gorm/routers"
	"gin_gorm/utils/validator"
)

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
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	router := routers.InitRouter()
	err = router.Run(fmt.Sprintf(":%d", global.Config.System.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
