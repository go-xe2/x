package xmap

import (
	"github.com/go-xe2/x/type/t"
)

type TStrIntMap map[string]int

func NewStrIntMap(mp ...map[string]int) TStrIntMap {
	var def map[string]int
	if len(mp) > 0 {
		def = mp[0]
	} else {
		def = make(map[string]int)
	}
	return TStrIntMap(def)
}

func (mp TStrIntMap) Map() map[string]int {
	return mp
}

func (mp TStrIntMap) ToArray() [][]interface{} {
	result := make([][]interface{}, 0)
	for k, v := range mp {
		result = append(result, []interface{}{k, v})
	}
	return result
}

func (mp TStrIntMap) GetString(key string, def ...int) int {
	s := 0
	if len(def) > 0 {
		s = def[0]
	}
	if s1, ok := mp[key]; ok {
		return s1
	}
	return s
}

func (mp TStrIntMap) GetInt(key string, def ...int) int {
	var n = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int(v, n)
	}
	return n
}

func (mp TStrIntMap) GetInt8(key string, def ...int8) int8 {
	var n int8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int8(v, n)
	}
	return n
}

func (mp TStrIntMap) GetInt16(key string, def ...int16) int16 {
	var n int16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int16(v, n)
	}
	return n
}

func (mp TStrIntMap) GetInt32(key string, def ...int32) int32 {
	var n int32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int32(v, n)
	}
	return n
}

func (mp TStrIntMap) GetInt64(key string, def ...int64) int64 {
	var n int64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Int64(v, n)
	}
	return n
}

func (mp TStrIntMap) GetUint(key string, def ...uint) uint {
	var n uint = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint(v, n)
	}
	return n
}

func (mp TStrIntMap) GetUint8(key string, def ...uint8) uint8 {
	var n uint8 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint8(v, n)
	}
	return n
}

func (mp TStrIntMap) GetUint16(key string, def ...uint16) uint16 {
	var n uint16 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint16(v, n)
	}
	return n
}

func (mp TStrIntMap) GetUint32(key string, def ...uint32) uint32 {
	var n uint32 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint32(v, n)
	}
	return n
}

func (mp TStrIntMap) GetUint64(key string, def ...uint64) uint64 {
	var n uint64 = 0
	if len(def) > 0 {
		n = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Uint64(v, n)
	}
	return n
}

func (mp TStrIntMap) GetFloat32(key string, def ...float32) float32 {
	var f float32 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Float32(v, f)
	}
	return f
}

func (mp TStrIntMap) GetFloat64(key string, def ...float64) float64 {
	var f float64 = 0
	if len(def) > 0 {
		f = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Float64(v, f)
	}
	return f
}

func (mp TStrIntMap) GetBool(key string, def ...bool) bool {
	b := false
	if len(def) > 0 {
		b = def[0]
	}
	if v, ok := mp[key]; ok {
		return t.Bool(v)
	}
	return b
}
