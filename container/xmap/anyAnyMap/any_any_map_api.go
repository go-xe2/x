package anyAnyMap

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

func New(m ...map[interface{}]interface{}) xmap.TAnyAnyMap {
	var mDef map[interface{}]interface{}
	if len(m) > 0 {
		mDef = m[0]
	} else {
		mDef = make(map[interface{}]interface{})
	}
	return xmap.TAnyAnyMap(mDef)
}

func GetString(m map[interface{}]interface{}, key interface{}, def ...string) string {
	s := ""
	if len(def) > 0 {
		s = def[0]
	}
	if s1, ok := m[key]; ok {
		return t.String(s1, s)
	}
	return s
}

func GetInt(m map[interface{}]interface{}, key interface{}, def ...int) int {
	var n = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int(v, n)
	}
	return n
}

func GetInt8(m map[interface{}]interface{}, key interface{}, def ...int8) int8 {
	var n int8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int8(v, n)
	}
	return n
}

func GetInt16(m map[interface{}]interface{}, key interface{}, def ...int16) int16 {
	var n int16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int16(v, n)
	}
	return n
}

func GetInt32(m map[interface{}]interface{}, key interface{}, def ...int32) int32 {
	var n int32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int32(v, n)
	}
	return n
}

func GetInt64(m map[interface{}]interface{}, key interface{}, def ...int64) int64 {
	var n int64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int64(v, n)
	}
	return n
}

func GetUint(m map[interface{}]interface{}, key interface{}, def ...uint) uint {
	var n uint = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint(v, n)
	}
	return n
}

func GetUint8(m map[interface{}]interface{}, key interface{}, def ...uint8) uint8 {
	var n uint8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint8(v, n)
	}
	return n
}

func GetUint16(m map[interface{}]interface{}, key interface{}, def ...uint16) uint16 {
	var n uint16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint16(v, n)
	}
	return n
}

func GetUint32(m map[interface{}]interface{}, key interface{}, def ...uint32) uint32 {
	var n uint32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint32(v, n)
	}
	return n
}

func GetUint64(m map[interface{}]interface{}, key interface{}, def ...uint64) uint64 {
	var n uint64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint64(v, n)
	}
	return n
}

func GetFloat32(m map[interface{}]interface{}, key interface{}, def ...float32) float32 {
	var f float32 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Float32(v, f)
	}
	return f
}

func GetFloat64(m map[interface{}]interface{}, key interface{}, def ...float64) float64 {
	var f float64 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Float64(v, f)
	}
	return f
}

func GetBool(m map[interface{}]interface{}, key interface{}, def ...bool) bool {
	b := false
	if len(def) > 0 {
		b = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Bool(v)
	}
	return b
}

func GetTime(m map[interface{}]interface{}, key interface{}, def ...time.Time) time.Time {
	var tDef = time.Date(1900, 1, 1, 0, 0, 1, 0, nil)
	if len(def) > 0 {
		tDef = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Time(v, "y-m-d H:i:s")
	}
	return tDef
}

func GetInterface(m map[interface{}]interface{}, key interface{}, def ...interface{}) interface{} {
	var inf interface{} = nil
	if len(def) > 0 {
		inf = def[0]
	}
	if v, ok := m[key].(interface{}); ok {
		return v
	}
	return inf
}

func GetMap(m map[interface{}]interface{}, key interface{}, def ...map[string]interface{}) map[string]interface{} {
	var mDef map[string]interface{} = nil
	if len(def) > 0 {
		mDef = def[0]
	}
	if v, ok := m[key].(map[string]interface{}); ok {
		return v
	}
	return mDef
}

func GetArray(m map[interface{}]interface{}, key interface{}, def ...[]interface{}) []interface{} {
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
func ToKeyString(m map[interface{}]interface{}) string {
	if m == nil {
		return ""
	}
	var result bytes.Buffer
	var keys []string
	for key := range m {
		keys = append(keys, fmt.Sprintf("%v", key))
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := m[key]
		result.WriteString(key)
		result.WriteString("=")
		result.WriteString(fmt.Sprintf("%v", value))
		result.WriteString("&")
	}
	return result.String()
}

// 交换键值
func Flip(m *map[interface{}]interface{}) {
	result := make(map[interface{}]interface{})
	for k, v := range *m {
		result[v] = k
	}
	*m = result
}

// 拷贝map
// fields 可传字段列表字符串，字符串数组，或map,如果为map可设置默认值
func Clone(m map[interface{}]interface{}, fields ...interface{}) map[interface{}]interface{} {
	var selectFields = make(map[interface{}]interface{})
	if len(fields) > 0 {
		switch items := fields[0].(type) {
		case []string:
			for _, key := range items {
				if v, ok := m[key]; ok {
					selectFields[key] = v
				}
			}
			return selectFields
		case map[string]interface{}:
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
				if v, ok := m[key]; ok {
					selectFields[key] = v
				}
			}
			return selectFields
		default:
			fmt.Println("clone unknown fields params type:", reflect.TypeOf(fields[0]))
		}
		return nil
	}
	for k, v := range m {
		selectFields[k] = v
	}
	return selectFields
}
