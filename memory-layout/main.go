package main

import (
	"fmt"
	"unsafe"
)

// go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
// fieldalignment -fix ./...

type Foo struct {
	bar Bar
	s   string
	b   int64
	a   bool
	c   bool
	d   bool
}

type Bar struct {
	s string
	a bool
}

func main() {
	y := Foo{}

	fmt.Println(unsafe.Sizeof(y))
}
