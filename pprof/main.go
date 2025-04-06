package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
func consumeCPU() {
	for {
		for i := 0; i < 1000000; i++ {
		}
	}
}

func lessConsumingCPU() {
	for {
		for i := 0; i < 10000; i++ {
		}
	}
}

// go tool pprof http://localhost:6060/debug/pprof/heap
// go tool pprof http://localhost:6060/debug/pprof/allocs
func leakMemory() {
	var leak []string
	for {
		leak = append(leak, "leak memory")
		time.Sleep(1 * time.Millisecond)
	}
}

func noMemoryLeak() {
	var memory string
	for {
		time.Sleep(1 * time.Second)
		memory = "no leak"
		_ = memory
	}
}

func veryBigMemoryLeak() {
	var leak [][]byte
	for {
		leak = append(leak, make([]byte, 10))
		time.Sleep(1 * time.Millisecond)
	}
}

func hugeMemoryLeak() {
	var leak [][]byte
	for {
		leak = append(leak, make([]byte, 1000))
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	go leakMemory()
	go veryBigMemoryLeak()
	go noMemoryLeak()
	go hugeMemoryLeak()

	time.Sleep(60 * time.Minute)
}
