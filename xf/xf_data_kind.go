package xf

import "reflect"

const (
	Bool       = reflect.Bool
	Int        = reflect.Int
	Int8       = reflect.Int8
	Int16      = reflect.Int16
	Int32      = reflect.Int32
	Int64      = reflect.Int64
	Uint       = reflect.Uint
	Uint8      = reflect.Uint8
	Uint16     = reflect.Uint16
	Uint32     = reflect.Uint32
	Uint64     = reflect.Uint64
	Float32    = reflect.Float32
	Float64    = reflect.Float64
	Complex64  = reflect.Complex64
	Complex128 = reflect.Complex128
	Array      = reflect.Array
	Func       = reflect.Func
	Any        = reflect.Interface
	Map        = reflect.Map
	Slice      = reflect.Slice
	String     = reflect.String
	Struct     = reflect.Struct
)

var (
	BoolType    = reflect.TypeOf(bool(false))
	IntType     = reflect.TypeOf(int(0))
	Int8Type    = reflect.TypeOf(int8(0))
	Int16Type   = reflect.TypeOf(int16(0))
	Int32Type   = reflect.TypeOf(int32(0))
	Int64Type   = reflect.TypeOf(int64(0))
	UintType    = reflect.TypeOf(uint(0))
	Uint8Type   = reflect.TypeOf(uint8(0))
	Uint16Type  = reflect.TypeOf(uint16(0))
	Uint32Type  = reflect.TypeOf(uint32(0))
	Uint64Type  = reflect.TypeOf(uint64(0))
	Float32Type = reflect.TypeOf(float32(0))
	Float64Type = reflect.TypeOf(float64(0))
	AnyType     = reflect.TypeOf([]interface{}{}).Elem()
)
