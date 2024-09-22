package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"template/global"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Fatal("未配置mysql，取消gorm连接")
	}

	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "dev" {
		// 开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 只打印错误的sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	//通过gorm.Open()函数连接到MySQL数据库，并设置日志记录器为上一步配置的mysqlLogger
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		lg.Fatal(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
	// 最多可容纳
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	// 连接最大复用时间，不能超过mysql的wait_timeout
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
