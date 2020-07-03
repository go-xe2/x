package xclass

import "reflect"

type classVTMethod struct {
	// 方法在类中的序号
	index int
	// 方法名称
	name string
	// 方法输入参数类型列表
	paramTypes  []reflect.Type
	resultTypes []reflect.Type
}
