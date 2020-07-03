package xfloat

import (
	"github.com/go-xe2/x/type/t"
)

type XFloat float64

func New(v ...float64) XFloat {
	if len(v) > 0 {
		return XFloat(v[0])
	}
	return XFloat(0)
}

func New32(v ...float32) XFloat {
	if len(v) > 0 {
		return XFloat(v[0])
	}
	return XFloat(0)
}

func (f XFloat) String() string {
	return t.String(f)
}
