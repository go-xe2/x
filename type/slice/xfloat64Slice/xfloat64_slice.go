package xfloat64Slice

import (
	. "github.com/go-xe2/x/type/slice/comm"
	"github.com/go-xe2/x/type/t"
	"sort"
)

func Contain(arr []float64, item float64) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func ForEach(arr []float64, fn Float64ArrForEachFunc) {
	if fn == nil {
		return
	}
	for k, v := range arr {
		fn(k, v)
	}
}

func Map(arr []float64, fn Float64ArrMapFunc) {
	if fn == nil {
		return
	}
	for k, v := range arr {
		nv := fn(k, v)
		arr[k] = nv
	}
}

func Find(arr []float64, fn Float64ArrSearchFunc) float64 {
	if fn == nil {
		return -1
	}
	for k, v := range arr {
		if fn(k, v) {
			return v
		}
	}
	return -1
}

func FindIndex(arr []float64, fn Float64ArrSearchFunc) int {
	if fn == nil {
		return -1
	}
	for k, v := range arr {
		if fn(k, v) {
			return k
		}
	}
	return -1
}

func Sort(arr []float64, comparer ...ArrSortCompareFunc) {
	var defComparer = func(aIndex int, bIndex int) int {
		n1 := arr[aIndex]
		n2 := arr[bIndex]
		return int(n1 - n2)
	}
	if len(comparer) > 0 && comparer[0] != nil {
		defComparer = comparer[0]
	}
	sort.Slice(arr, func(i, j int) bool {
		return defComparer(i, j) < 0
	})
}

func Join(arr []float64, sep string) string {
	var result = ""
	for _, v := range arr {
		if result != "" {
			result += sep
		}
		result += t.String(v)
	}
	return result
}

func String(this []float64) string {
	var result = ""
	for _, v := range this {
		if result != "" {
			result += ","
		}
		result += t.String(v)
	}
	return "[" + result + "]"
}

func AsInterface(this []float64) []interface{} {
	var result = make([]interface{}, 0)
	for _, v := range this {
		result = append(result, v)
	}
	return result
}
