package t

import (
	"reflect"
)

func (t Type) ToStructType(structType reflect.Type) interface{}  {
	values := MapDeep(t.val)
	vt := structType
	for vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	objV := reflect.New(vt).Interface()
	if err := Struct(values, objV); err != nil {
		panic(err)
	}
	ele := reflect.ValueOf(objV)
	for ele.Kind() == reflect.Ptr {
		ele = ele.Elem()
	}
	vt = structType
	ret := ele.Interface()
	for vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
		ret = any2anyPointer(ret)
	}
	return ret
}
