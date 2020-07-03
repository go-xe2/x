package xjson

import (
	"fmt"
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xtime"
	"github.com/go-xe2/x/utils/xutil"
	"time"
)

func (j *TJson) Value() interface{} {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return *(j.p)
}

func (j *TJson) IsNil() bool {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.p == nil || *(j.p) == nil
}

func (j *TJson) Get(pattern string, def ...interface{}) interface{} {
	j.mu.RLock()
	defer j.mu.RUnlock()

	var result *interface{}
	if j.vc {
		result = j.getPointerByPattern(pattern)
	} else {
		result = j.getPointerByPatternWithoutViolenceCheck(pattern)
	}
	if result != nil {
		return *result
	}
	if len(def) > 0 {
		return def[0]
	}
	return nil
}

func (j *TJson) GetVar(pattern string, def ...interface{}) *_type.TVar {
	return _type.NewVar(j.Get(pattern, def...), true)
}

func (j *TJson) GetMap(pattern string, def ...interface{}) map[string]interface{} {
	result := j.Get(pattern, def...)
	if result != nil {
		return t.Map(result)
	}
	return nil
}

func (j *TJson) GetJson(pattern string, def ...interface{}) *TJson {
	return New(j.Get(pattern, def...), true)
}

func (j *TJson) GetJsons(pattern string, def ...interface{}) []*TJson {
	array := j.GetArray(pattern, def...)
	if len(array) > 0 {
		jsonSlice := make([]*TJson, len(array))
		for i := 0; i < len(array); i++ {
			jsonSlice[i] = New(array[i], true)
		}
		return jsonSlice
	}
	return nil
}

func (j *TJson) GetJsonMap(pattern string, def ...interface{}) map[string]*TJson {
	m := j.GetMap(pattern, def...)
	if len(m) > 0 {
		jsonMap := make(map[string]*TJson, len(m))
		for k, v := range m {
			jsonMap[k] = New(v, true)
		}
		return jsonMap
	}
	return nil
}

func (j *TJson) GetArray(pattern string, def ...interface{}) []interface{} {
	return t.Interfaces(j.Get(pattern, def...))
}

func (j *TJson) GetString(pattern string, def ...interface{}) string {
	return t.String(j.Get(pattern, def...))
}

func (j *TJson) GetBytes(pattern string, def ...interface{}) []byte {
	return t.Bytes(j.Get(pattern, def...))
}

func (j *TJson) GetBool(pattern string, def ...interface{}) bool {
	return t.Bool(j.Get(pattern, def...))
}

func (j *TJson) GetInt(pattern string, def ...interface{}) int {
	return t.Int(j.Get(pattern, def...))
}

func (j *TJson) GetInt8(pattern string, def ...interface{}) int8 {
	return t.Int8(j.Get(pattern, def...))
}

func (j *TJson) GetInt16(pattern string, def ...interface{}) int16 {
	return t.Int16(j.Get(pattern, def...))
}

func (j *TJson) GetInt32(pattern string, def ...interface{}) int32 {
	return t.Int32(j.Get(pattern, def...))
}

func (j *TJson) GetInt64(pattern string, def ...interface{}) int64 {
	return t.Int64(j.Get(pattern, def...))
}

func (j *TJson) GetUint(pattern string, def ...interface{}) uint {
	return t.Uint(j.Get(pattern, def...))
}

func (j *TJson) GetUint8(pattern string, def ...interface{}) uint8 {
	return t.Uint8(j.Get(pattern, def...))
}

func (j *TJson) GetUint16(pattern string, def ...interface{}) uint16 {
	return t.Uint16(j.Get(pattern, def...))
}

func (j *TJson) GetUint32(pattern string, def ...interface{}) uint32 {
	return t.Uint32(j.Get(pattern, def...))
}

func (j *TJson) GetUint64(pattern string, def ...interface{}) uint64 {
	return t.Uint64(j.Get(pattern, def...))
}

func (j *TJson) GetFloat32(pattern string, def ...interface{}) float32 {
	return t.Float32(j.Get(pattern, def...))
}

func (j *TJson) GetFloat64(pattern string, def ...interface{}) float64 {
	return t.Float64(j.Get(pattern, def...))
}

