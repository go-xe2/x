package xclass

import (
	"errors"
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"reflect"
)

// 调用类实例方法
func (o *TObject) Call(method string, params ...interface{}) (results []interface{}, err error) {
	o.initClassVT()
	fnInfo, ok := o.vt.methodMaps[method]
	if ok {
		fn := o.classValue.Method(fnInfo.index)
		if fn.IsValid() {
			values := make([]reflect.Value, 0)
			for _, v := range params {
				values = append(values, reflect.ValueOf(v))
			}
			defer func() {
				if e := recover(); e != nil {
					err = errors.New(fmt.Sprintf("%v", e))
				}
			}()
			resultV := fn.Call(values)
			results = make([]interface{}, len(resultV))
			for k, v := range resultV {
				results[k] = v.Interface()
			}
			return
		}
	}
	return nil, exception.Newf("类%s.%s方法不存在", o.ClassName(), method)
}

// 检查类是否存在某方法
func (o *TObject) HasMethod(method string) bool {
	o.initClassVT()
	if _, ok := o.vt.methodMaps[method]; ok {
		return true
	}
	return false
}

// 获取类对象方法数
func (o *TObject) MethodCount() int {
	o.initClassVT()
	return len(o.vt.methodMaps)
}

func (o *TObject) Methods() []*ClassMethod {
	o.initClassVT()
	if o.methodCache == nil {
		l := len(o.vt.methodMaps)
		i := 0
		result := make([]*ClassMethod, l)
		for k, v := range o.vt.methodMaps {
			result[i] = newClassMethod(k, v.paramTypes, v.resultTypes)
			i++
		}
		o.methodCache = result
	}
	return o.methodCache
}

func (o *TObject) ForEachMethods(fn func(methodName string, params []reflect.Type, results []reflect.Type) bool) {
	o.initClassVT()
	for k, v := range o.vt.methodMaps {
		if !fn(k, v.paramTypes, v.resultTypes) {
			break
		}
	}
}
