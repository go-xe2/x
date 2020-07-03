package structs

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/xstring"
	"reflect"
)

const StrCallParamNumErrorMsg = "%s.%s需要%d个参数，实际传入%d个参数,参数类型不匹配"
const StrCallParamTypeErrorMsg = "%s.%s第%d个形参为%s类型，调用实参为%s类型,参数类型不匹配"

type MethodParams struct {
	owner *MethodInfo
	items []reflect.Type
}

func NewMethodParams(mi *MethodInfo, paramCount int) *MethodParams {
	return &MethodParams{
		owner: mi,
		items: make([]reflect.Type, paramCount),
	}
}

func (mp *MethodParams) String() string {
	count := len(mp.items)
	str := make([]string, count)
	for i := 0; i < count; i++ {
		str[i] = mp.items[i].Kind().String()
	}
	return xstring.Join(str, ",")
}

func (mp *MethodParams) Items() []reflect.Type {
	return mp.items
}

func (mp *MethodParams) Count() int {
	return len(mp.items)
}

func (mp *MethodParams) Add(item reflect.Type) {
	mp.items = append(mp.items, item)
}

func (mp *MethodParams) Set(index int, item reflect.Type) {
	mp.items[index] = item
}

func (mp *MethodParams) Clear() {
	mp.items = make([]reflect.Type, 0)
}

func (mp *MethodParams) CheckValues(instance interface{}, values ...interface{}) (bool, error) {
	count := mp.Count()
	preParams := make([]interface{}, 1)
	preParams[0] = instance
	preParams = append(preParams, values...)
	inCount := len(preParams)

	if inCount != count {
		return false, exception.Newf(StrCallParamNumErrorMsg, mp.owner.OwnerName(), mp.owner.name, count, inCount)
	}
	for i := 0; i < count; i++ {
		inType := reflect.TypeOf(preParams[i])
		rdType := mp.items[i]
		if rdType.Kind() != inType.Kind() {
			return false, exception.Newf(StrCallParamTypeErrorMsg, mp.owner.OwnerName(), mp.owner.name, i, rdType.Name(), inType.Name())
		}
		for inType.Kind() == reflect.Ptr {
			inType = inType.Elem()
		}
		for rdType.Kind() == reflect.Ptr {
			rdType = rdType.Elem()
		}
		if inType.Kind() != rdType.Kind() {
			return false, exception.Newf(StrCallParamTypeErrorMsg, mp.owner.OwnerName(), mp.owner.name, i, rdType.Name(), inType.Name())
		}
	}
	return true, nil
}
