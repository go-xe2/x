package t

import (
	"reflect"
)

var _ T = &Type{}

func ifNilDef(v interface{}, def interface{}) interface{} {
	if v == nil {
		return def
	}
	return v
}

func (t Type) Any(def ...interface{}) interface{} {
	return ifNilDef(t.val, def)
}

func (t Type) String(def ...string) string {
	return String(t.val, def...)
}

func (t Type) Float64(def ...float64) float64 {
	return Float64(t.val, def...)
}
func (t Type) Float32(def ...float32) float32 {
	return Float32(t.val, def...)
}

func (t Type) Int64(def ...int64) int64 {
	return Int64(t.val, def...)
}
func (t Type) Int(def ...int) int {
	return Int(t.val, def...)
}
func (t Type) Int32(def ...int32) int32 {
	return Int32(t.val, def...)
}
func (t Type) Int16(def ...int16) int16 {
	return Int16(t.val, def...)
}
func (t Type) Int8(def ...int8) int8 {
	return Int8(t.val, def...)
}

func (t Type) Uint64(def ...uint64) uint64 {
	return Uint64(t.val, def...)
}
func (t Type) Uint(def ...uint) uint {
	return Uint(t.val, def...)
}
func (t Type) Uint32(def ...uint32) uint32 {
	return Uint32(t.val, def...)
}
func (t Type) Uint16(def ...uint16) uint16 {
	return Uint16(t.val, def...)
}
func (t Type) Uint8(def ...uint8) uint8 {
	return Uint8(t.val, def...)
}

func (t Type) Bool(def ...bool) bool {
	return Bool(t.val, def...)
}

// 将数据转转换成其他类型
func (t Type) ToType(typ reflect.Type) interface{} {
	return valToType(t, typ)
}

func any2anyPointer(v interface{}) interface{} {
	pv := reflect.New(reflect.TypeOf(v))
	pv.Elem().Set(reflect.ValueOf(v))
	return pv.Interface()
}

func valToType(val T, toType reflect.Type) (valValue interface{}) {
	switch toType.Kind() {
	case reflect.String:
		valValue = val.String()
		break
	case reflect.Bool:
		valValue = val.Bool()
		break
	case reflect.Float32:
		valValue = val.Float32()
		break
	case reflect.Float64:
		valValue = val.Float64()
		break
	case reflect.Int:
		valValue = val.Int()
		break
	case reflect.Int8:
		valValue = val.Int8()
		break
	case reflect.Int16:
		valValue = val.Int16()
		break
	case reflect.Int32:
		valValue = val.Int32()
		break
	case reflect.Int64:
		valValue = val.Int64()
		break
	case reflect.Uint:
		valValue = val.Uint()
		break
	case reflect.Uint8:
		valValue = val.Uint8()
		break
	case reflect.Uint16:
		valValue = val.Uint16()
		break
	case reflect.Uint32:
		valValue = val.Uint32()
		break
	case reflect.Uint64:
		valValue = val.Uint64()
		break
	case reflect.Complex64:
		if c1, ok := val.Any().(complex64); ok {
			valValue = c1
		} else {
			valValue = complex64(0)
		}
		break
	case reflect.Complex128:
		if c1, ok := val.Any().(complex128); ok {
			valValue = c1
		} else {
			valValue = complex128(0)
		}
		break
	case reflect.Interface:
		valValue = val.Any()
		break
	case reflect.Slice:
		valValue = val.ToSliceType(toType)
		break
	case reflect.Map:
		valValue = val.ToMapType(toType)
		break
	case reflect.Struct:
		valValue = val.ToStructType(toType)
		break
	case reflect.Ptr:
		ptrElemType := toType.Elem()
		for ptrElemType.Kind() == reflect.Ptr {
			ptrElemType = ptrElemType.Elem()
		}
		if ptrElemType.Kind() == reflect.Struct {
			valValue = val.ToStructType(toType)
		} else {
			valValue = valToType(val, ptrElemType)
		}
		break
	default:
		return val.Any()
	}
	return
}
