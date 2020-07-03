package xclass

import "reflect"

type classVT struct {
	// class类类型
	clsType reflect.Type
	// class类继类，支持继承多个父类
	extends []*classVTExtendField
	// 字段映射
	fieldMaps map[string]*classVTField
	// 类方法序号名称映射
	methodMaps map[string]*classVTMethod
	classTag   ClassTag
}
