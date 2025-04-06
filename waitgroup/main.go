package main

import (
	"fmt"
	"sync"
)

func Fibonacci3Vars(n int) {
	if n <= 1 {
		fmt.Println("Résultat pour", n, "=", n)
		return
	}

	a := 0
	b := 1
	c := 0
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
}

func FibonacciBase(n int) int {
	if n <= 1 {
		return n
	}

	var n1, n2 = 0, 1
	for i := 2; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}

func Fibonacci(n int, wg *sync.WaitGroup) int {
	defer wg.Done()
	if n <= 1 {
		return n
	}

	var n1, n2 = 0, 1
	for i := 2; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}

const countItems = 200

func main() {
	/* 	var wg sync.WaitGroup

	   	for i := 1; i < countItems; i++ {
	   		wg.Add(1)
	   		go func(i int) {

	   		}(i)
	   	}

	   	wg.Wait() */
	fmt.Println("Résultat pour", countItems, "=", FibonacciBase(countItems))
}
