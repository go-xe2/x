package _type

import "sync/atomic"

type TInt struct {
	value int64
}

func NewInt(value ...int) *TInt {
	inst := &TInt{}
	if len(value) > 0 {
		inst.value = int64(value[0])
	}
	return inst
}

func (v *TInt) Clone() *TInt {
	return NewInt(v.Val())
}

func (v *TInt) Set(value int) (old int) {
	return int(atomic.SwapInt64(&v.value, int64(value)))
}

func (v *TInt) Val() int {
	return int(atomic.LoadInt64(&v.value))
}

func (v *TInt) Add(delta int) (new int) {
	return int(atomic.AddInt64(&v.value, int64(delta)))
}

func (v *TInt) CompareSwap(old, new int) bool {
	return atomic.CompareAndSwapInt64(&v.value, int64(old), int64(new))
}
