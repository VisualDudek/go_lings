// Slices of slices:
// creating multi-dimensional slices with nested slice literals
// and accessing elements with multiple indices.
package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for _, s := range board {
		fmt.Printf("%s\n", strings.Join(s, " "))
	}
}
