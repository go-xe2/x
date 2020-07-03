package xtest

import (
	"testing"
)

func TestCase(t *testing.T) {
	Case(t, func() {
		Assert(1, 1)
		AssertNE(1, 0)
		AssertLt(float32(123.455), float32(123.456))
		AssertEq(float64(123.456), float64(123.456))
	})
}
