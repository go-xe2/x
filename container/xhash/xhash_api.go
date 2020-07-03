package xhash

import (
	"bytes"
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
	"reflect"
	"sort"
	"strings"
	"time"
)

func GetPathString(m map[string]interface{}, path string, def ...string) string {
	v := GetPathInterface(m, path, nil)
	return t.String(v, def...)
}

func GetPathInt(m map[string]interface{}, path string, def ...int) int {
	v := GetPathInterface(m, path, nil)
	return t.Int(v, def...)
}

func GetPathInt8(m map[string]interface{}, path string, def ...int8) int8 {
	v := GetPathInterface(m, path, nil)
	return t.Int8(v, def...)
}

func GetPathInt16(m map[string]interface{}, path string, def ...int16) int16 {
	v := GetPathInterface(m, path, nil)
	return t.Int16(v, def...)
}

func GetPathInt32(m map[string]interface{}, path string, def ...int32) int32 {
	v := GetPathInterface(m, path, nil)
	return t.Int32(v, def...)
}

func GetPathInt64(m map[string]interface{}, path string, def ...int64) int64 {
	v := GetPathInterface(m, path, nil)
	return t.Int64(v, def...)
}

func GetPathUint(m map[string]interface{}, path string, def ...uint) uint {
	v := GetPathInterface(m, path, nil)
	return t.Uint(v, def...)
}

func GetPathUint8(m map[string]interface{}, path string, def ...uint8) uint8 {
	v := GetPathInterface(m, path, nil)
	return t.Uint8(v, def...)
}

func GetPathUint16(m map[string]interface{}, path string, def ...uint16) uint16 {
	v := GetPathInterface(m, path, nil)
	return t.Uint16(v, def...)
}

func GetPathUint32(m map[string]interface{}, path string, def ...uint32) uint32 {
	v := GetPathInterface(m, path, nil)
	return t.Uint32(v, def...)
}

func GetPathUint64(m map[string]interface{}, path string, def ...uint64) uint64 {
	v := GetPathInterface(m, path, nil)
	return t.Uint64(v, def...)
}

func GetPathBool(m map[string]interface{}, path string, def ...bool) bool {
	v := GetPathInterface(m, path, nil)
	return t.Bool(v, def...)
}

func GetPathFloat(m map[string]interface{}, path string, def ...float32) float32 {
	v := GetPathInterface(m, path, nil)
	return t.Float32(v, def...)
}

func GetPathFloat64(m map[string]interface{}, path string, def ...float64) float64 {
	v := GetPathInterface(m, path, nil)
	return t.Float64(v, def...)
}

func GetPathTime(m map[string]interface{}, path string, def ...time.Time) time.Time {
	v := GetPathInterface(m, path, nil)
	var tDef = time.Date(1900, 1, 1, 0, 0, 1, 0, nil)
	if len(def) > 0 {
		tDef = def[0]
	}
	if v != nil {
		return t.Time(v, "y-m-d H:i:s")
	}
	return tDef
}

func GetPathMap(m map[string]interface{}, path string, def ...map[string]interface{}) map[string]interface{} {
	v := GetPathInterface(m, path, nil)
	var mDef map[string]interface{} = nil
	if len(def) > 0 {
		mDef = def[0]
	}
	if v == nil {
		return mDef
	}
	if v1, ok := v.(map[string]interface{}); ok {
		return v1
	}
	return mDef
}

func GetPathArray(m map[string]interface{}, path string, def ...[]interface{}) []interface{} {
	v := GetPathInterface(m, path, nil)
	var mDef []interface{} = nil
	if len(def) > 0 {
		mDef = def[0]
	}
	if v == nil {
		return mDef
	}
	if v, ok := v.([]interface{}); ok {
		return v
	}
	return mDef
}

func GetString(m map[string]interface{}, key string, def ...string) string {
	s := ""
	if len(def) > 0 {
		s = def[0]
	}
	if s1, ok := m[key]; ok {
		return t.String(s1, s)
	}
	return s
}

