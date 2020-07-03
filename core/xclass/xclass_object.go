package xclass

import (
	"github.com/go-xe2/x/core/exception"
	"reflect"
)

type testTag string

type TObject struct {
	fieldCache    []*ClassField  // 字段缓存列表
	methodCache   []*ClassMethod // 类方法缓存列表
	vt            *classVT
	classValue    reflect.Value                      // 类对象reflect.value
	this          interface{}                        // 类实例指针
	extends       map[reflect.Type]*classExtendField // 继承父类列表
	dynamicFields map[string]interface{}             // 动态字段
	classTag      ClassTag                           // 类tag注解
}

var _ Object = (*TObject)(nil)

var TObjectClass = ClassOf((*TObject)(nil))

func (o *TObject) Constructor(props ...interface{}) interface{} {
	if o.this == nil {
		panic(exception.New("class类未正确初始化，请使用Create方法创建实例"))
	}
	return o.this
}

func (o *TObject) Destroy() {
	o.dynamicFields = nil
	o.extends = nil
	o.vt = nil
	o.methodCache = nil
	o.fieldCache = nil
	o.this = nil
	o.classValue = reflect.Value{}
}

func (o *TObject) This() interface{} {
	return o.this
}

func (o *TObject) NewInstance(props ...interface{}) interface{} {
	vt := classToClassVT(o.classValue.Type())
	return classAlloc(vt, props...)
}
