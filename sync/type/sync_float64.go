package _type

import (
	"math"
	"sync/atomic"
	"unsafe"
)

type TFloat64 struct {
	value uint64
}

func NewFloat64(value ...float64) *TFloat64 {
	inst := &TFloat64{}
	if len(value) > 0 {
		inst.value = math.Float64bits(value[0])
	}
	return &TFloat64{}
}

func (v *TFloat64) Clone() *TFloat64 {
	return NewFloat64(v.Val())
}

func (v *TFloat64) Set(value float64) (old float64) {
	return math.Float64frombits(atomic.SwapUint64(&v.value, math.Float64bits(value)))
}

func (v *TFloat64) Val() float64 {
	return math.Float64frombits(atomic.LoadUint64(&v.value))
}

func (v *TFloat64) Add(delta float64) (new float64) {
	for {
		old := math.Float64frombits(v.value)
		new = old + delta
		if atomic.CompareAndSwapUint64(
			(*uint64)(unsafe.Pointer(&v.value)),
			math.Float64bits(old),
			math.Float64bits(new),
		) {
			break
		}
	}
	return
}

func (v *TFloat64) CompareSwap(old, new float64) bool {
	return atomic.CompareAndSwapUint64(&v.value, uint64(old), uint64(new))
}
