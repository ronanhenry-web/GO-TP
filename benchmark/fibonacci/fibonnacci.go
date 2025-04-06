package fibonacci

import (
	"reflect"
)

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// FibonacciLoop Loop to calculate FibonacciRecursive numbers
func FibonacciLoop(n int) int {
	if n <= 1 {
		return n
	}

	var n1, n2 = 0, 1
	for i := 2; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}

// FibonacciDynamic Dynamic programming to calculate FibonacciRecursive numbers
func FibonacciDynamic(n int) int {
	if n <= 1 {
		return n
	}
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

type FibonacciMemoization struct {
	memo []int
}

func NewMemoization(size int) FibonacciMemoization {
	fibonacci := FibonacciMemoization{
		memo: make([]int, size),
	}
	fibonacci.Get(size)
	return fibonacci
}

func (f FibonacciMemoization) Size() int {
	return int(uintptr(len(f.memo)) * reflect.TypeOf(f.memo).Elem().Size())
}

func (f FibonacciMemoization) Get(n int) int {
	if n > len(f.memo) {
		newMemo := make([]int, n)
		copy(newMemo, f.memo)
		f.memo = newMemo
	}

	if n <= 1 {
		return n
	}

	if val := f.memo[n-1]; val != 0 {
		return val
	}
	if n <= 1 {
		f.memo[n-1] = n
	} else {
		var n1, n2 = 0, 1
		for i := 2; i <= n; i++ {
			n1, n2 = n2, n1+n2
		}

		f.memo[n-1] = n2
	}
	return f.memo[n-1]
}
