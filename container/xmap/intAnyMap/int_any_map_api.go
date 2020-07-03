package intAnyMap

import (
	"bytes"
	"fmt"
	"github.com/go-xe2/x/container/xmap"
	"github.com/go-xe2/x/type/t"
	"reflect"
	"sort"
	"strings"
	"time"
)

func New(m ...map[int]interface{}) xmap.TIntAnyMap {
	var mDef map[int]interface{}
	if len(m) > 0 {
		mDef = m[0]
	} else {
		mDef = make(map[int]interface{})
	}
	return xmap.TIntAnyMap(mDef)
}

func GetString(m map[int]interface{}, key int, def ...string) string {
	s := ""
	if len(def) > 0 {
		s = def[0]
	}
	if s1, ok := m[key]; ok {
		return t.String(s1, s)
	}
	return s
}

func GetInt(m map[int]interface{}, key int, def ...int) int {
	var n = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int(v, n)
	}
	return n
}

func GetInt8(m map[int]interface{}, key int, def ...int8) int8 {
	var n int8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int8(v, n)
	}
	return n
}

func GetInt16(m map[int]interface{}, key int, def ...int16) int16 {
	var n int16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int16(v, n)
	}
	return n
}

func GetInt32(m map[int]interface{}, key int, def ...int32) int32 {
	var n int32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int32(v, n)
	}
	return n
}

func GetInt64(m map[int]interface{}, key int, def ...int64) int64 {
	var n int64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int64(v, n)
	}
	return n
}

func GetUint(m map[int]interface{}, key int, def ...uint) uint {
	var n uint = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint(v, n)
	}
	return n
}

func GetUint8(m map[int]interface{}, key int, def ...uint8) uint8 {
	var n uint8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint8(v, n)
	}
	return n
}

func GetUint16(m map[int]interface{}, key int, def ...uint16) uint16 {
	var n uint16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint16(v, n)
	}
	return n
}

func GetUint32(m map[int]interface{}, key int, def ...uint32) uint32 {
	var n uint32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint32(v, n)
	}
	return n
}

func GetUint64(m map[int]interface{}, key int, def ...uint64) uint64 {
	var n uint64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint64(v, n)
	}
	return n
}

func GetFloat32(m map[int]interface{}, key int, def ...float32) float32 {
	var f float32 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Float32(v, f)
	}
	return f
}

func GetFloat64(m map[int]interface{}, key int, def ...float64) float64 {
	var f float64 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Float64(v, f)
	}
	return f
}

func GetBool(m map[int]interface{}, key int, def ...bool) bool {
	b := false
	if len(def) > 0 {
		b = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Bool(v)
	}
	return b
}

func GetTime(m map[int]interface{}, key int, def ...time.Time) time.Time {
	var tDef = time.Date(1900, 1, 1, 0, 0, 1, 0, nil)
	if len(def) > 0 {
		tDef = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Time(v, "y-m-d H:i:s")
	}
	return tDef
}

func GetInterface(m map[int]interface{}, key int, def ...interface{}) interface{} {
	var inf interface{} = nil
	if len(def) > 0 {
		inf = def[0]
	}
	if v, ok := m[key].(interface{}); ok {
		return v
	}
	return inf
}

func GetMap(m map[int]interface{}, key int, def ...map[string]interface{}) map[string]interface{} {
	var mDef map[string]interface{} = nil
	if len(def) > 0 {
		mDef = def[0]
	}
	if v, ok := m[key].(map[string]interface{}); ok {
		return v
	}
	return mDef
}

func GetArray(m map[int]interface{}, key int, def ...[]interface{}) []interface{} {
	var arr []interface{} = nil
	if len(def) > 0 {
		arr = def[0]
	}
	if v, ok := m[key].([]interface{}); ok {
		return v
	}
	return arr
}

// map转换成可作为键名的字符串
func ToKeyString(m map[int]interface{}) string {
	if m == nil {
		return ""
	}
	var result bytes.Buffer
	var keys []int
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		value := m[key]
		result.WriteString(t.String(key))
		result.WriteString("=")
		result.WriteString(fmt.Sprintf("%v", value))
		result.WriteString("&")
	}
	return result.String()
}

// 拷贝map
// fields 可传字段列表字符串，字符串数组，或map,如果为map可设置默认值
func Clone(m map[int]interface{}, fields ...interface{}) map[interface{}]interface{} {
	var selectFields = make(map[interface{}]interface{})
	if len(fields) > 0 {
		switch items := fields[0].(type) {
		case []int:
			for _, key := range items {
				if v, ok := m[key]; ok {
					selectFields[key] = v
				}
			}
			return selectFields
		case map[int]interface{}:
			for k, v := range items {
				if v1, ok := m[k]; ok {
					selectFields[k] = v1
				} else {
					selectFields[k] = v
				}
			}
			return selectFields
		case string:
			strItems := strings.Split(items, ",")
			for _, key := range strItems {
				key = strings.Trim(key, " ")

				if v, ok := m[t.Int(key)]; ok {
					selectFields[key] = v
				}
			}
			return selectFields
		default:
			fmt.Println("xmap.clone unknown fields params type:", reflect.TypeOf(fields[0]))
		}
		return nil
	}
	for k, v := range m {
		selectFields[k] = v
	}
	return selectFields
}
