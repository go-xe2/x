package xstrSlice

import (
	. "github.com/go-xe2/x/type/slice/comm"
	"github.com/go-xe2/x/type/t"
	"sort"
)

func Contain(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func ForEach(arr []string, fn StrArrForEachFunc) {
	if fn == nil {
		return
	}
	for k, v := range arr {
		fn(k, v)
	}
}

func Map(arr []string, fn StrArrMapFunc) {
	if fn == nil {
		return
	}
	for k, v := range arr {
		nv := fn(k, v)
		arr[k] = nv
	}
}

func Find(arr []string, fn StrArrSearchFunc) string {
	if fn == nil {
		return ""
	}
	for k, v := range arr {
		if fn(k, v) {
			return v
		}
	}
	return ""
}

func FindIndex(arr []string, fn StrArrSearchFunc) int {
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

func Sort(arr []string, comparer ...ArrSortCompareFunc) {
	var defComparer = func(aIndex int, bIndex int) int {
		s1 := arr[aIndex]
		s2 := arr[bIndex]
		if s1 < s2 {
			return -1
		} else if s1 > s2 {
			return 1
		} else {
			return 0
		}
	}
	if len(comparer) > 0 && comparer[0] != nil {
		defComparer = comparer[0]
	}
	sort.Slice(arr, func(i, j int) bool {
		return defComparer(i, j) < 0
	})
}

func Join(arr []string, sep string) string {
	var result = ""
	for _, s := range arr {
		if result != "" {
			result += sep
		}
		result += s
	}
	return result
}

func String(this []string) string {
	var result = ""
	for _, v := range this {
		if result != "" {
			result += ","
		}
		result += t.String(v)
	}
	return "[" + result + "]"
}

func AsInterface(this []string) []interface{} {
	var result = make([]interface{}, 0)
	for _, v := range this {
		result = append(result, v)
	}
	return result
}
