package main

import (
	"bruteforce/lib"
	"fmt"
)

func main() {
	fmt.Println("Password found: ", lib.FindPasswordByCombinationAsync("@kAl1"))
}
