package main

import (
	"fmt"
	"time"
)

func writeToChannel(ch chan<- int, value int) {
	ch <- value
}

func readFromChannel(ch <-chan int) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)

	go writeToChannel(ch, 42)
	go readFromChannel(ch)

	time.Sleep(time.Second)
}
