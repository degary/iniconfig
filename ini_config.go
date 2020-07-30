package iniconfig

import (
	"errors"
	"reflect"
	"strings"
)

func Marshal(data interface{}) (result []byte, err error) {
	return
}

func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	_ = lineArr
	//var m map[string]string
	//获取用户传进来的类型信息,如果result传进来的不是指针,则报错
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass address")
		return
	}

	return
}
