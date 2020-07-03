package xclass

import (
	"errors"
	"fmt"
	"github.com/go-xe2/x/core/exception"
	_type "github.com/go-xe2/x/sync/type"
	"reflect"
)

// 获取类对像字段数
func (o *TObject) FieldCount() int {
	o.initClassVT()
	return len(o.vt.fieldMaps)
}

func (o *TObject) Fields() []*ClassField {
	o.initClassVT()
	if o.fieldCache == nil {
		l := len(o.vt.fieldMaps)
		result := make([]*ClassField, l)
		i := 0
		for k, v := range o.vt.fieldMaps {
			result[i] = newClassField(k, v.fieldTag, v.fieldType)
			i++
		}
		o.fieldCache = result
	}
	return o.fieldCache
}

// 检查类是否存在某字段
func (o *TObject) HasField(fieldName string) bool {
	o.initClassVT()
	if _, ok := o.vt.fieldMaps[fieldName]; ok {
		return true
	}
	return false
}

func (o *TObject) HasDynamicField(fieldName string) bool {
	if _, ok := o.dynamicFields[fieldName]; ok {
		return true
	}
	return false
}

func (o *TObject) ForEachFields(fn func(fieldName string, tag ClassTag, fieldType reflect.Type) bool) {
	o.initClassVT()
	for k, v := range o.vt.fieldMaps {
		if !fn(k, v.fieldTag, v.fieldType) {
			break
		}
	}
}

// 设置字段值
func (o *TObject) Set(fieldName string, value interface{}) (err error) {
	fv, _, _ := o.getFieldValueByName(fieldName)
	if !fv.IsValid() {
		return exception.Newf("类%s.%s字段未定义", o.ClassName(), fieldName)
	}
	if !fv.CanSet() {
		return exception.Newf("类%s.%s字段只读", o.ClassName(), fieldName)
	}
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()
	oldV := fv.Interface()
	if reflect.DeepEqual(oldV, value) {
		// 新值与旧值相同时不设置
		return
	}
	var canSetValue = true
	if observer, ok := o.this.(ObjectFieldObserver); ok {
		canSetValue = observer.OnFieldChanged(fieldName, oldV, value)
	}
	if canSetValue {
		fv.Set(reflect.ValueOf(value))
	}
	return
}

// 数据类型自动转换的安全设置字段值
func (o *TObject) SafeSet(fieldName string, value interface{}) (err error) {
	fv, fo, fi := o.getFieldValueByName(fieldName)
	if !fv.IsValid() {
		return exception.Newf("类%s.%s字段未定义", o.ClassName(), fieldName)
	}
	if fi.setter != nil {
		oldV := fv.Interface()
		if reflect.DeepEqual(oldV, value) {
			// 值没有发生变化时不处理
			return nil
		}
		var canChanged = true
		if observer, ok := o.this.(ObjectFieldObserver); ok {
			canChanged = observer.OnFieldChanged(fieldName, oldV, value)
		}
		if canChanged {
			if err := fi.setter(fo, fi, value); err != nil {
				return err
			}
		}
		return nil
	}
	return exception.Newf("类%s.%s字段数据类型不支持动态设置", o.ClassName(), fieldName)
}

// 获取字段值
func (o *TObject) Get(fieldName string) interface{} {
	fv, _, _ := o.getFieldValueByName(fieldName)
	if fv.IsValid() {
		return fv.Interface()
	}
	return nil
}

func (o *TObject) GetVar(fieldName string) *_type.TVar {
	return _type.NewVar(o.Get(fieldName))
}

// 设置动态字段值，如果字段不存在则不进行设置
func (o *TObject) DynamicSet(fieldName string, value interface{}) interface{} {
	if v, ok := o.dynamicFields[fieldName]; ok {
		if reflect.DeepEqual(v, value) {
			// 值没有发生变化时不设置
			return v
		}
		var canSet = true
		if observer, ok := o.this.(ObjectDynamicFieldObserver); ok {
			canSet = observer.OnDynamicFieldChanged(fieldName, v, value)
		}
		if canSet {
			o.dynamicFields[fieldName] = value
		}
		return v
	}
	return nil
}

// 字段存不存在都设置新值
func (o *TObject) DynamicTrySet(fieldName string, value interface{}) interface{} {
	old := o.dynamicFields[fieldName]
	if reflect.DeepEqual(old, value) {
		return old
	}
	var canSet = true
	if observer, ok := o.this.(ObjectDynamicFieldObserver); ok {
		canSet = observer.OnDynamicFieldChanged(fieldName, old, value)
	}
	if canSet {
		o.dynamicFields[fieldName] = value
	}
	return old
}

func (o *TObject) DynamicGetOrSet(fieldName string, setValue interface{}) interface{} {
	if v, ok := o.dynamicFields[fieldName]; ok {
		return v
	}
	newValue := setValue
	// 不存在时设置
	if fn, ok := setValue.(func() interface{}); ok {
		newValue = fn()
	}
	canSet := true
	if observer, ok := o.this.(ObjectDynamicFieldObserver); ok {
		canSet = observer.OnDynamicFieldChanged(fieldName, nil, newValue)
	}
	if canSet {
		o.dynamicFields[fieldName] = newValue
	}
	return newValue
}

func (o *TObject) DynamicGetOrSetVar(fieldName string, setValue interface{}) *_type.TVar {
	return _type.NewVar(o.DynamicGetOrSet(fieldName, setValue))
}

// 获取动态字段值
func (o *TObject) DynamicGet(fieldName string) interface{} {
	if v, ok := o.dynamicFields[fieldName]; ok {
		return v
	}
	return nil
}

func (o *TObject) DynamicGetVar(fieldName string) *_type.TVar {
	return _type.NewVar(o.DynamicGet(fieldName))
}

func (o *TObject) DynamicFieldCount() int {
	return len(o.dynamicFields)
}

func (o *TObject) ForEachDynamicFields(fn func(k string, v interface{}) bool) {
	for k, v := range o.dynamicFields {
		if !fn(k, v) {
			break
		}
	}
}
