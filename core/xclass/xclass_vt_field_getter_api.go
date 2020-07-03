package xclass

import (
	"reflect"
	"unsafe"
)

func FieldGetStringFunc(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*string)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetIntFunc(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*int)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetInt8Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*int8)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetInt16Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*int16)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetInt32Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*int32)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetInt64Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*int64)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetUintFunc(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*uint)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetUint8Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*uint8)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetUint16Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*uint16)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetUint32Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*uint32)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetUint64Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*uint64)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetBoolFunc(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*bool)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetFloat32Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*float32)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetFloat64Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*float64)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetComplex64Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*complex64)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetComplex128Func(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	return *((*complex128)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))), nil
}

func FieldGetInterfaceFunc(inst reflect.Value, field *classVTField) (v interface{}, err error) {
	fv := reflect.NewAt(field.fieldType, unsafe.Pointer(inst.Elem().UnsafeAddr()+field.offset)).Elem()
	return fv.Interface(), nil
}
