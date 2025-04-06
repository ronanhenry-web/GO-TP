package main

import (
	"fmt"
)

// Multiply matrices A and B and store the result in C
func matrixMultiply(A, B, C [][]int) {
	for i := range A {
		for j := range B[0] {
			for k := range B {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
}

func main() {
	// Example matrices
	A := [][]int{{1, 2}, {3, 4}}
	B := [][]int{{5, 6}, {7, 8}}
	C := make([][]int, len(A))
	for i := range A {
		C[i] = make([]int, len(B[0]))
	}

	matrixMultiply(A, B, C)
	fmt.Println(C)
}