func GetInt(m map[string]interface{}, key string, def ...int) int {
	var n = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int(v, n)
	}
	return n
}

func GetInt8(m map[string]interface{}, key string, def ...int8) int8 {
	var n int8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int8(v, n)
	}
	return n
}

func GetInt16(m map[string]interface{}, key string, def ...int16) int16 {
	var n int16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int16(v, n)
	}
	return n
}

func GetInt32(m map[string]interface{}, key string, def ...int32) int32 {
	var n int32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int32(v, n)
	}
	return n
}

func GetInt64(m map[string]interface{}, key string, def ...int64) int64 {
	var n int64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Int64(v, n)
	}
	return n
}

func GetUint(m map[string]interface{}, key string, def ...uint) uint {
	var n uint = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint(v, n)
	}
	return n
}

func GetUint8(m map[string]interface{}, key string, def ...uint8) uint8 {
	var n uint8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint8(v, n)
	}
	return n
}

func GetUint16(m map[string]interface{}, key string, def ...uint16) uint16 {
	var n uint16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint16(v, n)
	}
	return n
}

func GetUint32(m map[string]interface{}, key string, def ...uint32) uint32 {
	var n uint32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint32(v, n)
	}
	return n
}

func GetUint64(m map[string]interface{}, key string, def ...uint64) uint64 {
	var n uint64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Uint64(v, n)
	}
	return n
}

func GetFloat32(m map[string]interface{}, key string, def ...float32) float32 {
	var f float32 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Float32(v, f)
	}
	return f
}

func GetFloat64(m map[string]interface{}, key string, def ...float64) float64 {
	var f float64 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Float64(v, f)
	}
	return f
}

func GetBool(m map[string]interface{}, key string, def ...bool) bool {
	b := false
	if len(def) > 0 {
		b = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Bool(v)
	}
	return b
}

func GetTime(m map[string]interface{}, key string, def ...time.Time) time.Time {
	var tDef = time.Date(1900, 1, 1, 0, 0, 1, 0, nil)
	if len(def) > 0 {
		tDef = def[0]
	}
	if v, ok := m[key]; ok {
		return t.Time(v, "y-m-d H:i:s")
	}
	return tDef
}

func GetInterface(m map[string]interface{}, key string, def ...interface{}) interface{} {
	var inf interface{} = nil
	if len(def) > 0 {
		inf = def[0]
	}
	if v, ok := m[key].(interface{}); ok {
		return v
	}
	return inf
}

func GetMap(m map[string]interface{}, key string, def ...map[string]interface{}) map[string]interface{} {
	var mDef map[string]interface{} = nil
	if len(def) > 0 {
		mDef = def[0]
	}
	if v, ok := m[key].(map[string]interface{}); ok {
		return v
	}
	return mDef
}

func GetArray(m map[string]interface{}, key string, def ...[]interface{}) []interface{} {
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
func ToKeyString(m map[string]interface{}) string {
	if m == nil {
		return ""
	}
	var result bytes.Buffer
	var keys []string
	for key := range m {
		keys = append(keys, key)
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

// 拷贝map
// fields 可传字段列表字符串，字符串数组，或map,如果为map可设置默认值
func Clone(m map[string]interface{}, fields ...interface{}) map[string]interface{} {
	var selectFields = make(map[string]interface{})
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
			fmt.Println("xmap.clone unknown fields params type:", reflect.TypeOf(fields[0]))
		}
		return nil
	}
	for k, v := range m {
		selectFields[k] = v
	}
	return selectFields
}

func GetPathInterface(m map[string]interface{}, path string, def ...interface{}) interface{} {
	if path == "" || m == nil {
		return nil
	}
	nodes := xstring.Split(path, "/")
	p := m
	nLen := len(nodes)
	var ok = false
	for i := 0; i < nLen-1; i++ {
		key := nodes[i]
		p, ok = p[key].(map[string]interface{})
		if !ok {
			p = nil
			break
		}
	}
	if p != nil {
		if v, ok := p[nodes[nLen-1]]; ok {
			return v
		}
		return def
	}
	return def
}
