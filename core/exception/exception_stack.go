package exception

import "runtime"

type IStack interface {
	Stack() string
}

type stack []uintptr

const (
	mMAX_STACK_DEPTH = 32
)

func callers() stack {
	var pcs [mMAX_STACK_DEPTH]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[0:n]
}
