package t

import (
	"reflect"
)

func (t Type) Map() TMap {
	ref := reflect.Indirect(reflect.ValueOf(t.val))
	var res = make(TMap)
	keys := ref.MapKeys()

	for _, item := range keys {
		res[New(item.Interface())] = New(ref.MapIndex(item).Interface())
	}
	return res
}

func (t Type) MapAny() TMapAny {
	m := t.Map()
	var res = make(TMapAny)
	for k, v := range m {
		res[k.Any()] = v
	}
	return res
}

func (t Type) MapString() TMapString {
	m := t.Map()
	var res = make(TMapString)
	for k, v := range m {
		res[k.String()] = v
	}
	return res
}

func (t Type) MapStrAny() map[string]interface{} {
	m := t.Map()
	var res = make(map[string]interface{})
	for k, v := range m {
		res[k.String()] = v.Any()
	}
	return res
}

func (t Type) Hash() THash {
	m := t.Map()
	var results = make(THash)
	for k, v := range m {
		results[k.String()] = v
	}
	return results
}

func (t Type) MapInt64() TMapInt64 {
	m := t.Map()
	var res = make(TMapInt64)
	for k, v := range m {
		res[k.Int64()] = v
	}
	return res
}

func (t Type) ToMapType(mapType reflect.Type) interface{} {
	val := t.val
	srcType := reflect.TypeOf(val)
	srcElem := srcType
	for srcElem.Kind() == reflect.Ptr {
		srcElem = srcElem.Elem()
	}
	switch srcElem.Kind() {
	case reflect.Slice, reflect.Array:
		// slice转map
		items := t.SliceAny()
		tmpMap := reflect.MakeMap(reflect.MapOf(IntType, srcElem.Elem()))
		for k, v := range items {
			tmpMap.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
		}
		val = tmpMap.Interface()
		break
	case reflect.Struct:
		// struct转map
		val = Map(t.val)
		break
	default:
		// 其他类型
		// 如果不是map类型，转换成map类型后处理
		tmpMap := reflect.MakeMap(reflect.MapOf(StringType, srcType))
		tmpMap.SetMapIndex(reflect.ValueOf(srcElem.Name()), reflect.ValueOf(val))
		val = tmpMap.Interface()
	}
	ret := reflect.MakeMap(mapType)
	keyType := mapType.Key()
	valType := mapType.Elem()
	ref := reflect.Indirect(reflect.ValueOf(val))
	keys := ref.MapKeys()
	for _, item := range keys {
		key := New(item.Interface())
		val := New(ref.MapIndex(item).Interface())
		var keyValue, valValue reflect.Value
		keyValue = reflect.ValueOf(valToType(key, keyType))
		valValue = reflect.ValueOf(valToType(val, valType))
		ret.SetMapIndex(keyValue, valValue)
	}
	return ret.Interface()
}
