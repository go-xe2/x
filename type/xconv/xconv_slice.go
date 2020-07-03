package xconv

import (
	"github.com/go-xe2/x/type/xstring"
	"reflect"
)

// SliceInt is alias of Ints.
func SliceInt(i interface{}) []int {
	return Ints(i)
}

// SliceStr is alias of Strings.
func SliceStr(i interface{}) []string {
	return Strings(i)
}

// SliceAny is alias of Interfaces.
func SliceAny(i interface{}) []interface{} {
	return Interfaces(i)
}

// SliceFloat is alias of Floats.
func SliceFloat(i interface{}) []float64 {
	return Floats(i)
}

// SliceMap is alias of Maps.
func SliceMap(i interface{}) []map[string]interface{} {
	return Maps(i)
}

// SliceMapDeep is alias of MapsDeep.
func SliceMapDeep(i interface{}) []map[string]interface{} {
	return MapsDeep(i)
}

// SliceStruct is alias of Structs.
func SliceStruct(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return Structs(params, pointer, mapping...)
}

// SliceStructDeep is alias of StructsDeep.
func SliceStructDeep(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return StructsDeep(params, pointer, mapping...)
}

// Ints converts <i> to []int.
func Ints(i interface{}) []int {
	if i == nil {
		return nil
	}
	if r, ok := i.([]int); ok {
		return r
	} else {
		array := make([]int, 0)
		switch value := i.(type) {
		case []string:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []int8:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []int16:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []int32:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []int64:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []uint:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []uint8:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []uint16:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []uint32:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []uint64:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []bool:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []float32:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []float64:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case []interface{}:
			for _, v := range value {
				array = append(array, Int(v))
			}
		case [][]byte:
			for _, v := range value {
				array = append(array, Int(v))
			}
		default:
			return []int{Int(i)}
		}
		return array
	}
}

// Strings converts <i> to []string.
func Strings(i interface{}) []string {
	if i == nil {
		return nil
	}
	if r, ok := i.([]string); ok {
		return r
	} else {
		array := make([]string, 0)
		switch value := i.(type) {
		case []int:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []int8:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []int16:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []int32:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []int64:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []uint:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []uint8:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []uint16:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []uint32:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []uint64:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []bool:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []float32:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []float64:
			for _, v := range value {
				array = append(array, String(v))
			}
		case []interface{}:
			for _, v := range value {
				array = append(array, String(v))
			}
		case [][]byte:
			for _, v := range value {
				array = append(array, String(v))
			}
		default:
			return []string{String(i)}
		}
		return array
	}
}

// Strings converts <i> to []float64.
func Floats(i interface{}) []float64 {
	if i == nil {
		return nil
	}
	if r, ok := i.([]float64); ok {
		return r
	} else {
		array := make([]float64, 0)
		switch value := i.(type) {
		case []string:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []int:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []int8:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []int16:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []int32:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []int64:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []uint:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []uint8:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []uint16:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []uint32:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []uint64:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []bool:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []float32:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		case []interface{}:
			for _, v := range value {
				array = append(array, Float64(v))
			}
		default:
			return []float64{Float64(i)}
		}
		return array
	}
}

// Interfaces converts <i> to []interface{}.
func Interfaces(i interface{}) []interface{} {
	if i == nil {
		return nil
	}
	if r, ok := i.([]interface{}); ok {
		return r
	} else {
		array := make([]interface{}, 0)
		switch value := i.(type) {
		case []string:
			for _, v := range value {
				array = append(array, v)
			}
		case []int:
			for _, v := range value {
				array = append(array, v)
			}
		case []int8:
			for _, v := range value {
				array = append(array, v)
			}
		case []int16:
			for _, v := range value {
				array = append(array, v)
			}
		case []int32:
			for _, v := range value {
				array = append(array, v)
			}
		case []int64:
			for _, v := range value {
				array = append(array, v)
			}
		case []uint:
			for _, v := range value {
				array = append(array, v)
			}
		case []uint8:
			for _, v := range value {
				array = append(array, v)
			}
		case []uint16:
			for _, v := range value {
				array = append(array, v)
			}
		case []uint32:
			for _, v := range value {
				array = append(array, v)
			}
		case []uint64:
			for _, v := range value {
				array = append(array, v)
			}
		case []bool:
			for _, v := range value {
				array = append(array, v)
			}
		case []float32:
			for _, v := range value {
				array = append(array, v)
			}
		case []float64:
			for _, v := range value {
				array = append(array, v)
			}
		default:
			// Finally we use reflection.
			rv := reflect.ValueOf(i)
			kind := rv.Kind()
			// If it's pointer, find the real type.
			if kind == reflect.Ptr {
				rv = rv.Elem()
				kind = rv.Kind()
			}
			switch kind {
			case reflect.Slice, reflect.Array:
				for i := 0; i < rv.Len(); i++ {
					array = append(array, rv.Index(i).Interface())
				}
			case reflect.Struct:
				rt := rv.Type()
				for i := 0; i < rv.NumField(); i++ {
					// Only public attributes.
					if xstring.IsFirstLetterLower(rt.Field(i).Name) {
						continue
					}
					array = append(array, rv.Field(i).Interface())
				}
			default:
				return []interface{}{i}
			}
		}
		return array
	}
}

// Maps converts <i> to []map[string]interface{}.
func Maps(value interface{}, tags ...string) []map[string]interface{} {
	if value == nil {
		return nil
	}
	if r, ok := value.([]map[string]interface{}); ok {
		return r
	} else {
		array := Interfaces(value)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]interface{}, len(array))
		for k, v := range array {
			list[k] = Map(v, tags...)
		}
		return list
	}
}

// MapsDeep converts <i> to []map[string]interface{} recursively.
func MapsDeep(value interface{}, tags ...string) []map[string]interface{} {
	if value == nil {
		return nil
	}
	if r, ok := value.([]map[string]interface{}); ok {
		return r
	} else {
		array := Interfaces(value)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]interface{}, len(array))
		for k, v := range array {
			list[k] = MapDeep(v, tags...)
		}
		return list
	}
}
