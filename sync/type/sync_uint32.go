package _type

import "sync/atomic"

type TUint32 struct {
	value uint32
}

func NewUint32(value ...uint32) *TUint32 {
	inst := &TUint32{}
	if len(value) > 0 {
		inst.value = value[0]
	}
	return inst
}

func (v *TUint32) Clone() *TUint32 {
	return NewUint32(v.Val())
}

func (v *TUint32) Set(value uint32) (old uint32) {
	return atomic.SwapUint32(&v.value, value)
}

func (v *TUint32) Val() uint32 {
	return atomic.LoadUint32(&v.value)
}

func (v *TUint32) Add(delta uint32) (new uint32) {
	return atomic.AddUint32(&v.value, delta)
}

func (v *TUint32) CompareSwap(old, new uint32) bool {
	return atomic.CompareAndSwapUint32(&v.value, old, new)
}
