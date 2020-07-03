package xclass

import (
	"errors"
	"fmt"
	"github.com/go-xe2/x/type/t"
	"reflect"
	"unsafe"
)


func FieldSetAnyFunc(inst reflect.Value, field *classVTField, value interface{}) error  {
	*((*interface{})(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = value
	return nil
}

func FieldSetStringFunc(inst reflect.Value, field *classVTField, value interface{}) error {
	s := t.String(value, "")
	*((*string)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = s
	return nil
}

func FieldSetIntFunc(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Int(value, 0)
	*((*int)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetInt8Func(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Int8(value, 0)
	*((*int8)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetInt16Func(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Int16(value, 0)
	*((*int16)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetInt32Func(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Int32(value, 0)
	*((*int32)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetInt64Func(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Int64(value, 0)
	*((*int64)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetUintFunc(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Uint(value, 0)
	*((*uint)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetUint8Func(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Uint8(value, 0)
	*((*uint8)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetUint16Func(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Uint16(value, 0)
	*((*uint16)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetUint32Func(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Uint32(value, 0)
	*((*uint32)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetUint64Func(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Uint64(value, 0)
	*((*uint64)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetFloat32Func(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Float32(value, 0)
	*((*float32)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetFloat64Func(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Float64(value, 0)
	*((*float64)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetBoolFunc(inst reflect.Value, field *classVTField, value interface{}) error {
	v := t.Bool(value)
	*((*bool)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	return nil
}

func FieldSetComplex64Func(inst reflect.Value, field *classVTField, value interface{}) error {
	if v, ok := value.(complex64); ok {
		*((*complex64)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	} else {
		return errors.New("字段[" + field.name + "]只接受complex64类型值")
	}
	return nil
}

func FieldSetComplex128Func(inst reflect.Value, field *classVTField, value interface{}) error {
	if v, ok := value.(complex128); ok {
		*((*complex128)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	} else {
		return errors.New("字段[" + field.name + "]只接受complex128类型值")
	}
	return nil
}

func FieldSetUintPtrFunc(inst reflect.Value, field *classVTField, value interface{}) error {
	if v, ok := value.(uintptr); ok {
		*((*uintptr)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = v
	} else {
		return errors.New("字段[" + field.name + "]只接受uintptr类型值")
	}
	return nil
}

func FieldSetPtrFunc(inst reflect.Value, field *classVTField, value interface{}) ( err error) {
	v := t.New(value).ToType(field.fieldType)
	fd := reflect.NewAt(field.fieldType, unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset)).Elem()
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()
	fd.Set(reflect.ValueOf(v))
	return
}

func FieldSetSliceFunc(inst reflect.Value, field *classVTField, value interface{}) (err error) {
	v := t.New(value).ToSliceType(field.fieldType)
	fd := reflect.NewAt(field.fieldType, unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset)).Elem()
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()
	fd.Set(reflect.ValueOf(v))
	return
}

func FieldSetMapFunc(inst reflect.Value, field *classVTField, value interface{}) (err error) {
	v := t.New(value).ToMapType(field.fieldType)
	fd := reflect.NewAt(field.fieldType,unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset)).Elem()
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()
	fd.Set(reflect.ValueOf(v))
	return
}

func FieldSetChanFunc(inst reflect.Value, field *classVTField, value interface{}) error {
	if reflect.TypeOf(value).Kind() != reflect.Chan {
		return errors.New(fmt.Sprintf("字段[%s]只接受chan类型", field.name))
	}
	*((*uintptr)(unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset))) = reflect.ValueOf(value).Pointer()
	return nil
}

func FieldSetStructFunc(inst reflect.Value, field *classVTField, value interface{}) (err error) {
	v := t.New(value).ToStructType(field.fieldType)
	fd := reflect.NewAt(field.fieldType, unsafe.Pointer(inst.Elem().UnsafeAddr() + field.offset)).Elem()
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()
	fd.Set(reflect.ValueOf(v))
	return
}
