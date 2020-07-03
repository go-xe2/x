package xrunPool

import (
	"testing"
)

func increment() {
	for i := 0; i < 1000000; i++ {
	}
}

func BenchmarkGrpool_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(increment)
	}
}

func BenchmarkGoroutine_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go increment()
	}
}
