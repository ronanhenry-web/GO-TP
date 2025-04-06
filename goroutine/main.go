package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	for char := 'A'; char <= 'E'; char++ {
		fmt.Println(string(char))
	}
}

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(time.Second)
}
