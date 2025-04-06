package lib

import (
	"testing"
)

func BenchmarkFibonaci(b *testing.B) {
	for b.Loop() {
		Fibonaci(10)
	}
}

func BenchmarkFibonaciIterative(b *testing.B) {
	for b.Loop() {
		FibonaciIterative(10)
	}
}
