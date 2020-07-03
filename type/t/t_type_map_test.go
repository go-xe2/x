package t

import (
	"fmt"
	"reflect"
	"testing"
)

type TestFruit struct {
	Color string `json:"color"`
	Price float32 `json:"price"`
}

type testStructMap struct {
	*TestFruit `json:"testFruit"`
	Name string
	Age int
}

var structMap1 = &testStructMap{
	TestFruit: &TestFruit {
		Color: "red",
		Price: 33.2,
	},
	Name: "my name is structMap1",
	Age:1,
}
func TestType_ToType(t *testing.T) {
	needle := reflect.TypeOf(make(map[string]string))
	fmt.Println("needle type:", needle)

	toStrStrMap := New(333.22).ToType(needle)
	fmt.Println("toStrStrMap:", toStrStrMap, ", type:", reflect.TypeOf(toStrStrMap))

	structToMap := New(structMap1).ToType(needle)
	fmt.Println("structToMap:", structToMap, ", type:", reflect.TypeOf(structToMap))

}
