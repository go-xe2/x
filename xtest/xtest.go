// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gtest provides convenient test utilities for unit testing.

package xtest

import (
	"fmt"
	"github.com/go-xe2/x/core/debug"
	"github.com/go-xe2/x/type/t"
	"os"
	"reflect"
	"testing"
)

const (
	mPATH_FILTER_KEY = "github.com/go-xe2/x/xtest/xtest"
)

// 测试案例，检查panic并输出
func Case(t *testing.T, f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n%s", err, debug.StackWithFilter(mPATH_FILTER_KEY))
			t.Fail()
		}
	}()
	f()
}

// 是否相等测试检查
func Assert(value, expect interface{}) {
	rvExpect := reflect.ValueOf(expect)
	if isNil(value) {
		value = nil
	}
	if rvExpect.Kind() == reflect.Map {
		if err := compareMap(value, expect); err != nil {
			panic(err)
		}
		return
	}
	strValue := t.String(value)
	strExpect := t.String(expect)
	if strValue != strExpect {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v == %v`, strValue, strExpect))
	}
}

// 是否相等测试检查
func AssertEq(value, expect interface{}) {
	// Value assert.
	rvExpect := reflect.ValueOf(expect)
	if isNil(value) {
		value = nil
	}
	if rvExpect.Kind() == reflect.Map {
		if err := compareMap(value, expect); err != nil {
			panic(err)
		}
		return
	}
	strValue := t.String(value)
	strExpect := t.String(expect)
	if strValue != strExpect {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v == %v`, strValue, strExpect))
	}
	// Type assert.
	t1 := reflect.TypeOf(value)
	t2 := reflect.TypeOf(expect)
	if t1 != t2 {
		panic(fmt.Sprintf(`[ASSERT] EXPECT TYPE %v[%v] == %v[%v]`, strValue, t1, strExpect, t2))
	}
}

// 不相等测试检查
func AssertNE(value, expect interface{}) {
	rvExpect := reflect.ValueOf(expect)
	if isNil(value) {
		value = nil
	}
	if rvExpect.Kind() == reflect.Map {
		if err := compareMap(value, expect); err == nil {
			panic(fmt.Sprintf(`[ASSERT] EXPECT %v != %v`, value, expect))
		}
		return
	}
	strValue := t.String(value)
	strExpect := t.String(expect)
	if strValue == strExpect {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v != %v`, strValue, strExpect))
	}
}

// 大于测试检查
func AssertGt(value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = t.String(value) > t.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = t.Int(value) > t.Int(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = t.Uint(value) > t.Uint(expect)

	case reflect.Float32, reflect.Float64:
		passed = t.Float64(value) > t.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v > %v`, value, expect))
	}
}

// 大于等于测试检查, AssertGE的别名
func AssertGte(value, expect interface{}) {
	AssertGE(value, expect)
}

// 大于等于测试检查
func AssertGE(value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = t.String(value) >= t.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = t.Int(value) >= t.Int(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = t.Uint(value) >= t.Uint(expect)

	case reflect.Float32, reflect.Float64:
		passed = t.Float64(value) >= t.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v >= %v`, value, expect))
	}
}

// 小于测试检查
func AssertLt(value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = t.String(value) < t.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = t.Int(value) < t.Int(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = t.Uint(value) < t.Uint(expect)

	case reflect.Float32, reflect.Float64:
		passed = t.Float64(value) < t.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v < %v`, value, expect))
	}
}

// 小于等于测试检查，AssertLE别名
func AssertLte(value, expect interface{}) {
	AssertLE(value, expect)
}

// 小于等于测试检查
func AssertLE(value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = t.String(value) <= t.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = t.Int(value) <= t.Int(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = t.Uint(value) <= t.Uint(expect)

	case reflect.Float32, reflect.Float64:
		passed = t.Float64(value) <= t.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v <= %v`, value, expect))
	}
}

// 数组中包含某项检查
func AssertIN(value, expect interface{}) {
	passed := true
	switch reflect.ValueOf(expect).Kind() {
	case reflect.Slice, reflect.Array:
		expectSlice := t.Interfaces(expect)
		for _, v1 := range t.Interfaces(value) {
			result := false
			for _, v2 := range expectSlice {
				if v1 == v2 {
					result = true
					break
				}
			}
			if !result {
				passed = false
				break
			}
		}
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v IN %v`, value, expect))
	}
}

// 数组中不包含某项检查
func AssertNI(value, expect interface{}) {
	passed := true
	switch reflect.ValueOf(expect).Kind() {
	case reflect.Slice, reflect.Array:
		for _, v1 := range t.Interfaces(value) {
			result := true
			for _, v2 := range t.Interfaces(expect) {
				if v1 == v2 {
					result = false
					break
				}
			}
			if !result {
				passed = false
				break
			}
		}
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] EXPECT %v NOT IN %v`, value, expect))
	}
}

// panic错误信息, 使用Case捕获
func Error(message ...interface{}) {
	panic(fmt.Sprintf("[ERROR] %s", fmt.Sprint(message...)))
}

// 输出错误信息并退出进程
func Fatal(message ...interface{}) {
	fmt.Fprintf(os.Stderr, "[FATAL] %s\n%s", fmt.Sprint(message...), debug.StackWithFilter(mPATH_FILTER_KEY))
	os.Exit(1)
}

// 比较两个map是否相等，相等时返回nil,否则返回错误
func compareMap(value, expect interface{}) error {
	rvValue := reflect.ValueOf(value)
	rvExpect := reflect.ValueOf(expect)
	if isNil(value) {
		value = nil
	}
	if rvExpect.Kind() == reflect.Map {
		if rvValue.Kind() == reflect.Map {
			if rvExpect.Len() == rvValue.Len() {
				// Turn two interface maps to the same type for comparison.
				// Direct use of rvValue.MapIndex(key).Interface() will panic
				// when the key types are inconsistent.
				mValue := make(map[string]string)
				mExpect := make(map[string]string)
				ksValue := rvValue.MapKeys()
				ksExpect := rvExpect.MapKeys()
				for _, key := range ksValue {
					mValue[t.String(key.Interface())] = t.String(rvValue.MapIndex(key).Interface())
				}
				for _, key := range ksExpect {
					mExpect[t.String(key.Interface())] = t.String(rvExpect.MapIndex(key).Interface())
				}
				for k, v := range mExpect {
					if v != mValue[k] {
						return fmt.Errorf(`[ASSERT] EXPECT VALUE map["%v"]:%v == map["%v"]:%v`+
							"\nGIVEN : %v\nEXPECT: %v", k, mValue[k], k, v, mValue, mExpect)
					}
				}
			} else {
				return fmt.Errorf(`[ASSERT] EXPECT MAP LENGTH %d == %d`, rvValue.Len(), rvExpect.Len())
			}
		} else {
			return fmt.Errorf(`[ASSERT] EXPECT VALUE TO BE A MAP`)
		}
	}
	return nil
}

// 检查值是否为nil
func isNil(value interface{}) bool {
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.Ptr, reflect.Func:
		return rv.IsNil()
	default:
		return value == nil
	}
}
