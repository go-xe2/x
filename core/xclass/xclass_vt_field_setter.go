package xclass

import "reflect"

type classFieldSetter func(inst reflect.Value, field *classVTField, value interface{}) error

// class字段数据类型设置方法集合
var classFieldSetters = map[reflect.Kind]classFieldSetter{
	reflect.Interface:  FieldSetAnyFunc,
	reflect.String:     FieldSetStringFunc,
	reflect.Bool:       FieldSetBoolFunc,
	reflect.Float32:    FieldSetFloat32Func,
	reflect.Float64:    FieldSetFloat64Func,
	reflect.Int:        FieldSetIntFunc,
	reflect.Int8:       FieldSetInt8Func,
	reflect.Int16:      FieldSetInt16Func,
	reflect.Int32:      FieldSetInt32Func,
	reflect.Int64:      FieldSetInt64Func,
	reflect.Uint:       FieldSetUintFunc,
	reflect.Uint8:      FieldSetUint8Func,
	reflect.Uint16:     FieldSetUint16Func,
	reflect.Uint32:     FieldSetUint32Func,
	reflect.Uint64:     FieldSetUint64Func,
	reflect.Complex64:  FieldSetComplex64Func,
	reflect.Complex128: FieldSetComplex128Func,
	reflect.Uintptr:    FieldSetUintPtrFunc,
	reflect.Ptr:        FieldSetPtrFunc,
	reflect.Slice:      FieldSetSliceFunc,
	reflect.Array:      FieldSetSliceFunc,
	reflect.Map:        FieldSetMapFunc,
	reflect.Struct:     FieldSetStructFunc,
	reflect.Chan:       FieldSetChanFunc,
}
