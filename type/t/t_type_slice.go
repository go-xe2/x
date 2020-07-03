package t

import (
	"reflect"
)

func (t Type) Slice() TSlice {
	ref := reflect.Indirect(reflect.ValueOf(t.val))
	l := ref.Len()
	v := ref.Slice(0, l)
	var res = TSlice{}
	for i := 0; i < l; i++ {
		res = append(res, New(v.Index(i).Interface()))
	}
	return res
}

func (t Type) SliceAny() []interface{} {
	s := t.Slice()
	var res = make([]interface{}, 0)
	for _, item := range s {
		res = append(res, item.Any())
	}
	return res
}

func (t Type) SliceMapStrAny() []map[string]interface{} {
	s := t.Slice()
	var res = make([]map[string]interface{}, 0)
	for _, item := range s {
		res = append(res, item.MapStrAny())
	}
	return res
}

func (t Type) SliceHash() []THash {
	s := t.Slice()
	var results []THash
	for _, item := range s {
		results = append(results, item.Hash())
	}
	return results
}

func (t Type) SliceString() []string {
	s := t.Slice()
	var res = make([]string, 0)
	for _, item := range s {
		res = append(res, item.String())
	}
	return res
}

func (t Type) SliceInt64() []int64 {
	s := t.Slice()
	var res = make([]int64, 0)
	for _, item := range s {
		res = append(res, item.Int64())
	}
	return res
}

func InArray(needle, arr interface{}) bool {
	nt := New(needle)
	for _, item := range New(arr).Slice() {
		if nt.String() == item.String() {
			return true
		}
	}
	return false
}

func (t Type) ToSliceType(sliceType reflect.Type) interface{} {
	val := t.val
	valType := reflect.TypeOf(val)
	valElem := valType
	for valElem.Kind() == reflect.Ptr {
		valElem = valElem.Elem()
	}
	if valElem.Kind() != reflect.Slice && valElem.Kind() != reflect.Array {
		// 如果不是slice类型值，则创建
		tmpSliceType := reflect.SliceOf(valType)
		tmp := reflect.MakeSlice(tmpSliceType, 0, 1)
		tmp = reflect.Append(tmp, reflect.ValueOf(val))
		val = tmp.Interface()
	}
	ref := reflect.Indirect(reflect.ValueOf(val))
	l := ref.Len()
	v := ref.Slice(0, l)
	var res = reflect.MakeSlice(sliceType, l, l)
	elemTyp := sliceType.Elem()
	for i := 0; i < l; i++ {
		item := New(v.Index(i).Interface())
		toVal := reflect.ValueOf(valToType(item, elemTyp))
		res.Index(i).Set(toVal)
	}
	return res.Interface()
}
