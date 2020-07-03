package _type

import "sync/atomic"

type TInt32 struct {
	value int32
}

func NewInt32(value ...int32) *TInt32 {
	inst := &TInt32{}
	if len(value) > 0 {
		inst.value = value[0]
	}
	return inst
}

func (v *TInt32) Clone() *TInt32 {
	return NewInt32(v.Val())
}

func (v *TInt32) Set(value int32) (old int32) {
	return atomic.SwapInt32(&v.value, value)
}

func (v *TInt32) Val() int32 {
	return atomic.LoadInt32(&v.value)
}

func (v *TInt32) Add(delta int32) (new int32) {
	return atomic.AddInt32(&v.value, delta)
}

func (v *TInt32) CompareSwap(old, new int32) bool {
	return atomic.CompareAndSwapInt32(&v.value, old, new)
}
