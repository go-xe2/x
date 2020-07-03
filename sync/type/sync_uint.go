package _type

import "sync/atomic"

type TUint struct {
	value uint64
}

func NewUint(value ...uint) *TUint {
	inst := &TUint{}
	if len(value) > 0 {
		inst.value = uint64(value[0])
	}
	return inst
}

func (v *TUint) Clone() *TUint {
	return NewUint(v.Val())
}

func (v *TUint) Set(value uint) (old uint) {
	return uint(atomic.SwapUint64(&v.value, uint64(value)))
}

func (v *TUint) Val() uint {
	return uint(atomic.LoadUint64(&v.value))
}

func (v *TUint) Add(delta uint) (new uint) {
	return uint(atomic.AddUint64(&v.value, uint64(delta)))
}

func (v *TUint) CompareSwap(old, new uint) bool {
	return atomic.CompareAndSwapUint64(&v.value, uint64(old), uint64(new))
}
