package structs

import (
	"github.com/go-xe2/x/type/xstring"
	"reflect"
)

func MapField(structPtr interface{}, priority []string, recursive bool) (map[string]IField, error) {
	result := make(map[string]IField)
	obj, err := New(structPtr)
	if err != nil {
		return nil, err
	}
	fields := obj.Fields()
	var fieldName = ""
	var fieldTag = ""
	for _, field := range fields {
		fieldName = field.Name()
		if xstring.IsFirstLetterLower(fieldName) {
			continue
		}
		result[fieldName] = field
		fieldTag = ""
		tags := field.Tag()
		for _, p := range priority {
			fieldTag = tags.Get(p).Value()
			if fieldTag != "" && fieldTag != "-" {
				break
			}
		}
		result[fieldTag] = field
		if recursive {
			rv := reflect.ValueOf(field.Type())
			kind := rv.Kind()
			if kind == reflect.Ptr {
				rv = rv.Elem()
				kind = rv.Kind()
			}
			if kind == reflect.Struct {
				mFields, err := MapField(rv, priority, true)
				if err != nil {
					continue
				}
				for k, v := range mFields {
					if _, ok := result[k]; !ok {
						result[k] = v
					}
				}
			}
		}
	}
	return result, nil
}
