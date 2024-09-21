package flags

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"template/global"
)

func DB(c *cli.Context) (err error) {
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate()
	if err != nil {
		zap.L().Error("[ error ] 生成数据库表结构失败")
		return nil
	}
	zap.L().Info("[ success ] 生成数据库表结构成功！")
	return nil
}
