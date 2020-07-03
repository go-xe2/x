package _type

import "sync/atomic"

type TBytes struct {
	value atomic.Value
}

func NewBytes(value ...[]byte) *TBytes {
	inst := &TBytes{}
	if len(value) > 0 {
		inst.value.Store(value[0])
	}
	return inst
}

func (v *TBytes) Clone() *TBytes {
	return NewBytes(v.Val())
}

func (v *TBytes) Set(value []byte) (old []byte) {
	old = v.Val()
	v.value.Store(value)
	return
}

func (v *TBytes) Val() []byte {
	if s := v.value.Load(); s != nil {
		return s.([]byte)
	}
	return nil
}
