package xmap

import (
	"github.com/go-xe2/x/type/t"
	"time"
)

type TIntAnyMap map[int]interface{}

func NewIntAnyMap(mp ...map[int]interface{}) TIntAnyMap {
	var def map[int]interface{}
	if len(mp) > 0 {
		def = mp[0]
	} else {
		def = make(map[int]interface{})
	}
	return TIntAnyMap(def)
}

func (mp TIntAnyMap) Map() map[int]interface{} {
	return mp
}

func (mp TIntAnyMap) ToArray() [][]interface{} {
	result := make([][]interface{}, 0)
	for k, v := range mp {
		result = append(result, []interface{}{k, v})
	}
	return result
}

func (mp TIntAnyMap) GetString(key int, def ...string) string {
	s := ""
	if len(def) > 0 {
		s = def[0]
	}
	if s1, ok := mp[key]; ok {
		return t.String(s1, s)
	}
	return s
}

func (mp TIntAnyMap) GetInt(key int, def ...int) int {
	var n = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int(v, n)
	}
	return n
}

func (mp TIntAnyMap) GetInt8(key int, def ...int8) int8 {
	var n int8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int8(v, n)
	}
	return n
}

func (mp TIntAnyMap) GetInt16(key int, def ...int16) int16 {
	var n int16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int16(v, n)
	}
	return n
}

func (mp TIntAnyMap) GetInt32(key int, def ...int32) int32 {
	var n int32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int32(v, n)
	}
	return n
}

func (mp TIntAnyMap) GetInt64(key int, def ...int64) int64 {
	var n int64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int64(v, n)
	}
	return n
}

func (mp TIntAnyMap) GetUint(key int, def ...uint) uint {
	var n uint = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint(v, n)
	}
	return n
}

func (mp TIntAnyMap) GetUint8(key int, def ...uint8) uint8 {
	var n uint8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint8(v, n)
	}
	return n
}

func (mp TIntAnyMap) GetUint16(key int, def ...uint16) uint16 {
	var n uint16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint16(v, n)
	}
	return n
}

func (mp TIntAnyMap) GetUint32(key int, def ...uint32) uint32 {
	var n uint32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint32(v, n)
	}
	return n
}

func (mp TIntAnyMap) GetUint64(key int, def ...uint64) uint64 {
	var n uint64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint64(v, n)
	}
	return n
}

func (mp TIntAnyMap) GetFloat32(key int, def ...float32) float32 {
	var f float32 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Float32(v, f)
	}
	return f
}

func (mp TIntAnyMap) GetFloat64(key int, def ...float64) float64 {
	var f float64 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Float64(v, f)
	}
	return f
}

func (mp TIntAnyMap) GetBool(key int, def ...bool) bool {
	b := false
	if len(def) > 0 {
		b = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Bool(v)
	}
	return b
}

func (mp TIntAnyMap) GetInterface(key int, def ...interface{}) interface{} {
	var inf interface{} = nil
	if len(def) > 0 {
		inf = def[0]
	}
	if v, ok := mp[key].(interface{}); ok {
		return v
	}
	return inf
}

func (mp TIntAnyMap) GetMap(key int, def ...map[string]interface{}) interface{} {
	var mDef map[string]interface{} = nil
	if len(def) > 0 {
		mDef = def[0]
	}
	if v, ok := mp[key].(map[string]interface{}); ok {
		return v
	}
	return mDef
}

func (mp TIntAnyMap) GetArray(key int, def ...[]interface{}) []interface{} {
	var arr []interface{} = nil
	if len(def) > 0 {
		arr = def[0]
	}
	if v, ok := mp[key].([]interface{}); ok {
		return v
	}
	return arr
}

func (mp TIntAnyMap) GetTime(key int, def ...time.Time) time.Time {
	var tDef = time.Date(1900, 1, 1, 0, 0, 1, 0, nil)
	if len(def) > 0 {
		tDef = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Time(v, "y-m-d H:i:s")
	}
	return tDef
}
