package structs

import (
	"reflect"
)

// tag映射字段名
func TagMapName(structPtr interface{}, priority []string, recursive bool) (map[string]string, error) {
	tagMap, err := TagMapField(structPtr, priority, recursive)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string, len(tagMap))
	for k, v := range tagMap {
		m[k] = v.Name()
	}
	return m, nil
}

// 获取priority中的tagName映射字段
func TagMapField(structPtr interface{}, priority []string, recursive bool) (map[string]IField, error) {
	tagMap := make(map[string]IField)
	structObj, err := New(structPtr)
	if err != nil {
		return nil, err
	}
	fields := structObj.Fields()
	tag := ""
	for _, field := range fields {
		tag = ""
		for _, p := range priority {
			tag = field.Tag().Get(p).Value()
			if tag != "" {
				break
			}
		}
		if tag != "" {
			tagMap[tag] = field
		}
		if recursive {
			rv := field.Type()
			kind := rv.Kind()
			if kind == reflect.Ptr {
				rv = rv.Elem()
				kind = rv.Kind()
			}
			if kind == reflect.Struct {
				mMaps, err := TagMapField(rv, priority, true)
				if err != nil {
					continue
				}
				for k, v := range mMaps {
					if _, ok := tagMap[k]; !ok {
						tagMap[k] = v
					}
				}
			}
		}
	}
	return tagMap, nil
}
