package _type

import "sync/atomic"

type TBool struct {
	value int32
}

func NewBool(b ...bool) *TBool {
	inst := &TBool{}
	if len(b) > 0 {
		if b[0] {
			inst.value = 1
		} else {
			inst.value = 0
		}
	}
	return inst
}

func (b *TBool) Set(v bool) (old bool) {
	if v {
		old = atomic.SwapInt32(&b.value, 1) == 1
	} else {
		old = atomic.SwapInt32(&b.value, 0) == 1
	}
	return
}

func (b *TBool) Val() bool {
	return atomic.LoadInt32(&b.value) > 0
}

func (b *TBool) Clone() *TBool {
	return NewBool(b.Val())
}

func (b *TBool) CompareSwap(old, new bool) bool {
	var nOld, nNew int32
	if old {
		nOld = 1
	}
	if new {
		nNew = 1
	}
	return atomic.CompareAndSwapInt32(&b.value, nOld, nNew)
}
