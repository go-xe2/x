package structs

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

// struct字段值设置函数
type FieldSetFunc func(inst interface{}, field IField, value interface{}) error
type FieldSetFuncV func(instV reflect.Value, field IField, value interface{}) error
type FieldGetFunc func(inst interface{}, field IField) (interface{}, error)
type FieldGetFuncV func(instV reflect.Value, field IField) (interface{}, error)

func GetStructInstValue(inst interface{}) (reflect.Value, error) {
	dataType := reflect.TypeOf(inst)
	dataValue := reflect.ValueOf(inst)
	if dataType.Kind() == reflect.Ptr {
		if dataValue.IsNil() {
			return dataValue, errors.New("转入的指针为nil")
		}
		for ; dataType.Kind() == reflect.Ptr; dataType = dataType.Elem() {
			dataValue = dataValue.Elem()
		}
	} else {
		return dataValue, errors.New("转入的参数不是结构体指针类型")
	}
	if dataValue.Kind() != reflect.Struct {
		return dataValue, errors.New("转入的不是struct的指针")
	}
	return dataValue, nil
}

/*======================================================================================*/
/* 以下为反射设置struct字段值函数 															*/
/*=====================================================================================**/

func FieldSetStringFunc(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetStringFuncV(instValue, field, value)
}

