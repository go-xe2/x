package xclass

import (
	"fmt"
	"testing"
)

func TestClassToClassVT(t *testing.T) {
	rt := classToClassVT(TDemo2Class.Type())
	fmt.Print(rt)
}
