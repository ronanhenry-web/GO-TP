package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/Chroq/benchmark/fibonacci"
)

func main() {
	fmt.Println(fibonacci.FibonacciLoop(40))
}

/*
func main() {
	// profiles := profile.Start(profile.MemProfile)
	// defer profiles.Stop()

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	// memutils.DisplayMemoryUsage()
	// fmt.Println("Initial Memory Usage")

	// FibonacciRecursive(10)
	// runtime.GC()
	go func() {
		i := 0
		for {
			fibonacci.FibonacciRecursive(i)
			i++
		}
	}()
	// memutils.DisplayMemoryUsage()
	// fmt.Println("FibonacciRecursive")

	// FibonacciLoop(10)
	// runtime.GC()
	go func() {
		i := 0
		for {
			fibonacci.FibonacciLoop(i)
			i++
		}
	}()
	// memutils.DisplayMemoryUsage()
	// fmt.Println("FibonacciLoop")

	// FibonacciDynamic(10)
	// runtime.GC()
	go func() {
		i := 0
		for {
			fibonacci.FibonacciDynamic(i)
			i++
		}
	}()
	// memutils.DisplayMemoryUsage()
	// fmt.Println("FibonacciDynamic")

	// FibonacciMemoization(10)
	// runtime.GC()
	go func() {
		memo := fibonacci.NewMemoization(100)
		i := 0
		for {
			memo.Get(i)
			i++
		}
	}()
	// memutils.DisplayMemoryUsage()
	// fmt.Println("FibonacciMemoization")
	// fmt.Println("Memoization size : ", memo.Size(), " bytes")
	wg.Wait()
}
*/
