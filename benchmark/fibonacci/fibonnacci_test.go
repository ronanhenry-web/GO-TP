package fibonacci_test

import (
	"testing"

	"github.com/Chroq/benchmark/fibonacci"
)

/* func BenchmarkFibonnacciReccursive100(b *testing.B) {
	for b.Loop() {
		fibonacci.FibonacciRecursive(100)
	}
} */

func BenchmarkFibonacciLoop100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci.FibonacciLoop(100)
	}
}

func BenchmarkFibonacciDynamic100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci.FibonacciDynamic(100)
	}
}

func BenchmarkFibonacciMemoization100(b *testing.B) {
	memoization := fibonacci.NewMemoization(100)

	for b.Loop() {
		memoization.Get(100)
	}
}
