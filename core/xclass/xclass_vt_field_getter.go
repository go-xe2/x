package xclass

import "reflect"

type classFieldGetter func(inst reflect.Value, field *classVTField) (v interface{}, err error)

var classFieldGetters = map[reflect.Kind]classFieldGetter{
	reflect.Interface:  FieldGetInterfaceFunc,
	reflect.String:     FieldGetStringFunc,
	reflect.Bool:       FieldGetBoolFunc,
	reflect.Float32:    FieldGetFloat32Func,
	reflect.Float64:    FieldGetFloat64Func,
	reflect.Int:        FieldGetIntFunc,
	reflect.Int8:       FieldGetInt8Func,
	reflect.Int16:      FieldGetInt16Func,
	reflect.Int32:      FieldGetInt32Func,
	reflect.Int64:      FieldGetInt64Func,
	reflect.Uint:       FieldGetUintFunc,
	reflect.Uint8:      FieldGetUint8Func,
	reflect.Uint16:     FieldGetUint16Func,
	reflect.Uint32:     FieldGetUint32Func,
	reflect.Uint64:     FieldGetUint64Func,
	reflect.Complex64:  FieldGetComplex64Func,
	reflect.Complex128: FieldGetComplex128Func,
	reflect.Ptr:        FieldGetInterfaceFunc,
	reflect.Uintptr:    FieldGetInterfaceFunc,
	reflect.Slice:      FieldGetInterfaceFunc,
	reflect.Array:      FieldGetInterfaceFunc,
	reflect.Map:        FieldGetInterfaceFunc,
	reflect.Chan:       FieldGetInterfaceFunc,
	reflect.Struct:     FieldGetInterfaceFunc,
}
