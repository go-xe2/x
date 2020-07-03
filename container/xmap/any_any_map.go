package xmap

import (
	"github.com/go-xe2/x/type/t"
	"time"
)

type TAnyAnyMap map[interface{}]interface{}

func NewAnyAnyMap(mp ...map[interface{}]interface{}) TAnyAnyMap {
	var def map[interface{}]interface{}
	if len(mp) > 0 {
		def = mp[0]
	} else {
		def = make(map[interface{}]interface{})
	}
	return TAnyAnyMap(def)
}

func (mp TAnyAnyMap) Map() map[interface{}]interface{} {
	return mp
}

func (mp TAnyAnyMap) ToArray() [][]interface{} {
	result := make([][]interface{}, 0)
	for k, v := range mp {
		result = append(result, []interface{}{k, v})
	}
	return result
}

func (mp TAnyAnyMap) GetString(key interface{}, def ...string) string {
	s := ""
	if len(def) > 0 {
		s = def[0]
	}
	if s1, ok := mp[key]; ok {
		return t.String(s1, s)
	}
	return s
}

func (mp TAnyAnyMap) GetInt(key interface{}, def ...int) int {
	var n = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int(v, n)
	}
	return n
}

func (mp TAnyAnyMap) GetInt8(key interface{}, def ...int8) int8 {
	var n int8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int8(v, n)
	}
	return n
}

func (mp TAnyAnyMap) GetInt16(key interface{}, def ...int16) int16 {
	var n int16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int16(v, n)
	}
	return n
}

func (mp TAnyAnyMap) GetInt32(key interface{}, def ...int32) int32 {
	var n int32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int32(v, n)
	}
	return n
}

func (mp TAnyAnyMap) GetInt64(key interface{}, def ...int64) int64 {
	var n int64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int64(v, n)
	}
	return n
}

func (mp TAnyAnyMap) GetUint(key interface{}, def ...uint) uint {
	var n uint = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint(v, n)
	}
	return n
}

func (mp TAnyAnyMap) GetUint8(key interface{}, def ...uint8) uint8 {
	var n uint8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint8(v, n)
	}
	return n
}

func (mp TAnyAnyMap) GetUint16(key interface{}, def ...uint16) uint16 {
	var n uint16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint16(v, n)
	}
	return n
}

func (mp TAnyAnyMap) GetUint32(key interface{}, def ...uint32) uint32 {
	var n uint32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint32(v, n)
	}
	return n
}

func (mp TAnyAnyMap) GetUint64(key interface{}, def ...uint64) uint64 {
	var n uint64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint64(v, n)
	}
	return n
}

func (mp TAnyAnyMap) GetFloat32(key interface{}, def ...float32) float32 {
	var f float32 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Float32(v, f)
	}
	return f
}

func (mp TAnyAnyMap) GetFloat64(key interface{}, def ...float64) float64 {
	var f float64 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Float64(v, f)
	}
	return f
}

func (mp TAnyAnyMap) GetBool(key interface{}, def ...bool) bool {
	b := false
	if len(def) > 0 {
		b = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Bool(v)
	}
	return b
}

func (mp TAnyAnyMap) GetInterface(key interface{}, def ...interface{}) interface{} {
	var inf interface{} = nil
	if len(def) > 0 {
		inf = def[0]
	}
	if v, ok := mp[key].(interface{}); ok {
		return v
	}
	return inf
}

func (mp TAnyAnyMap) GetMap(key interface{}, def ...map[string]interface{}) interface{} {
	var mDef map[string]interface{} = nil
	if len(def) > 0 {
		mDef = def[0]
	}
	if v, ok := mp[key].(map[string]interface{}); ok {
		return v
	}
	return mDef
}

func (mp TAnyAnyMap) GetArray(key interface{}, def ...[]interface{}) []interface{} {
	var arr []interface{} = nil
	if len(def) > 0 {
		arr = def[0]
	}
	if v, ok := mp[key].([]interface{}); ok {
		return v
	}
	return arr
}

func (mp TAnyAnyMap) GetTime(key interface{}, def ...time.Time) time.Time {
	var tDef = time.Date(1900, 1, 1, 0, 0, 1, 0, nil)
	if len(def) > 0 {
		tDef = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Time(v, "y-m-d H:i:s")
	}
	return tDef
}

// 交换键值
func (mp TAnyAnyMap) Flip() TAnyAnyMap {
	result := make(map[interface{}]interface{})
	for k, v := range mp {
		result[v] = k
	}
	return TAnyAnyMap(result)
}
