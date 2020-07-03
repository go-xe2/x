package xanySlice

import (
	. "github.com/go-xe2/x/type/slice/comm"
	"github.com/go-xe2/x/type/t"
	"reflect"
	"sort"
)

func Contain(arr interface{}, item interface{}) bool {
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(arr)
		for i := 0; i < a.Len(); i++ {
			if reflect.DeepEqual(item, a.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

func ForEach(arr interface{}, fn ArrForEachFunc) {
	if fn == nil {
		return
	}
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(arr)
		for i := 0; i < a.Len(); i++ {
			fn(i, a.Index(i).Interface())
		}
	}
}

func Map(arr interface{}, fn ArrMapFunc) {
	if fn == nil {
		return
	}
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(arr)
		for i := 0; i < a.Len(); i++ {
			v := fn(i, a.Index(i).Interface())
			a.Index(i).Set(reflect.ValueOf(v))
		}
	}
}

func Find(arr interface{}, fn ArrSearchFunc) interface{} {
	if fn == nil {
		return nil
	}
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(arr)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			if fn(i, v) {
				return v
			}
		}
	}
	return nil
}

func FindIndex(arr interface{}, fn ArrSearchFunc) int {
	if fn == nil {
		return -1
	}
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(arr)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			if fn(i, v) {
				return i
			}
		}
	}
	return -1
}

func Sort(arr interface{}, comparer ...ArrSortCompareFunc) {
	if len(comparer) == 0 || comparer[0] == nil {
		return
	}
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		sort.Slice(arr, func(i, j int) bool {
			return comparer[0](i, j) < 0
		})
	}
}

func Join(arr interface{}, sep string) string {
	result := ""
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(arr)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			if result != "" {
				result += sep
			}
			result += t.String(v)
		}
		return result
	}
	return ""
}

func String(this interface{}) string {
	result := ""
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			if result != "" {
				result += ","
			}
			result += t.String(v)
		}
		return "[" + result + "]"
	}
	return ""
}

func AsInterface(this interface{}) []interface{} {
	var result = make([]interface{}, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, v)
		}
		return result
	}
	return nil
}

func AsString(this interface{}) []string {
	var result = make([]string, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.String(v))
		}
		return result
	}
	return nil
}

func AsInt(this interface{}) []int {
	var result = make([]int, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Int(v))
		}
		return result
	}
	return nil
}

func AsInt8(this interface{}) []int8 {
	var result = make([]int8, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Int8(v))
		}
		return result
	}
	return nil
}

func AsInt16(this interface{}) []int16 {
	var result = make([]int16, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Int16(v))
		}
		return result
	}
	return nil
}

func AsInt32(this interface{}) []int32 {
	var result = make([]int32, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Int32(v))
		}
		return result
	}
	return nil
}

func AsInt64(this interface{}) []int64 {
	var result = make([]int64, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Int64(v))
		}
		return result
	}
	return nil
}

func AsUint(this interface{}) []uint {
	var result = make([]uint, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Uint(v))
		}
		return result
	}
	return nil
}

func AsUint8(this interface{}) []uint8 {
	var result = make([]uint8, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Uint8(v))
		}
		return result
	}
	return nil
}

func AsUint16(this interface{}) []uint16 {
	var result = make([]uint16, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Uint16(v))
		}
		return result
	}
	return nil
}

func AsUint32(this interface{}) []uint32 {
	var result = make([]uint32, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Uint32(v))
		}
		return result
	}
	return nil
}

func AsUint64(this interface{}) []uint64 {
	var result = make([]uint64, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Uint64(v))
		}
		return result
	}
	return nil
}

func AsFloat(this interface{}) []float32 {
	var result = make([]float32, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Float32(v))
		}
		return result
	}
	return nil
}

func AsFloat64(this interface{}) []float64 {
	var result = make([]float64, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Float64(v))
		}
		return result
	}
	return nil
}

func AsBool(this interface{}) []bool {
	var result = make([]bool, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result = append(result, t.Bool(v))
		}
		return result
	}
	return nil
}

func AsMap(this interface{}) map[string]interface{} {
	var result = make(map[string]interface{}, 0)
	switch reflect.TypeOf(this).Kind() {
	case reflect.Slice:
		a := reflect.ValueOf(this)
		for i := 0; i < a.Len(); i++ {
			v := a.Index(i).Interface()
			result[t.String(i)] = v
		}
		return result
	}
	return nil
}
