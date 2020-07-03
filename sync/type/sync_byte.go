package _type

import "sync/atomic"

type TByte struct {
	value int32
}

func NewByte(value ...byte) *TByte {
	inst := &TByte{}
	if len(value) > 0 {
		inst.value = int32(value[0])
	}
	return inst
}

func (v *TByte) Clone() *TByte {
	return NewByte(v.Val())
}

func (v *TByte) Set(value byte) (old byte) {
	return byte(atomic.SwapInt32(&v.value, int32(value)))
}

func (v *TByte) Val() byte {
	return byte(atomic.LoadInt32(&v.value))
}

func (v *TByte) Add(delta byte) (new byte) {
	return byte(atomic.AddInt32(&v.value, int32(delta)))
}

func (v *TByte) CompareSwap(old, new byte) bool {
	return atomic.CompareAndSwapInt32(&v.value, int32(old), int32(new))
}
