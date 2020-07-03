package xconv

import (
	"errors"
	"fmt"
	"github.com/go-xe2/x/core/empty"
	"github.com/go-xe2/x/type/xstring"
	"reflect"
	"strings"
)

// 其他类型转换成map[string]interface类型
func Map(value interface{}, tags ...string) map[string]interface{} {
	if value == nil {
		return nil
	}
	if r, ok := value.(map[string]interface{}); ok {
		return r
	} else {
		// Only assert the common combination of types, and finally it uses reflection.
		m := make(map[string]interface{})
		switch value.(type) {
		case map[interface{}]interface{}:
			for k, v := range value.(map[interface{}]interface{}) {
				m[String(k)] = v
			}
		case map[interface{}]string:
			for k, v := range value.(map[interface{}]string) {
				m[String(k)] = v
			}
		case map[interface{}]int:
			for k, v := range value.(map[interface{}]int) {
				m[String(k)] = v
			}
		case map[interface{}]uint:
			for k, v := range value.(map[interface{}]uint) {
				m[String(k)] = v
			}
		case map[interface{}]float32:
			for k, v := range value.(map[interface{}]float32) {
				m[String(k)] = v
			}
		case map[interface{}]float64:
			for k, v := range value.(map[interface{}]float64) {
				m[String(k)] = v
			}
		case map[string]bool:
			for k, v := range value.(map[string]bool) {
				m[k] = v
			}
		case map[string]int:
			for k, v := range value.(map[string]int) {
				m[k] = v
			}
		case map[string]uint:
			for k, v := range value.(map[string]uint) {
				m[k] = v
			}
		case map[string]float32:
			for k, v := range value.(map[string]float32) {
				m[k] = v
			}
		case map[string]float64:
			for k, v := range value.(map[string]float64) {
				m[k] = v
			}
		case map[int]interface{}:
			for k, v := range value.(map[int]interface{}) {
				m[String(k)] = v
			}
		case map[int]string:
			for k, v := range value.(map[int]string) {
				m[String(k)] = v
			}
		case map[uint]string:
			for k, v := range value.(map[uint]string) {
				m[String(k)] = v
			}
		// Not a common type, use reflection
		default:
			rv := reflect.ValueOf(value)
			kind := rv.Kind()
			// If it is a pointer, we should find its real data type.
			if kind == reflect.Ptr {
				rv = rv.Elem()
				kind = rv.Kind()
			}
			switch kind {
			case reflect.Map:
				ks := rv.MapKeys()
				for _, k := range ks {
					m[String(k.Interface())] = rv.MapIndex(k).Interface()
				}
			case reflect.Struct:
				rt := rv.Type()
				name := ""
				tagArray := structTagPriority
				switch len(tags) {
				case 0:
					// No need handle.
				case 1:
					tagArray = append(strings.Split(tags[0], ","), structTagPriority...)
				default:
					tagArray = append(tags, structTagPriority...)
				}
				for i := 0; i < rv.NumField(); i++ {
					// Only convert the public attributes.
					fieldName := rt.Field(i).Name
					if xstring.IsFirstLetterLower(fieldName) || fieldName[0] == '_' {
						continue
					}
					name = ""
					fieldTag := rt.Field(i).Tag
					for _, tag := range tagArray {
						if name = fieldTag.Get(tag); name != "" {
							break
						}
					}
					if name == "" {
						name = strings.TrimSpace(fieldName)
					} else {
						// Support json tag feature: -, omitempty
						name = strings.TrimSpace(name)
						if name == "-" {
							continue
						}
						array := strings.Split(name, ",")
						if len(array) > 1 {
							switch strings.TrimSpace(array[1]) {
							case "omitempty":
								if empty.IsEmpty(rv.Field(i).Interface()) {
									continue
								} else {
									name = strings.TrimSpace(array[0])
								}
							default:
								name = strings.TrimSpace(array[0])
							}
						}
					}
					m[name] = rv.Field(i).Interface()
				}
			default:
				return nil
			}
		}
		return m
	}
}

// struct转换成map[string]interface, 取取内层struct的字段
func MapDeep(value interface{}, tags ...string) map[string]interface{} {
	data := Map(value, tags...)
	for key, value := range data {
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		if kind == reflect.Ptr {
			rv = rv.Elem()
			kind = rv.Kind()
		}
		switch kind {
		case reflect.Struct:
			delete(data, key)
			for k, v := range MapDeep(value, tags...) {
				data[k] = v
			}
		}
	}
	return data
}

func Structs(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return doStructs(params, pointer, false, mapping...)
}

func StructsDeep(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return doStructs(params, pointer, true, mapping...)
}

// 转换[]struct为[]map类型
func doStructs(params interface{}, pointer interface{}, deep bool, mapping ...map[string]string) (err error) {
	if params == nil {
		return errors.New("params cannot be nil")
	}
	if pointer == nil {
		return errors.New("object pointer cannot be nil")
	}
	pointerRt := reflect.TypeOf(pointer)
	if kind := pointerRt.Kind(); kind != reflect.Ptr {
		return fmt.Errorf("pointer should be type of pointer, but got: %v", kind)
	}

	rv := reflect.ValueOf(params)
	kind := rv.Kind()
	if kind == reflect.Ptr {
		rv = rv.Elem()
		kind = rv.Kind()
	}
	switch kind {
	case reflect.Slice, reflect.Array:
		// If <params> is an empty slice, no conversion.
		if rv.Len() == 0 {
			return nil
		}
		array := reflect.MakeSlice(pointerRt.Elem(), rv.Len(), rv.Len())
		itemType := array.Index(0).Type()
		for i := 0; i < rv.Len(); i++ {
			if itemType.Kind() == reflect.Ptr {
				// Slice element is type pointer.
				e := reflect.New(itemType.Elem()).Elem()
				if deep {
					if err = StructDeep(rv.Index(i).Interface(), e, mapping...); err != nil {
						return err
					}
				} else {
					if err = Struct(rv.Index(i).Interface(), e, mapping...); err != nil {
						return err
					}
				}
				array.Index(i).Set(e.Addr())
			} else {
				// Slice element is not type of pointer.
				e := reflect.New(itemType).Elem()

				if deep {
					if err = StructDeep(rv.Index(i).Interface(), e, mapping...); err != nil {
						return err
					}
				} else {
					if err = Struct(rv.Index(i).Interface(), e, mapping...); err != nil {
						return err
					}
				}
				array.Index(i).Set(e)
			}
		}
		reflect.ValueOf(pointer).Elem().Set(array)
		return nil
	default:
		return fmt.Errorf("params should be type of slice, but got: %v", kind)
	}
}
