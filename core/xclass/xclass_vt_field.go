package xclass

import "reflect"

type classVTField struct {
	// 字段所在序号
	index []int
	// 字段名称
	name string
	// 字段类型
	fieldType reflect.Type
	// 是否可以设置值
	canSet bool
	// 字段tag字符串
	fieldTag ClassTag
	rawValue interface{}
	// 字段所属的类
	owner *classVT
	// 字段偏移量
	offset uintptr
	// 字段读取器指针
	getter classFieldGetter
	// 字段设置器指针
	setter classFieldSetter
}
