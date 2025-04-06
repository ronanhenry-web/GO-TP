package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Print("Hello ")
		}
	}()

	time.Sleep(time.Second)

	fmt.Print("World")
}
