package xclass

import (
	"fmt"
	"reflect"
	"testing"
)

var demoRt = &classVT{
	clsType: TDemo2Class.Type(),
	extends: []*classVTExtendField{
		{
			fieldIndex: []int{0, 1},
			extendType: &classVT{
				clsType: TDemo1Class.Type(),
				extends: []*classVTExtendField{
					{
						fieldIndex: []int{1, 2},
						extendType: &classVT{
							clsType: TObjectClass.Type(),
							extends: nil,
						},
					},
				},
			},
		},
	},
}

func TestAlloc(t *testing.T) {
	clsVt := classToClassVT(TDemo2Class.Type())
	if clsVt == nil {
		t.Error("get class virtual table error")
	}
	fmt.Println("clsVT:")
	fmt.Print(clsVt.String())

	clsTree := classAllocTree(clsVt)
	fmt.Print(clsTree)

	inst := classAlloc(clsVt)
	fmt.Println("inst:", inst)
	obj := inst.(*TDemo2)
	fmt.Println("obj:", obj)
	obj.cc = "this is init cc value"
	fmt.Println("obj sayHello:", obj.SayHello())
	fmt.Println("obj.this:", obj.This())
	fmt.Println("obj.className:", obj.ClassName())
	fmt.Println("obj.classType:", obj.ClassType())

	obj2 := obj.NewInstance().(*TDemo2)
	fmt.Println("obj2:", obj2)
	obj2.cc = "this is obj2 assign value"
	fmt.Println("obj2.sayHello:", obj2.SayHello())
	fmt.Println("obj2.Parents:", obj2.extends)
	fmt.Println("obj2.implements TDemo1:", obj2.Implements(TDemo1Class))
	fmt.Println("obj2.implements TDemo2:", obj2.Implements(TDemo2Class))
	fmt.Println("obj2.implements TObject:", obj2.Implements(TObjectClass))
	fmt.Println("obj2.fieldCount:", obj2.FieldCount())
	fmt.Println("obj2.methodCount:", obj2.MethodCount())
	fmt.Println("obj2.tag:", obj2.ClassTag())
	fmt.Println("obj2 fields:")
	obj2.ForEachFields(func(fieldName string, tag ClassTag, fieldType reflect.Type) bool {
		fmt.Println("obj2.field:", fieldName, ", tag:", tag, ", type:", fieldType)
		return true
	})
	fmt.Println("obj2.fields:", obj2.Fields())

	fmt.Println("obj2.methods:", obj2.Methods())

	var objectInst = obj2.ToParent(TObjectClass).(*TObject)
	fmt.Println("obj2 assign to TObject:", objectInst)
	var demo1Inst = obj2.ToParent(TDemo1Class).(*TDemo1)
	fmt.Println("obj2 assign to TDemo1:", demo1Inst)
	var obj3 = demo1Inst.This().(*TDemo2)
	fmt.Println("obj2.TDemo1 assign to TDemo2:", obj3)

	// test get|set
	obj2.AA = "this is obj2, AA declare in demo1"
	fmt.Println("obj2.AA:", obj2.AA)
	fmt.Println("obj2.get(AA):", obj2.Get("AA"))
	fmt.Println("obj2.set(dd, 123456) err:", obj2.Set("dd", 123456))
	fmt.Println("obj2.set(dd, '123456') err:", obj2.Set("dd", "123456"))
	fmt.Println("obj2.get(dd):", obj2.Get("dd"))
	fmt.Println("obj2.get(this):", obj2.Get("this"))
	fmt.Println("obj2.get(cc):", obj2.Get("cc"))
	fmt.Println("obj2.set(Sex) err:", obj2.Set("Sex", "男"))
	fmt.Println("obj2.set(Sex) err:", obj2.Set("Sex", "女"))
	fmt.Println("obj2.get(Sex):", obj2.Get("Sex"))
	fmt.Println("obj2.set(Sex) err:", obj2.Set("Sex", true))
	fmt.Println("obj2.SafeSet(Sex) err:", obj2.SafeSet("Sex", true))
	fmt.Println("obj2.sex:", obj2.Sex)
	fmt.Println("obj2.set(Sex) err:", obj2.Set("Sex", 123))
	fmt.Println("obj2.SafeSet(Sex) err:", obj2.SafeSet("Sex", 123))
	fmt.Println("obj2.sex:", obj2.Sex)
	fmt.Println("obj2.set(Sex) err:", obj2.Set("Sex", 22.333))
	fmt.Println("obj2.SafeSet(Sex) err:", obj2.SafeSet("Sex", 22.333))
	fmt.Println("obj2.Sex:", obj2.Sex)
	fmt.Println("obj2.SafeSet(AA) err:", obj2.SafeSet("AA", "dd2323232322"))
	fmt.Println("obj2.AA:", obj2.AA)

	s, err := obj2.Call("SayHello")
	if err != nil {
		t.Error("call SayHello err:", err)
	}
	fmt.Println("call SayHello result:", s)
	s, err = obj2.Call("HelloName", " this is call param")
	fmt.Println("call HelloName result:", s, ", err:", err)

}
