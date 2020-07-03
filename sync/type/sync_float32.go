package _type

import (
	"math"
	"sync/atomic"
	"unsafe"
)

type TFloat32 struct {
	value uint32
}

func NewFloat32(value ...float32) *TFloat32 {
	inst := &TFloat32{}
	if len(value) > 0 {
		inst.value = math.Float32bits(value[0])
	}
	return inst
}

func (v *TFloat32) Clone() *TFloat32 {
	return NewFloat32(v.Val())
}

func (v *TFloat32) Set(value float32) (old float32) {
	return math.Float32frombits(atomic.SwapUint32(&v.value, math.Float32bits(value)))
}

func (v *TFloat32) Val() float32 {
	return math.Float32frombits(atomic.LoadUint32(&v.value))
}

func (v *TFloat32) Add(delta float32) (new float32) {
	for {
		old := math.Float32frombits(v.value)
		new = old + delta
		if atomic.CompareAndSwapUint32(
			(*uint32)(unsafe.Pointer(&v.value)),
			math.Float32bits(old),
			math.Float32bits(new),
		) {
			break
		}
	}
	return
}

func (v *TFloat32) CompareSwap(old, new float32) bool {
	return atomic.CompareAndSwapUint32(&v.value, uint32(old), uint32(new))
}
