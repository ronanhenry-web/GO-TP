package main

import "testing"

func BenchmarkFibonacciBase(b *testing.B) {
	for b.Loop() {
		FibonacciBase(100_000)
	}
}

func BenchmarkFibonacci3Vars(b *testing.B) {
	for b.Loop() {
		Fibonacci3Vars(100_000)
	}
}
