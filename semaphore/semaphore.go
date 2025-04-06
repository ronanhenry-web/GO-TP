package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	maxWorkers := runtime.NumCPU()
	sem := make(chan struct{}, maxWorkers)

	for i := 0; i < 100; i++ {
		sem <- struct{}{} // Acquiert un jeton
		go func(id int) {
			defer func() { <-sem }() // Libère un jeton
			time.Sleep(time.Millisecond * 100)
			fmt.Println("Goroutine ", id, " done : sem len = ", len(sem))
		}(i)
	}
	// Attend que tous les "jetons" soient libérés
	for i := 0; i < maxWorkers; i++ {
		sem <- struct{}{}
	}
}
