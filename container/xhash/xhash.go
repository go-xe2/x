package xhash

import (
	"time"
)

type THashMap map[string]interface{}

func New(mp ...map[string]interface{}) THashMap {
	var def map[string]interface{}
	if len(mp) > 0 {
		def = mp[0]
	} else {
		def = make(map[string]interface{})
	}
	return THashMap(def)
}

func (mp THashMap) Map() map[string]interface{} {
	return mp
}

func (mp THashMap) ToArray() [][]interface{} {
	result := make([][]interface{}, 0)
	for k, v := range mp {
		result = append(result, []interface{}{k, v})
	}
	return result
}

func (mp THashMap) GetString(key string, def ...string) string {
	return GetString(mp, key, def...)
}

func (mp THashMap) GetInt(key string, def ...int) int {
	return GetInt(mp, key, def...)
}

func (mp THashMap) GetInt8(key string, def ...int8) int8 {
	return GetInt8(mp, key, def...)
}

func (mp THashMap) GetInt16(key string, def ...int16) int16 {
	return GetInt16(mp, key, def...)
}

func (mp THashMap) GetInt32(key string, def ...int32) int32 {
	return GetInt32(mp, key, def...)
}

func (mp THashMap) GetInt64(key string, def ...int64) int64 {
	return GetInt64(mp, key, def...)
}

func (mp THashMap) GetUint(key string, def ...uint) uint {
	return GetUint(mp, key, def...)
}

func (mp THashMap) GetUint8(key string, def ...uint8) uint8 {
	return GetUint8(mp, key, def...)
}

func (mp THashMap) GetUint16(key string, def ...uint16) uint16 {
	return GetUint16(mp, key, def...)
}

func (mp THashMap) GetUint32(key string, def ...uint32) uint32 {
	return GetUint32(mp, key, def...)
}

func (mp THashMap) GetUint64(key string, def ...uint64) uint64 {
	return GetUint64(mp, key, def...)
}

func (mp THashMap) GetFloat32(key string, def ...float32) float32 {
	return GetFloat32(mp, key, def...)
}

func (mp THashMap) GetFloat64(key string, def ...float64) float64 {
	return GetFloat64(mp, key, def...)
}

func (mp THashMap) GetBool(key string, def ...bool) bool {
	return GetBool(mp, key, def...)
}

func (mp THashMap) GetInterface(key string, def ...interface{}) interface{} {
	return GetInterface(mp, key, def...)
}

func (mp THashMap) GetMap(key string, def ...map[string]interface{}) interface{} {
	return GetMap(mp, key, def...)
}

func (mp THashMap) GetArray(key string, def ...[]interface{}) []interface{} {
	return GetArray(mp, key, def...)
}

func (mp THashMap) GetTime(key string, def ...time.Time) time.Time {
	return GetTime(mp, key, def...)
}

func (mp THashMap) GetPathInterface(path string, def ...interface{}) interface{} {
	return GetPathInterface(mp, path, def...)
}

func (mp THashMap) GetPathString(path string, def ...string) string {
	return GetPathString(mp, path, def...)
}

func (mp THashMap) GetPathInt(path string, def ...int) int {
	return GetPathInt(mp, path, def...)
}

func (mp THashMap) GetPathInt8(path string, def ...int8) int8 {
	return GetPathInt8(mp, path, def...)
}

func (mp THashMap) GetPathInt16(path string, def ...int16) int16 {
	return GetPathInt16(mp, path, def...)
}

func (mp THashMap) GetPathInt32(path string, def ...int32) int32 {
	return GetPathInt32(mp, path, def...)
}

func (mp THashMap) GetPathInt64(path string, def ...int64) int64 {
	return GetPathInt64(mp, path, def...)
}

func (mp THashMap) GetPathUint(path string, def ...uint) uint {
	return GetPathUint(mp, path, def...)
}

func (mp THashMap) GetPathUint8(path string, def ...uint8) uint8 {
	return GetPathUint8(mp, path, def...)
}

func (mp THashMap) GetPathUint16(path string, def ...uint16) uint16 {
	return GetPathUint16(mp, path, def...)
}

func (mp THashMap) GetPathUint32(path string, def ...uint32) uint32 {
	return GetPathUint32(mp, path, def...)
}

func (mp THashMap) GetPathUint64(path string, def ...uint64) uint64 {
	return GetPathUint64(mp, path, def...)
}

func (mp THashMap) GetPathFloat(path string, def ...float32) float32 {
	return GetPathFloat(mp, path, def...)
}

func (mp THashMap) GetPathFloat64(path string, def ...float64) float64 {
	return GetPathFloat64(mp, path, def...)
}

func (mp THashMap) GetPathBool(path string, def ...bool) bool {
	return GetPathBool(mp, path, def...)
}

func (mp THashMap) GetPathMap(path string, def ...map[string]interface{}) map[string]interface{} {
	return GetPathMap(mp, path, def...)
}

func (mp THashMap) GetPathArray(path string, def ...[]interface{}) []interface{} {
	return GetPathArray(mp, path, def...)
}
