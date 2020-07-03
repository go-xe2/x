package xclass

import (
	"github.com/go-xe2/x/sync/xsafeMap"
	"reflect"
)

type TClass struct {
	class     reflect.Type
	classTags ClassTag
}

type Class = *TClass

var classTags = xsafeMap.NewStrAnyMap()

// classTag为class对象设置备注标记
func ClassOf(class Object, tag ...ClassTag) Class {
	clsType := reflect.TypeOf(class)
	for clsType.Kind() == reflect.Ptr {
		clsType = clsType.Elem()
	}
	var classTag ClassTag
	if len(tag) > 0 {
		classTag = tag[0]
	}
	var clsTag ClassTag
	if len(tag) > 0 {
		clsTag = tag[0]
	}
	if clsTag != nil && !classTags.Contains(clsType.String()) {
		classTags.Set(clsType.String(), clsTag)
	}
	cls := &TClass{
		class:     reflect.PtrTo(clsType),
		classTags: classTag,
	}
	return cls
}

func (cls *TClass) Create(props ...interface{}) interface{} {
	return Create(cls, props...)
}

func (cls *TClass) Type() reflect.Type {
	return cls.class
}

func (cls *TClass) Tag() ClassTag {
	className := cls.class.Elem().String()
	if classTags.Contains(className) {
		return classTags.Get(className).(ClassTag)
	}
	return NewClassTag("")
}

func GetClassTag(classType reflect.Type) ClassTag {
	clsT := classType
	for clsT.Kind() == reflect.Ptr {
		clsT = clsT.Elem()
	}
	key := clsT.String()
	if classTags.Contains(key) {
		return classTags.Get(key).(ClassTag)
	}
	return NewClassTag("")
}
