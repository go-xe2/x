package xclass

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTObject(t *testing.T) {
	//o := &TObject{}
	//o := Create(ObjectClass).(*TObject)
	//fmt.Println("after init o.dd:", o.dd)
	//fmt.Println("after init o.TBase.dd:", o.TBase.dd)
	//o.dd = "fdfdfdf"
	//fmt.Println("initObject o:", o)
	//fmt.Println("o.dd:", o.dd)
	//fmt.Println("o.TBase.dd", o.TBase.dd)
	//fmt.Println("o.hello:", o.Hello())

	o1 := Create(TDemo1Class).(*TDemo1)
	fmt.Println(o1)
	o1.dd = "123456"
	fmt.Println("o1 hello:", o1.SayHello())

	o2 := o1.NewInstance().(*TDemo1)
	fmt.Println("o2:", o2)
	o2.dd = "abcdefg"
	fmt.Println("o2 hello:", o2.SayHello())
	fmt.Println("o1.this:", reflect.TypeOf(o1.this))
	fmt.Println("o1.TObject.this:", reflect.TypeOf(o1.TObject.this))

	o3 := Create(TDemo2Class).(*TDemo2)
	fmt.Println("o3:", o3)

}
