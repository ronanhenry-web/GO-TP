package main

import "fmt"

type FibonacciMemoization struct {
	memo []int
}

func NewFibonacci(size int) FibonacciMemoization {
	fibonacci := FibonacciMemoization{
		memo: make([]int, size),
	}
	fibonacci.Get(size)
	return fibonacci
}

func (f FibonacciMemoization) Get(n int) int {
	if n > len(f.memo) {
		newMemo := make([]int, n)
		copy(newMemo, f.memo)
		f.memo = newMemo
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

func main() {
	memoFib := NewFibonacci(50)

	fmt.Println(memoFib.Get(50))
	fmt.Println(memoFib.Get(50))
	fmt.Println(memoFib.Get(50))
}
