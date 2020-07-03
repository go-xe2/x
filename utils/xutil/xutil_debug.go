package xutil

import "github.com/go-xe2/x/core/debug"

func PrintStack(skip ...int) {
	number := 1
	if len(skip) > 0 {
		number = skip[0] + 1
	}
	debug.PrintStack(number)
}

func Stack(skip ...int) string {
	number := 1
	if len(skip) > 0 {
		number = skip[0] + 1
	}
	return debug.Stack(number)
}