func FieldSetStringFuncV(instV reflect.Value, field IField, value interface{}) error {
	s := anyToString(value, "")
	*((*string)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = s
	return nil
}

func FieldSetIntFunc(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetIntFuncV(instValue, field, value)
}

func FieldSetIntFuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToInt(value, 0)
	*((*int)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetInt8Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetInt8FuncV(instValue, field, value)
}

func FieldSetInt8FuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToInt8(value, 0)
	*((*int8)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetInt16Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetInt16FuncV(instValue, field, value)
}

func FieldSetInt16FuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToInt16(value, 0)
	*((*int16)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetInt32Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetInt32FuncV(instValue, field, value)
}

func FieldSetInt32FuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToInt32(value, 0)
	*((*int32)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetInt64Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetInt64FuncV(instValue, field, value)
}

func FieldSetInt64FuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToInt64(value, 0)
	*((*int64)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetUintFunc(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetUintFuncV(instValue, field, value)
}

func FieldSetUintFuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToUint(value, 0)
	*((*uint)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetUint8Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetUint8FuncV(instValue, field, value)
}

func FieldSetUint8FuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToUint8(value, 0)
	*((*uint8)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetUint16Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetUint16FuncV(instValue, field, value)
}

func FieldSetUint16FuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToUint16(value, 0)
	*((*uint16)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetUint32Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetUint32FuncV(instValue, field, value)
}

func FieldSetUint32FuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToUint32(value, 0)
	*((*uint32)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetUint64Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetUint64FuncV(instValue, field, value)
}

func FieldSetUint64FuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToUint64(value, 0)
	*((*uint64)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetFloat32Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetFloat32FuncV(instValue, field, value)
}

func FieldSetFloat32FuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToFloat(value, 0)
	*((*float32)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetFloat64Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetFloat64FuncV(instValue, field, value)
}

func FieldSetFloat64FuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToFloat64(value, 0)
	*((*float64)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetBoolFunc(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetBoolFuncV(instValue, field, value)
}

func FieldSetBoolFuncV(instV reflect.Value, field IField, value interface{}) error {
	v := anyToBool(value)
	*((*bool)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	return nil
}

func FieldSetComplex64Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetComplex64FuncV(instValue, field, value)
}

func FieldSetComplex64FuncV(instV reflect.Value, field IField, value interface{}) error {
	if v, ok := value.(complex64); ok {
		*((*complex64)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	} else {
		return errors.New("字段[" + field.Name() + "]只接受complex64类型值")
	}
	return nil
}

func FieldSetComplex128Func(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetComplex128FuncV(instValue, field, value)
}

func FieldSetComplex128FuncV(instV reflect.Value, field IField, value interface{}) error {
	if v, ok := value.(complex128); ok {
		*((*complex128)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	} else {
		return errors.New("字段[" + field.Name() + "]只接受complex128类型值")
	}
	return nil
}

func FieldSetUIntPtrFunc(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetUIntPtrFuncV(instValue, field, value)
}

func FieldSetUIntPtrFuncV(instV reflect.Value, field IField, value interface{}) error {
	if v, ok := value.(uintptr); ok {
		*((*uintptr)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = v
	} else {
		return errors.New("字段[" + field.Name() + "]只接受uintptr类型值")
	}
	return nil
}

func FieldSetPtrFunc(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetPtrFuncV(instValue, field, value)
}

func FieldSetPtrFuncV(instV reflect.Value, field IField, value interface{}) error {
	*((*uintptr)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = reflect.ValueOf(value).Pointer()
	return nil
}

func FieldSetSliceFunc(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetSliceFuncV(instValue, field, value)
}

func FieldSetSliceFuncV(instV reflect.Value, field IField, value interface{}) error {
	sliceMember, ok := field.(ISliceField)
	if !ok {
		return errors.New(fmt.Sprintf("字段[%s]不是slice类型", field.Name()))
	}
	fieldElemType := sliceMember.ItemType().Kind()
	valueKind := reflect.ValueOf(value)
	if valueKind.Kind() != reflect.Slice && valueKind.Kind() != reflect.Array {
		return errors.New(fmt.Sprintf("字段[%s]只接受slice类型", field.Name()))
	}
	valueElemType := valueKind.Type().Elem().Kind()
	if fieldElemType != valueElemType {
		return errors.New(fmt.Sprintf("字段[%s]只接受[]%v类型", field.Name(), fieldElemType))
	}
	instV.FieldByIndex(field.Index()).Set(reflect.ValueOf(value))
	return nil
}

func FieldSetMapFunc(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetMapFuncV(instValue, field, value)
}

func FieldSetMapFuncV(instV reflect.Value, field IField, value interface{}) error {
	mapMember, ok := field.(IMapField)
	if !ok {
		return errors.New(fmt.Sprintf("字段[%s]不是map类型", field.Name()))
	}
	valueType := reflect.TypeOf(value)
	if valueType.Kind() != reflect.Map {
		return errors.New(fmt.Sprintf("字段[%s]只接受map类型", field.Name()))
	}
	if valueType.Key().Kind() != mapMember.KeyType().Kind() || valueType.Elem().Kind() != mapMember.ValueType().Kind() {
		return errors.New(fmt.Sprintf("字段[%s]只接受map<%s,%s>类型", field.Name(), mapMember.KeyType().Name(), mapMember.ValueType().Name()))
	}
	*((*uintptr)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = reflect.ValueOf(value).Pointer()
	return nil
}

func FieldSetChanFunc(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetChanFuncV(instValue, field, value)
}

func FieldSetChanFuncV(instV reflect.Value, field IField, value interface{}) error {
	if reflect.TypeOf(value).Kind() != reflect.Chan {
		return errors.New(fmt.Sprintf("字段[%s]只接受chan类型", field.Name()))
	}
	*((*uintptr)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset()))) = reflect.ValueOf(value).Pointer()
	return nil
}

func FieldSetStructFunc(inst interface{}, field IField, value interface{}) error {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return err
	}
	return FieldSetStructFuncV(instValue, field, value)
}

func FieldSetStructFuncV(instV reflect.Value, field IField, value interface{}) error {
	if reflect.ValueOf(value).Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("字段[%s]只接受struct类型值", field.Name()))
	}
	instV.FieldByIndex(field.Index()).Set(reflect.ValueOf(value))
	return nil
}

/*======================================================================================*/
/* 以下为反射获取struct字段值函数 															*/
/*=====================================================================================**/

func FieldGetStringFunc(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetStringFuncV(instValue, field)
}

func FieldGetStringFuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetStringFuncVT(instV, field), nil
}

func FieldGetStringFuncVT(instV reflect.Value, field IField) string {
	return *((*string)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetIntFunc(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetIntFuncV(instValue, field)
}

func FieldGetIntFuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetIntFuncVT(instV, field), nil
}

func FieldGetIntFuncVT(instV reflect.Value, field IField) int {
	return *((*int)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetInt8Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetInt8FuncV(instValue, field)
}

func FieldGetInt8FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return MemberGetInt8FuncVT(instV, field), nil
}

func MemberGetInt8FuncVT(instV reflect.Value, field IField) int8 {
	return *((*int8)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetInt16Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetInt16FuncV(instValue, field)
}

func FieldGetInt16FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetInt16FuncVT(instV, field), nil
}

func FieldGetInt16FuncVT(instV reflect.Value, field IField) int16 {
	return *((*int16)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetInt32Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetInt32FuncV(instValue, field)
}

func FieldGetInt32FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetInt32FuncVT(instV, field), nil
}

func FieldGetInt32FuncVT(instV reflect.Value, field IField) int32 {
	return *((*int32)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetInt64Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetInt64FuncV(instValue, field)
}

func FieldGetInt64FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetInt64FuncVT(instV, field), nil
}

func FieldGetInt64FuncVT(instV reflect.Value, field IField) int64 {
	return *((*int64)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetUintFunc(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetUintFuncV(instValue, field)
}

func FieldGetUintFuncV(instV reflect.Value, field IField) (interface{}, error) {
	return MemberGetUintFuncVT(instV, field), nil
}

func MemberGetUintFuncVT(instV reflect.Value, field IField) uint {
	return *((*uint)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetUint8Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetUint8FuncV(instValue, field)
}

func FieldGetUint8FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetUint8FuncVT(instV, field), nil
}

func FieldGetUint8FuncVT(instV reflect.Value, field IField) uint8 {
	return *((*uint8)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetUint16Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetUint16FuncV(instValue, field)
}

func FieldGetUint16FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetUint16FuncVT(instV, field), nil
}

func FieldGetUint16FuncVT(instV reflect.Value, field IField) uint16 {
	return *((*uint16)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetUint32Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetUint32FuncV(instValue, field)
}

func FieldGetUint32FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetUint32FuncVT(instV, field), nil
}

func FieldGetUint32FuncVT(instV reflect.Value, field IField) uint32 {
	return *((*uint32)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetUint64Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetUint64FuncV(instValue, field)
}

func FieldGetUint64FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetUint64FuncVT(instV, field), nil
}

func FieldGetUint64FuncVT(instV reflect.Value, field IField) uint64 {
	return *((*uint64)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetBoolFunc(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetBoolFuncV(instValue, field)
}

func FieldGetBoolFuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetBoolFuncVT(instV, field), nil
}

func FieldGetBoolFuncVT(instV reflect.Value, field IField) bool {
	return *((*bool)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetFloat32Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetFloat32FuncV(instValue, field)
}

func FieldGetFloat32FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetFloat32FuncVT(instV, field), nil
}

func FieldGetFloat32FuncVT(instV reflect.Value, field IField) float32 {
	return *((*float32)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetFloat64Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetFloat64FuncV(instValue, field)
}

func FieldGetFloat64FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetFloat64FuncVT(instV, field), nil
}

func FieldGetFloat64FuncVT(instV reflect.Value, field IField) float64 {
	return *((*float64)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetComplex64Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetComplex64FuncV(instValue, field)
}

func FieldGetComplex64FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetComplex64FuncVT(instV, field), nil
}

func FieldGetComplex64FuncVT(instV reflect.Value, field IField) complex64 {
	return *((*complex64)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetComplex128Func(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetComplex128FuncV(instValue, field)
}

func FieldGetComplex128FuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetComplex128FuncVT(instV, field), nil
}

func FieldGetComplex128FuncVT(instV reflect.Value, field IField) complex128 {
	return *((*complex128)(unsafe.Pointer(instV.UnsafeAddr() + field.Offset())))
}

func FieldGetChanFunc(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetChanFuncV(instValue, field)
}

func FieldGetChanFuncV(instV reflect.Value, field IField) (interface{}, error) {
	return FieldGetInterfaceFuncV(instV, field)
}

func FieldGetSliceFunc(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetSliceFuncV(instValue, field)
}

func FieldGetSliceFuncV(instV reflect.Value, field IField) (interface{}, error) {
	fdV := reflect.NewAt(field.Type(), unsafe.Pointer(instV.UnsafeAddr()+field.Offset())).Elem()
	return fdV.Interface(), nil
}

func FieldGetMapFunc(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetMapFuncV(instValue, field)
}

func FieldGetMapFuncV(instV reflect.Value, field IField) (interface{}, error) {
	fdV := reflect.NewAt(field.Type(), unsafe.Pointer(instV.UnsafeAddr()+field.Offset())).Elem()
	return fdV.Interface(), nil
}

func FieldGetInterfaceFunc(inst interface{}, field IField) (interface{}, error) {
	instValue, err := GetStructInstValue(inst)
	if err != nil {
		return nil, err
	}
	return FieldGetInterfaceFuncV(instValue, field)
}

func FieldGetInterfaceFuncV(instV reflect.Value, field IField) (interface{}, error) {
	fdV := reflect.NewAt(field.Type(), unsafe.Pointer(instV.UnsafeAddr()+field.Offset())).Elem()
	return fdV.Interface(), nil
}
