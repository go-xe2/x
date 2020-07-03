package xanySlice

import (
	"errors"
	. "github.com/go-xe2/x/type/slice/comm"
	"github.com/go-xe2/x/type/t"
	"sort"
)

func (arr TAnySlice) Get(index int) interface{} {
	return arr[index]
}

func (arr TAnySlice) Contain(item interface{}) bool {
	size := arr.Size()
	for i := 0; i < size; i++ {
		if arr.Get(i) == item {
			return true
		}
	}
	return false
}

func (arr TAnySlice) ForEach(fn ArrForEachFunc) {
	if fn == nil {
		return
	}
	for i := 0; i < arr.Size(); i++ {
		fn(i, arr[i])
	}
}

func (arr TAnySlice) Map(fn ArrMapFunc) {
	if fn == nil {
		return
	}
	for i := 0; i < arr.Size(); i++ {
		v := fn(i, arr[i])
		arr[i] = v
	}
}

func (arr TAnySlice) Find(fn ArrSearchFunc) interface{} {
	if fn == nil {
		return nil
	}
	for i := 0; i < arr.Size(); i++ {
		v := arr[i]
		if fn(i, v) {
			return v
		}
	}
	return nil
}

func (arr TAnySlice) FindIndex(fn ArrSearchFunc) int {
	if fn == nil {
		return -1
	}
	for i := 0; i < arr.Size(); i++ {
		if fn(i, arr[i]) {
			return i
		}
	}
	return -1
}

func (arr TAnySlice) Sort(comparer ArrSortCompareFunc) {
	if comparer == nil {
		return
	}
	sort.Slice(arr, func(i, j int) bool {
		return comparer(i, j) < 0
	})
}

func (arr TAnySlice) Size() int {
	return len(arr)
}

func (arr *TAnySlice) Append(item ...interface{}) int {
	*arr = TAnySlice(append(*arr, item...))
	return len(*arr)
}

func (arr *TAnySlice) Prepend(item ...interface{}) int {
	old := *arr
	*arr = append(make([]interface{}, 0), item...)
	*arr = append(*arr, old...)
	return len(*arr)
}

func (arr *TAnySlice) Insert(index int, item interface{}) error {
	if index < 0 || index >= arr.Size() {
		return errors.New("数组下标越界")
	}
	old := *arr
	*arr = append(old[:index], item)
	*arr = append(*arr, old[index:]...)
	return nil
}

func (arr *TAnySlice) Clear() {
	*arr = make([]interface{}, 0)
}

func (arr *TAnySlice) Delete(index int) error {
	if index < 0 || index >= arr.Size() {
		return errors.New("数组下标越界")
	}
	old := *arr
	*arr = append(make([]interface{}, 0), old[:index-1]...)
	*arr = append(*arr, old[index+1:]...)
	return nil
}

func (arr *TAnySlice) Concat(arrs ...[]interface{}) int {
	for _, nar := range arrs {
		*arr = append(*arr, nar...)
	}
	return arr.Size()
}

func (arr TAnySlice) Json(sep string) string {
	result := ""
	for i := 0; i < arr.Size(); i++ {
		if result != "" {
			result += sep
		}
		result += t.String(arr[i])
	}
	return result
}

func (arr TAnySlice) String() string {
	return "[" + arr.Json(",") + "]"
}

func (arr TAnySlice) AsString() []string {
	result := make([]string, 0)
	for _, v := range arr {
		result = append(result, t.String(v))
	}
	return result
}

func (arr TAnySlice) AsInt() []int {
	result := make([]int, 0)
	for _, v := range arr {
		result = append(result, t.Int(v))
	}
	return result
}

func (arr TAnySlice) AsInt8() []int8 {
	result := make([]int8, 0)
	for _, v := range arr {
		result = append(result, t.Int8(v))
	}
	return result
}

func (arr TAnySlice) AsInt16() []int16 {
	result := make([]int16, 0)
	for _, v := range arr {
		result = append(result, t.Int16(v))
	}
	return result
}

func (arr TAnySlice) AsInt32() []int32 {
	result := make([]int32, 0)
	for _, v := range arr {
		result = append(result, t.Int32(v))
	}
	return result
}

func (arr TAnySlice) AsInt64() []int64 {
	result := make([]int64, 0)
	for _, v := range arr {
		result = append(result, t.Int64(v))
	}
	return result
}

func (arr TAnySlice) AsUint() []uint {
	result := make([]uint, 0)
	for _, v := range arr {
		result = append(result, t.Uint(v))
	}
	return result
}

func (arr TAnySlice) AsUint8() []uint8 {
	result := make([]uint8, 0)
	for _, v := range arr {
		result = append(result, t.Uint8(v))
	}
	return result
}

func (arr TAnySlice) AsUint16() []uint16 {
	result := make([]uint16, 0)
	for _, v := range arr {
		result = append(result, t.Uint16(v))
	}
	return result
}

func (arr TAnySlice) AsUint32() []uint32 {
	result := make([]uint32, 0)
	for _, v := range arr {
		result = append(result, t.Uint32(v))
	}
	return result
}

func (arr TAnySlice) AsUint64() []uint64 {
	result := make([]uint64, 0)
	for _, v := range arr {
		result = append(result, t.Uint64(v))
	}
	return result
}

func (arr TAnySlice) AsFloat() []float32 {
	result := make([]float32, 0)
	for _, v := range arr {
		result = append(result, t.Float(v))
	}
	return result
}

func (arr TAnySlice) AsFloat64() []float64 {
	result := make([]float64, 0)
	for _, v := range arr {
		result = append(result, t.Float64(v))
	}
	return result
}

func (arr TAnySlice) AsBool() []bool {
	result := make([]bool, 0)
	for _, v := range arr {
		result = append(result, t.Bool(v))
	}
	return result
}

func (arr TAnySlice) AsMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range arr {
		result[t.String(k)] = v
	}
	return result
}
