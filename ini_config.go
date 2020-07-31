package iniconfig

import (
	"errors"
	"fmt"
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
	//获取用户传进来的类型信息,如果result传进来的不是指针类型,则报错
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("please pass address")
		return
	}

	//获取用户传进来的类型信息,如果result不是结构体类型,则报错
	typeInfo2 := typeInfo.Elem()
	if typeInfo2.Kind() != reflect.Struct {
		err = errors.New("please pass struct")
		return err
	}
	// 定义变量,存储section的名字
	var lastSectionFieldName string
	for index, value := range lineArr {
		line := strings.TrimSpace(value)
		//如果长度为0,则认为是空行, continue
		if len(line) == 0 {
			continue
		}

		//如果首字母是 ';'或者'#',则 continue
		if line[0] == ';' || line[0] == '#' {
			continue
		}

		//首字母是'[',并且长度小于等于2,则认为非法
		if line[0] == '[' && len(line) <= 2 {
			err = fmt.Errorf("syntax error,invalid section: %s,lineNo: %d", line, index+1)
			return
		}
		//首字母是 [ 并且尾字母不是 ],则认为非法
		if line[0] == '[' && line[len(line)-1] != ']' {
			err = fmt.Errorf("syntax error,invalid section: %s,lineNo: %d", line, index+1)
			return
		}

		//判断section的是否为空,为空则认为非法
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error,invalid section: %s,lineNo: %d", line, index+1)
			return
		}

		for i := 0; i < typeInfo2.NumField(); i++ {
			field := typeInfo2.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == sectionName {
				lastSectionFieldName = field.Name
				fmt.Printf("field name: %s\n", lastSectionFieldName)
				break
			}
		}

	}
	return
}
