package _type

import "sync/atomic"

type TInt64 struct {
	value int64
}

func NewInt64(value ...int64) *TInt64 {
	inst := &TInt64{}
	if len(value) > 0 {
		inst.value = value[0]
	}
	return inst
}

func (v *TInt64) Clone() *TInt64 {
	return NewInt64(v.Val())
}

func (v *TInt64) Set(value int64) (old int64) {
	return atomic.SwapInt64(&v.value, value)
}

func (v *TInt64) Val() int64 {
	return atomic.LoadInt64(&v.value)
}

func (v *TInt64) Add(delta int64) (new int64) {
	return atomic.AddInt64(&v.value, delta)
}

func (v *TInt64) CompareSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&v.value, old, new)
}
