package _type

import (
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xtime"
	"time"
)

type TVar struct {
	safe  bool
	value interface{}
}

func NewVar(value interface{}, unsafe ...bool) *TVar {
	v := &TVar{}
	if len(unsafe) == 0 || !unsafe[0] {
		v.safe = true
		v.value = NewInterface(value)
	} else {
		v.value = value
	}
	return v
}

func (v *TVar) IsSafe() bool {
	return v.safe
}

func (v *TVar) Val() interface{} {
	if v.safe {
		return v.value.(*TInterface).Val()
	} else {
		return v.value
	}
}

func (v *TVar) Interface() interface{} {
	return v.Val()
}

func (v *TVar) Set(val interface{}) (old interface{}) {
	if v.safe {
		old = v.value.(*TInterface).Set(val)
	} else {
		old = v.value
		v.value = val
	}
	return
}

func (v *TVar) Clone() *TVar {
	return NewVar(v.value, v.safe)
}

func (v *TVar) Assign(src *TVar) {
	v.safe = src.safe
	v.value = src.value
}

func (v *TVar) IsNil() bool {
	return v.Val() == nil
}

func (v *TVar) Bytes() []byte {
	return t.Bytes(v.Val())
}

func (v *TVar) String() string {
	return t.String(v.Val())
}

func (v *TVar) Bool() bool {
	return t.Bool(v.Val())
}

func (v *TVar) Int() int {
	return t.Int(v.Val())
}

func (v *TVar) Int8() int8 {
	return t.Int8(v.Val())
}

func (v *TVar) Int16() int16 {
	return t.Int16(v.Val())
}

func (v *TVar) Int32() int32 {
	return t.Int32(v.Val())
}

func (v *TVar) Int64() int64 {
	return t.Int64(v.Val())
}

func (v *TVar) Uint() uint {
	return t.Uint(v.Val())
}

func (v *TVar) Uint8() uint8 {
	return t.Uint8(v.Val())
}

func (v *TVar) Uint16() uint16 {
	return t.Uint16(v.Val())
}

func (v *TVar) Uint32() uint32 {
	return t.Uint32(v.Val())
}

func (v *TVar) Uint64() uint64 {
	return t.Uint64(v.Val())
}

func (v *TVar) Float32() float32 {
	return t.Float32(v.Val())
}

func (v *TVar) Float64() float64 {
	return t.Float64(v.Val())
}

func (v *TVar) Ints() []int {
	return t.Ints(v.Val())
}

func (v *TVar) Floats() []float64 {
	return t.Floats(v.Val())
}

func (v *TVar) Strings() []string {
	return t.Strings(v.Val())
}

func (v *TVar) Interfaces() []interface{} {
	return t.Interfaces(v.Val())
}

func (v *TVar) Time(format ...string) time.Time {
	return t.Time(v.Val(), format...)
}

func (v *TVar) Duration() time.Duration {
	return t.Duration(v.Val())
}

func (v *TVar) XTime(format ...string) *xtime.Time {
	return t.XTime(v.Val(), format...)
}

func (v *TVar) Map(tags ...string) map[string]interface{} {
	return t.Map(v.Val(), tags...)
}

func (v *TVar) MapDeep(tags ...string) map[string]interface{} {
	return t.MapDeep(v.Val(), tags...)
}

func (v *TVar) Struct(pointer interface{}, mapping ...map[string]string) error {
	return t.Struct(v.Val(), pointer, mapping...)
}

func (v *TVar) StructDeep(pointer interface{}, mapping ...map[string]string) error {
	return t.StructDeep(v.Val(), pointer, mapping...)
}

func (v *TVar) Structs(pointer interface{}, mapping ...map[string]string) (err error) {
	return t.Structs(v.Val(), pointer, mapping...)
}

func (v *TVar) StructsDeep(pointer interface{}, mapping ...map[string]string) (err error) {
	return t.StructsDeep(v.Val(), pointer, mapping...)
}
