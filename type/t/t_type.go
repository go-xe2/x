package t

import "reflect"

type T interface {
	Any(def ...interface{}) interface{}
	String(def ...string) string

	Float32(def ...float32) float32
	Float64(def ...float64) float64

	Int(def ...int) int
	Int8(def ...int8) int8
	Int16(def ...int16) int16
	Int32(def ...int32) int32
	Int64(def ...int64) int64

	Uint(def ...uint) uint
	Uint8(def ...uint8) uint8
	Uint16(def ...uint16) uint16
	Uint32(def ...uint32) uint32
	Uint64(def ...uint64) uint64

	Bool(def ...bool) bool

	Slice() TSlice
	SliceAny() []interface{}
	SliceMapStrAny() []map[string]interface{}
	SliceHash() []THash
	SliceString() []string
	SliceInt64() []int64
	ToSliceType(sliceType reflect.Type) interface{}

	Map() TMap
	MapAny() TMapAny
	MapString() TMapString
	MapInt64() TMapInt64
	MapStrAny() map[string]interface{}
	Hash() THash
	ToMapType(mapType reflect.Type) interface{}

	ToStructType(structType reflect.Type) interface{}

	ToType(typ reflect.Type) interface{}
}

type TMap map[T]T
type TMapAny map[interface{}]T
type TMapString map[string]T
type TMapInt64 map[int64]T
type TMapInt32 map[int32]T
type TMapInt map[int64]T
type THash map[string]interface{}

type TSlice []T
type TSliceAny []interface{}
type TSliceString []string
type TSliceInt64 []int64

type Type struct {
	val interface{}
}

var AnyType = reflect.TypeOf((interface{})(nil))
var StringType = reflect.TypeOf(string(""))
var BoolType = reflect.TypeOf(Bool(false))
var IntType = reflect.TypeOf(int(0))
var Int8Type = reflect.TypeOf(int8(0))
var Int16Type = reflect.TypeOf(int16(0))
var Int32Type = reflect.TypeOf(int32(0))
var Int64Type = reflect.TypeOf(int64(0))
var UintType = reflect.TypeOf(uint(0))
var Uint8Type = reflect.TypeOf(uint8(0))
var Uint16Type = reflect.TypeOf(uint16(0))
var Uint32Type = reflect.TypeOf(uint32(0))
var Uint64Type = reflect.TypeOf(uint64(0))
var Float32Type = reflect.TypeOf(Float32(0))
var Float64Type = reflect.ValueOf(Float64(0))
var Complex64Type = reflect.TypeOf(complex64(0))
var Complex128Type = reflect.TypeOf(complex128(0))

// 基础slice数据类型
var SliceAnyType = reflect.TypeOf([]interface{}{})
var SliceStringType = reflect.TypeOf([]string{})
var SliceBoolType = reflect.TypeOf([]bool{})
var SliceIntType = reflect.TypeOf([]int{})
var SliceInt8Type = reflect.TypeOf([]int8{})
var SliceInt16Type = reflect.TypeOf([]int16{})
var SliceInt32Type = reflect.TypeOf([]int32{})
var SliceInt64Type = reflect.TypeOf([]int64{})
var SliceUintType = reflect.TypeOf([]uint{})
var SliceUint8Type = reflect.TypeOf([]uint8{})
var SliceUint16Type = reflect.TypeOf([]uint16{})
var SliceUint32Type = reflect.TypeOf([]uint32{})
var SliceUint64Type = reflect.TypeOf([]uint64{})
var SliceFloat32Type = reflect.TypeOf([]float32{})
var SliceFloat64Type = reflect.TypeOf([]float64{})
var SliceComplex64Type = reflect.TypeOf([]complex64{})
var SliceComplex128Type = reflect.TypeOf([]complex128{})
