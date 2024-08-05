package core

import (
	"fmt"
	"gin_gorm/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const filePath = "settings.yaml"

func InitConf() {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		fmt.Printf("读取配置信息失败, err:%v\n", err)
		return
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("viper反序列化失败, err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变...")
		if err := viper.Unmarshal(global.Config); err != nil {
			fmt.Printf("viper反序列化失败, err:%v\n", err)
		}
	})
	return
}
