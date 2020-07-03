package _type

import (
	"fmt"
	"sync/atomic"
)

type TInterface struct {
	value atomic.Value
}

func NewInterface(value ...interface{}) *TInterface {
	t := &TInterface{}
	if len(value) > 0 && value[0] != nil {
		t.value.Store(value[0])
	}
	return t
}

func (v *TInterface) Clone() *TInterface {
	return NewInterface(v.Val())
}

func (v *TInterface) Set(value interface{}) (old interface{}) {
	old = v.Val()
	v.value.Store(value)
	return
}

func (v *TInterface) Val() interface{} {
	return v.value.Load()
}

func (v *TInterface) String() string {
	val := v.Val()
	return fmt.Sprintf("%v", val)
}
