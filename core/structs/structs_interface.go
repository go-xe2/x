package structs

import "reflect"

type ISliceField interface {
	ItemType() reflect.Type
}

// 指针引用类型字段
type IPtrField interface {
	// 指针指向类型名称
	OriginType() reflect.Type
}

type IMapField interface {
	KeyType() reflect.Type
	ValueType() reflect.Type
}

type IStructField interface {
	StructType() reflect.Type
}

type IField interface {
	Name() string
	Value() reflect.Value
	Tag() *FieldTag
	Type() reflect.Type
	Offset() uintptr
	Index() []int
	Get(inst ...interface{}) interface{}
	GetV(instValue ...reflect.Value) interface{}
	Set(val interface{}, inst ...interface{})
	SetV(val interface{}, instValue ...reflect.Value)
	Init(inst interface{}) IField
	Self() interface{}
}
