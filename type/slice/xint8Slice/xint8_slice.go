package int8Array

import (
	. "github.com/go-xe2/x/type/slice/comm"
	"github.com/go-xe2/x/type/t"
	"sort"
	"strconv"
)

func Contain(arr []int8, item int8) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func ForEach(arr []int8, fn Int8ArrForEachFunc) {
	if fn == nil {
		return
	}
	for k, v := range arr {
		fn(k, v)
	}
}

func Map(arr []int8, fn Int8ArrMapFunc) {
	if fn == nil {
		return
	}
	for k, v := range arr {
		nv := fn(k, v)
		arr[k] = nv
	}
}

func Find(arr []int8, fn Int8ArrSearchFunc) int8 {
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

func FindIndex(arr []int8, fn Int8ArrSearchFunc) int {
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

func Sort(arr []int8, comparer ...ArrSortCompareFunc) {
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

func Join(arr []int8, sep string) string {
	var result = ""
	for _, v := range arr {
		if result != "" {
			result += sep
		}
		result += strconv.Itoa(int(v))
	}
	return result
}

func String(this []int8) string {
	var result = ""
	for _, v := range this {
		if result != "" {
			result += ","
		}
		result += t.String(v)
	}
	return "[" + result + "]"
}

func AsInterface(this []int8) []interface{} {
	var result = make([]interface{}, 0)
	for _, v := range this {
		result = append(result, v)
	}
	return result
}
