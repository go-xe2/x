package xparser

import (
	"github.com/go-xe2/x/encoding/xjson"
	_type "github.com/go-xe2/x/sync/type"
	"github.com/go-xe2/x/type/xtime"
	"time"
)

type TParser struct {
	json *xjson.TJson
}

func (p *TParser) Value() interface{} {
	return p.json.Value()
}

func (p *TParser) Get(pattern string, def ...interface{}) interface{} {
	return p.json.Get(pattern, def...)
}

func (p *TParser) GetVar(pattern string, def ...interface{}) *_type.TVar {
	return p.json.GetVar(pattern, def...)
}

func (p *TParser) GetMap(pattern string, def ...interface{}) map[string]interface{} {
	return p.json.GetMap(pattern, def...)
}

func (p *TParser) GetArray(pattern string, def ...interface{}) []interface{} {
	return p.json.GetArray(pattern, def...)
}

func (p *TParser) GetString(pattern string, def ...interface{}) string {
	return p.json.GetString(pattern, def...)
}

func (p *TParser) GetBool(pattern string, def ...interface{}) bool {
	return p.json.GetBool(pattern, def...)
}

func (p *TParser) GetInt(pattern string, def ...interface{}) int {
	return p.json.GetInt(pattern, def...)
}

func (p *TParser) GetInt8(pattern string, def ...interface{}) int8 {
	return p.json.GetInt8(pattern, def...)
}

func (p *TParser) GetInt16(pattern string, def ...interface{}) int16 {
	return p.json.GetInt16(pattern, def...)
}

func (p *TParser) GetInt32(pattern string, def ...interface{}) int32 {
	return p.json.GetInt32(pattern, def...)
}

func (p *TParser) GetInt64(pattern string, def ...interface{}) int64 {
	return p.json.GetInt64(pattern, def...)
}

func (p *TParser) GetInts(pattern string, def ...interface{}) []int {
	return p.json.GetInts(pattern, def...)
}

func (p *TParser) GetUint(pattern string, def ...interface{}) uint {
	return p.json.GetUint(pattern, def...)
}

func (p *TParser) GetUint8(pattern string, def ...interface{}) uint8 {
	return p.json.GetUint8(pattern, def...)
}

func (p *TParser) GetUint16(pattern string, def ...interface{}) uint16 {
	return p.json.GetUint16(pattern, def...)
}

func (p *TParser) GetUint32(pattern string, def ...interface{}) uint32 {
	return p.json.GetUint32(pattern, def...)
}

func (p *TParser) GetUint64(pattern string, def ...interface{}) uint64 {
	return p.json.GetUint64(pattern, def...)
}

func (p *TParser) GetFloat32(pattern string, def ...interface{}) float32 {
	return p.json.GetFloat32(pattern, def...)
}

func (p *TParser) GetFloat64(pattern string, def ...interface{}) float64 {
	return p.json.GetFloat64(pattern, def...)
}

func (p *TParser) GetFloats(pattern string, def ...interface{}) []float64 {
	return p.json.GetFloats(pattern, def...)
}

func (p *TParser) GetStrings(pattern string, def ...interface{}) []string {
	return p.json.GetStrings(pattern, def...)
}

func (p *TParser) GetInterfaces(pattern string, def ...interface{}) []interface{} {
	return p.json.GetInterfaces(pattern, def...)
}

func (p *TParser) GetTime(pattern string, format ...string) time.Time {
	return p.json.GetTime(pattern, format...)
}

func (p *TParser) GetDuration(pattern string, def ...interface{}) time.Duration {
	return p.json.GetDuration(pattern, def...)
}

func (p *TParser) GetGTime(pattern string, format ...string) *xtime.Time {
	return p.json.GetGTime(pattern, format...)
}

func (p *TParser) GetToVar(pattern string, pointer interface{}) error {
	return p.json.GetToVar(pattern, pointer)
}

func (p *TParser) GetStruct(pattern string, pointer interface{}, mapping ...map[string]string) error {
	return p.json.GetStruct(pattern, pointer, mapping...)
}

func (p *TParser) GetStructDeep(pattern string, pointer interface{}, mapping ...map[string]string) error {
	return p.json.GetStructDeep(pattern, pointer, mapping...)
}

func (p *TParser) GetStructs(pattern string, pointer interface{}, mapping ...map[string]string) error {
	return p.json.GetStructs(pattern, pointer, mapping...)
}

func (p *TParser) GetStructsDeep(pattern string, pointer interface{}, mapping ...map[string]string) error {
	return p.json.GetStructsDeep(pattern, pointer, mapping...)
}

func (p *TParser) GetToStruct(pattern string, pointer interface{}, mapping ...map[string]string) error {
	return p.json.GetStruct(pattern, pointer, mapping...)
}

func (p *TParser) Set(pattern string, value interface{}) error {
	return p.json.Set(pattern, value)
}

func (p *TParser) Len(pattern string) int {
	return p.json.Len(pattern)
}

func (p *TParser) Append(pattern string, value interface{}) error {
	return p.json.Append(pattern, value)
}

func (p *TParser) Remove(pattern string) error {
	return p.json.Remove(pattern)
}

func (p *TParser) ToMap() map[string]interface{} {
	return p.json.ToMap()
}

func (p *TParser) ToArray() []interface{} {
	return p.json.ToArray()
}

func (p *TParser) ToStruct(pointer interface{}) error {
	return p.json.ToStruct(pointer)
}

func (p *TParser) ToStructDeep(pointer interface{}) error {
	return p.json.ToStructDeep(pointer)
}

func (p *TParser) ToStructs(pointer interface{}) error {
	return p.json.ToStructs(pointer)
}

func (p *TParser) ToStructsDeep(pointer interface{}) error {
	return p.json.ToStructsDeep(pointer)
}

func (p *TParser) Dump() {
	p.json.Dump()
}

func (p *TParser) Export() string {
	return p.json.Export()
}
