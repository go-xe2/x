package xclass

import (
	"fmt"
)

type TDemo1 struct {
	*TObject
	dd string
	AA string
}

type TDemo2 struct {
	*TDemo1
	cc  string
	Sex string
}

var TDemo1Class = ClassOf((*TDemo1)(nil))
var TDemo2Class = ClassOf((*TDemo2)(nil), NewClassTagByMap(map[string]string{
	"filter":  "11111",
	"summary": "this is Demo2 class tag test",
}))

func (d *TDemo1) Constructor(props ...interface{}) interface{} {
	d.TObject.Constructor(props...)
	fmt.Println("TDemo1 constructor call.")
	return d.this
}

func (d *TDemo1) Hello() string {
	return "my define in TDemo1, dd:" + d.dd
}

func (d2 *TDemo2) Constructor(props ...interface{}) interface{} {
	d2.TDemo1.Constructor(props...)
	fmt.Println("TDemo2 constructor call. class tag:", d2.ClassTag())
	return d2.this
}

func (d2 *TDemo2) Hello() string {
	return "this is TDemo2 cc:" + d2.cc
}

func (d2 *TDemo2) HelloName(name string) string {
	d2.cc = name
	return d2.Hello()
}

func (d2 *TDemo2) OnFieldChanged(fieldName string, oldV, newV interface{}) bool {
	fmt.Println("====TDemo2 field:", fieldName, " changed old value:", oldV, ", value:", newV)
	return true
}
