package t

import (
	"fmt"
	"reflect"
	"testing"
)

type StructA struct {
	Str string
	I int
}

type StructB struct {
	StructA
	F float32
	B bool
}

type StructC struct {
	Str string
	F float32
	B bool
}

var structB = &StructB{
	StructA:StructA{
		Str:"this is structA str",
		I: 22,
	},
	F:33.34,
	B:true,
}
func TestStructDeep(t *testing.T) {
	var params = map[string]interface{}{
		"str22": "this is new value.",
		"i": 1234,
		"f": 55.66,
		"b": false,
		"M": " this is more field.",
	}
	var mapping = map[string]string{
		"str22": "Str",
		"i": "I",
	}
	err := StructDeep(params, structB, mapping)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("structDeep structB:", structB, ", mapping:", mapping)
	err = Struct(map[string]interface{}{
		"str22": "this is map value.",
		"i": 2222,
		"f": 88.88,
		"b": true,
	}, structB, mapping)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("struct structB:", structB)

}

func TestMapDeep(t *testing.T) {
	m := MapDeep(structB)
	fmt.Println("m:", m)
}

func TestType_ToStructType(t *testing.T) {
	var t1 = &StructC{
		Str: "structC",
		F:1111.1,
		B:true,
	}
	var t2 = &t1
	toType := reflect.ValueOf(t1)
	toType2 := reflect.ValueOf(t2)
	fmt.Println("toType:", toType.Type())
	fmt.Println("toType2:", toType2.Type())

	v1 := New(structB).ToStructType(toType.Type())
	fmt.Println("structB convert to structC v1:", v1,", type:", reflect.TypeOf(v1))
	//var tx **StructC
	//tx = v1.(**StructC)
	//fmt.Println("tx:", tx)
	v2 := New(structB).ToType(toType2.Type())
	fmt.Println("structB convert to structC v2:", v2, ", type:", reflect.TypeOf(v2))

}