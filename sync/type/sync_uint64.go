package _type

import "sync/atomic"

type TUint64 struct {
	value uint64
}

func NewUint64(value ...uint64) *TUint64 {
	inst := &TUint64{}
	if len(value) > 0 {
		inst.value = value[0]
	}
	return inst
}

func (v *TUint64) Clone() *TUint64 {
	return NewUint64(v.Val())
}

func (v *TUint64) Set(value uint64) (old uint64) {
	return atomic.SwapUint64(&v.value, value)
}

func (v *TUint64) Val() uint64 {
	return atomic.LoadUint64(&v.value)
}

func (v *TUint64) Add(delta uint64) (new uint64) {
	return atomic.AddUint64(&v.value, delta)
}

func (v *TUint64) CompareSwap(old, new uint64) bool {
	return atomic.CompareAndSwapUint64(&v.value, old, new)
}
