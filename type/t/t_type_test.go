package t

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T)  {
	fmt.Println("anyType:", AnyType)
	fmt.Println("stringType:", StringType)
	fmt.Println("boolType:", BoolType)
	fmt.Println("intType:", IntType)
	fmt.Println("int8Type:", Int8Type)
	fmt.Println("int16Type:", Int16Type)
	fmt.Println("int32Type:", Int32Type)
	fmt.Println("int64Type:", Int64Type)
	fmt.Println("uintType:", UintType)
	fmt.Println("uint8Type:", Uint8Type)
	fmt.Println("uint16Type:", Uint16Type)
	fmt.Println("uint32Type:", Uint32Type)
	fmt.Println("uint64Type:", Uint64Type)
	fmt.Println("float32Type:", Float32Type)
	fmt.Println("float64Type:", Float64Type)
	fmt.Println("complex64Type:", Complex64Type)
	fmt.Println("complex128Type:", Complex128Type)

	fmt.Println("sliceAnyType:", SliceAnyType)
	fmt.Println("sliceStringType:", SliceStringType)
	fmt.Println("sliceBoolType:", SliceBoolType)
	fmt.Println("sliceIntType:", SliceIntType)
	fmt.Println("sliceInt8Type:", SliceInt8Type)
	fmt.Println("sliceInt16Type:", SliceInt16Type)
	fmt.Println("sliceInt32Type:", SliceInt32Type)
	fmt.Println("sliceInt64Type:", SliceInt64Type)
	fmt.Println("sliceUintType:", SliceUintType)
	fmt.Println("sliceUint8Type:", SliceUint8Type)
	fmt.Println("sliceUint16Type:", SliceUint16Type)
	fmt.Println("sliceUint32Type:", SliceUint32Type)
	fmt.Println("sliceUint64Type:", SliceUint64Type)
	fmt.Println("sliceFloat32Type:", SliceFloat32Type)
	fmt.Println("sliceFloat64Type:", SliceFloat64Type)
	fmt.Println("sliceComplex64Type:", SliceComplex64Type)
	fmt.Println("sliceComplex128Type:", SliceComplex128Type)
}