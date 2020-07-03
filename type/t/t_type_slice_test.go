package t

import (
	"fmt"
	"reflect"
	"testing"
)

var slice1 = []int{
	1,2,3,4,5,6,
}

var slice2 = []bool {
	true,
	false,
	false,
	true,
}

var slice3 = []float32{
	1.1,
	1.2,
	1.3,
	1.4,
	1.5,
}

var slice5 = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
	"6.1",
	"6.2",
	"6.3",
}

var slice6 = [][]int{
	{1,2},
	{1,3},
	{1,4},
	{1,5},
}

var slice7 = []map[int]string{
	{
		1: "1.2",
		2: "2.2",
		3: "3.3",
	},
}


func TestType_ToSliceType(t *testing.T) {

	strFloatMapSlice := reflect.TypeOf(make([]map[string]float64, 0))
	fmt.Println("strFloatMapSlice:", strFloatMapSlice)


	int2strSlice := New(slice1).ToSliceType(SliceStringType)
	fmt.Println("int2strSlice:", int2strSlice, ", type:", reflect.TypeOf(int2strSlice), "old type:", reflect.TypeOf(slice1))

	bool2strSlice := New(slice2).ToSliceType(SliceStringType)
	fmt.Println("bool2strSlice:", bool2strSlice, ", type:", reflect.TypeOf(bool2strSlice), ", old type:", reflect.TypeOf(slice2))

	float2strSlice := New(slice3).ToSliceType(SliceStringType)
	fmt.Println("float2strSlice:", float2strSlice, ", type:", reflect.TypeOf(float2strSlice), ", old type:", reflect.TypeOf(slice3))

	intint2strSlice := New(slice6).ToSliceType(SliceStringType)
	fmt.Println("intint2strSlice:", intint2strSlice, ", type:", reflect.TypeOf(intint2strSlice), ", old type:", reflect.TypeOf(slice6))

	str2intSlice := New(slice5).ToSliceType(SliceIntType)
	fmt.Println("str2intSlice:", str2intSlice, ", type:", reflect.TypeOf(str2intSlice), ", old type:", reflect.TypeOf(slice5))

	str2floatSlice := New(slice5).ToSliceType(SliceFloat32Type)
	fmt.Println("str2floatSlice:", str2floatSlice, ", type:", reflect.TypeOf(str2floatSlice), ", old type:", reflect.TypeOf(slice5))

	intint2AnySlice := New(slice6).ToSliceType(SliceAnyType)
	fmt.Println("intint2AnySlice:", intint2AnySlice, ", type:", reflect.TypeOf(intint2AnySlice), ", old type:", reflect.TypeOf(slice6))

	intStrMap2strFloatMap := New(slice7).ToSliceType(strFloatMapSlice)
	fmt.Println("intStrMap2strFloatMap:", intStrMap2strFloatMap, ", type:", reflect.TypeOf(intStrMap2strFloatMap), ", old type:", reflect.TypeOf(slice7))

	intStrMap2AnySlice := New(slice7).ToSliceType(SliceAnyType)
	fmt.Println("intStrMap2AnySlice:", intStrMap2AnySlice, ", type:", reflect.TypeOf(intStrMap2AnySlice), ", old type:", reflect.TypeOf(slice7))

	toIntInt2AnySlice := New(slice6).ToType(SliceAnyType)
	fmt.Println("toIntInt2AnySlice:", toIntInt2AnySlice, ", type:", reflect.TypeOf(toIntInt2AnySlice))

	toIntInt2StrSlice := New(slice6).ToType(SliceStringType)
	fmt.Println("toIntInt2StrSlice:", toIntInt2StrSlice, ", type:", reflect.TypeOf(toIntInt2StrSlice))

	sg2StrSlice := New(3.323).ToType(SliceStringType)
	fmt.Println("sg2StrSlice:", sg2StrSlice, ", type:", reflect.TypeOf(sg2StrSlice))
}
