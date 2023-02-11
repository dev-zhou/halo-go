package _struct

import (
	"encoding/json"
	"reflect"
)

func New[S comparable](class S) *S {
	return &class
}

// StructCopy 结构体拷贝 参数均为 指针类型 没有考虑嵌套的问题
func StructCopy(origin, result interface{}) {
	originValue := reflect.ValueOf(origin)
	resultValue := reflect.ValueOf(result)
	if originValue.Kind() != reflect.Ptr || resultValue.Kind() != reflect.Ptr {
		return
	}
	if originValue.IsNil() || resultValue.IsNil() {
		return
	}
	originElem := originValue.Elem()
	resultElem := resultValue.Elem()
	for i := 0; i < resultElem.NumField(); i++ {
		resultField := resultElem.Type().Field(i)
		originFieldName, ok := originElem.Type().FieldByName(resultField.Name)
		// TODO 根据需要判断是否是空值从而是否 reset 值
		if ok && originFieldName.Type == resultField.Type {
			resultElem.Field(i).Set(originElem.FieldByName(resultField.Name))
		}
		elemType := resultElem.Field(i)
		if elemType.Kind() == reflect.Struct {
			childStructElem := resultElem.Field(i)
			for j := 0; j < childStructElem.NumField(); j++ {
				childStructField := childStructElem.Type().Field(j)
				originFieldName, ok := originElem.Type().FieldByName(childStructField.Name)
				if ok && originFieldName.Type == childStructField.Type {
					childStructElem.Field(j).Set(originElem.FieldByName(childStructField.Name))
				}
			}
		}
	}
}

// StructToMap 结构体转map
func StructToMap[T comparable](obj T) (map[string]interface{}, error) {
	// 结构体转json
	j, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	// json转map
	var m map[string]interface{}
	err1 := json.Unmarshal(j, &m)
	if err1 != nil {
		return nil, err1
	}
	return m, nil
}
