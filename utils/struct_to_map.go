package utils

import (
	"fmt"
	"reflect"
)

// ToMap 结构体转为Map[string]interface{}
func ToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{}) // 创建一个空的 map，用于存储结果

	// 使用反射获取输入参数的值和类型
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // 如果是指针，则获取其指向的值
	}

	// 检查输入参数类型是否为结构体
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", in)
	}

	// 遍历结构体的字段
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		// 获取字段的标签值
		tagValue := field.Tag.Get(tagName)
		if tagValue != "" {
			// 将字段的标签值和值存入 map
			out[tagValue] = v.Field(i).Interface()
		}
	}

	return out, nil
}