func (j *TJson) GetFloats(pattern string, def ...interface{}) []float64 {
	return t.Floats(j.Get(pattern, def...))
}

func (j *TJson) GetInts(pattern string, def ...interface{}) []int {
	return t.Ints(j.Get(pattern, def...))
}

func (j *TJson) GetStrings(pattern string, def ...interface{}) []string {
	return t.Strings(j.Get(pattern, def...))
}

func (j *TJson) GetInterfaces(pattern string, def ...interface{}) []interface{} {
	return t.Interfaces(j.Get(pattern, def...))
}

func (j *TJson) GetTime(pattern string, format ...string) time.Time {
	return t.Time(j.Get(pattern), format...)
}

func (j *TJson) GetDuration(pattern string, def ...interface{}) time.Duration {
	return t.Duration(j.Get(pattern, def...))
}

func (j *TJson) GetGTime(pattern string, format ...string) *xtime.Time {
	return t.XTime(j.Get(pattern), format...)
}

func (j *TJson) Set(pattern string, value interface{}) error {
	return j.setValue(pattern, value, false)
}

func (j *TJson) Remove(pattern string) error {
	return j.setValue(pattern, nil, true)
}

func (j *TJson) Contains(pattern string) bool {
	return j.Get(pattern) != nil
}

func (j *TJson) Len(pattern string) int {
	p := j.getPointerByPattern(pattern)
	if p != nil {
		switch (*p).(type) {
		case map[string]interface{}:
			return len((*p).(map[string]interface{}))
		case []interface{}:
			return len((*p).([]interface{}))
		default:
			return -1
		}
	}
	return -1
}

func (j *TJson) Append(pattern string, value interface{}) error {
	p := j.getPointerByPattern(pattern)
	if p == nil {
		return j.Set(fmt.Sprintf("%s.0", pattern), value)
	}
	switch (*p).(type) {
	case []interface{}:
		return j.Set(fmt.Sprintf("%s.%d", pattern, len((*p).([]interface{}))), value)
	}
	return fmt.Errorf("invalid variable type of %s", pattern)
}

func (j *TJson) GetToVar(pattern string, pointer interface{}) error {
	r := j.Get(pattern)
	if r != nil {
		if t, err := Encode(r); err == nil {
			return DecodeTo(t, pointer)
		} else {
			return err
		}
	} else {
		pointer = nil
	}
	return nil
}

func (j *TJson) GetStruct(pattern string, pointer interface{}, mapping ...map[string]string) error {
	return t.Struct(j.Get(pattern), pointer, mapping...)
}

func (j *TJson) GetStructDeep(pattern string, pointer interface{}, mapping ...map[string]string) error {
	return t.StructDeep(j.Get(pattern), pointer, mapping...)
}

func (j *TJson) GetStructs(pattern string, pointer interface{}, mapping ...map[string]string) error {
	return t.Structs(j.Get(pattern), pointer, mapping...)
}

func (j *TJson) GetStructsDeep(pattern string, pointer interface{}, mapping ...map[string]string) error {
	return t.StructsDeep(j.Get(pattern), pointer, mapping...)
}

func (j *TJson) GetToStruct(pattern string, pointer interface{}, mapping ...map[string]string) error {
	return j.GetStruct(pattern, pointer, mapping...)
}

func (j *TJson) ToMap() map[string]interface{} {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return t.Map(*(j.p))
}

func (j *TJson) ToArray() []interface{} {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return t.Interfaces(*(j.p))
}

func (j *TJson) ToStruct(pointer interface{}) error {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return t.Struct(*(j.p), pointer)
}

func (j *TJson) ToStructDeep(pointer interface{}) error {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return t.StructDeep(*(j.p), pointer)
}

func (j *TJson) ToStructs(pointer interface{}) error {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return t.Structs(*(j.p), pointer)
}

func (j *TJson) ToStructsDeep(pointer interface{}) error {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return t.StructsDeep(*(j.p), pointer)
}

func (j *TJson) Dump() {
	j.mu.RLock()
	defer j.mu.RUnlock()
	xutil.Dump(*j.p)
}

func (j *TJson) Export() string {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return xutil.Export(*j.p)
}
