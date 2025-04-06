package main

import "fmt"

func escapeAnalysis(escapeMain *int) *int {
	escaped := 3
	escapeMain = &escaped

	return escapeMain
}

func NoEscape(notEscaped int) int {
	return notEscaped

}

// go build -gcflags="-m" -o escape_analysis main.go
func main() {
	escapeMain := 1
	notEscaped := 2

	escapeAnalysis(&escapeMain)

	NoEscape(notEscaped)
	fmt.Println(escapeMain)
}
