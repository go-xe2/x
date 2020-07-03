package structs

import (
	"errors"
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"reflect"
)

type MethodInfo struct {
	instance  interface{}
	instValue reflect.Value
	name      string
	index     int
	inParams  *MethodParams
	outParams *MethodParams
}

func NewMethodInfo(instance interface{}, instValue reflect.Value, name string, idx int) *MethodInfo {
	return &MethodInfo{
		instance:  instance,
		instValue: instValue,
		name:      name,
		index:     idx,
	}
}

func (mi *MethodInfo) String() string {
	inParams := mi.inParams.String()
	outParams := mi.outParams.String()
	if mi.outParams.Count() > 1 {
		outParams = "(" + outParams + ")"
	}
	return fmt.Sprintf("%s.%s(%s) %s", mi.OwnerName(), mi.name, inParams, outParams)
}

func (mi *MethodInfo) Name() string {
	return mi.name
}

func (mi *MethodInfo) Index() int {
	return mi.index
}

func (mi *MethodInfo) OutParams() *MethodParams {
	return mi.outParams
}

func (mi *MethodInfo) InParams() *MethodParams {
	return mi.inParams
}

func (mi *MethodInfo) OwnerName() string {
	return mi.instValue.Type().Name()
}

func (mi *MethodInfo) SetInParams(params *MethodParams) {
	mi.inParams = params
}

func (mi *MethodInfo) SetOutParams(params *MethodParams) {
	mi.outParams = params
}

// 调用方法
func (mi *MethodInfo) Invoke(params ...interface{}) (result []interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()
	result = make([]interface{}, 0)
	fn := mi.instValue.Type().Method(mi.index)

	isValid, err := mi.inParams.CheckValues(mi.instance, params...)
	if err != nil {
		return
	}
	if !isValid {
		err = exception.New("调用参数不正确")
		return
	}

	callParams := make([]reflect.Value, 1)
	callParams[0] = reflect.ValueOf(mi.instance)
	for _, p := range params {
		callParams = append(callParams, reflect.ValueOf(p))
	}
	numOut := mi.outParams.Count()
	res := fn.Func.Call(callParams)
	result = make([]interface{}, numOut)
	for i := 0; i < numOut; i++ {
		result[i] = res[i].Interface()
	}
	return
}
