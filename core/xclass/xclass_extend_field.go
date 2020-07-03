package xclass

import "reflect"

type classExtendField struct {
	// 继承类字段信息
	field *classVTExtendField
	// 被继承类字段类型实例
	inst interface{}
	// 被继承类字段实例value
	value reflect.Value
}

func newClassExtendField(field *classVTExtendField, value reflect.Value, inst interface{}) *classExtendField {
	return &classExtendField{
		field: field,
		inst:  inst,
		value: value,
	}
}

func (cef *classExtendField) Field() *classVTExtendField {
	return cef.field
}

func (cef *classExtendField) Type() reflect.Type {
	return cef.field.extendType.clsType
}

func (cef *classExtendField) Offset() uintptr {
	return cef.field.offset
}

func (cef *classExtendField) Instance() interface{} {
	return cef.inst
}

func (cef *classExtendField) Value() reflect.Value {
	return cef.value
}
