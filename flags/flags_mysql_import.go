package flags

import (
	"gin_gorm/global"
	"os"
	"strings"
)

func MysqlImport(path string) {
	byteData, err := os.ReadFile(path)
	if err != nil {
		global.Log.Error("mysql import err:", err)
	}
	//分割数据 一定要按照\r\n分割
	sqlList := strings.Split(string(byteData), ";\r\n")
	for _, sql := range sqlList {
		//去除字符串开头和结尾的空白符
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		//执行sql语句
		err = global.DB.Exec(sql).Error
		if err != nil {
			global.Log.Error("mysql import err:", err)
			continue
		}
	}
}
