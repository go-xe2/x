package xintSlie

import (
	. "github.com/go-xe2/x/type/slice/comm"
	"github.com/go-xe2/x/type/t"
	"sort"
	"strconv"
)

func Contain(arr []int, item int) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func ForEach(arr []int, fn IntArrForEachFunc) {
	if fn == nil {
		return
	}
	for k, v := range arr {
		fn(k, v)
	}
}

func Map(arr []int, fn IntArrMapFunc) {
	if fn == nil {
		return
	}
	for k, v := range arr {
		nv := fn(k, v)
		arr[k] = nv
	}
}

func Find(arr []int, fn IntArrSearchFunc) int {
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

func FindIndex(arr []int, fn IntArrSearchFunc) int {
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

func Sort(arr []int, comparer ...ArrSortCompareFunc) {
	var defComparer = func(aIndex int, bIndex int) int {
		n1 := arr[aIndex]
		n2 := arr[bIndex]
		return n1 - n2
	}
	if len(comparer) > 0 && comparer[0] != nil {
		defComparer = comparer[0]
	}
	sort.Slice(arr, func(i, j int) bool {
		return defComparer(i, j) < 0
	})
}

func Join(arr []int, sep string) string {
	var result = ""
	for _, v := range arr {
		if result != "" {
			result += sep
		}
		result += strconv.Itoa(v)
	}
	return result
}

func String(this []int) string {
	var result = ""
	for _, v := range this {
		if result != "" {
			result += ","
		}
		result += t.String(v)
	}
	return "[" + result + "]"
}

func AsInterface(this []int) []interface{} {
	var result = make([]interface{}, 0)
	for _, v := range this {
		result = append(result, v)
	}
	return result
}
