package main

import (
	"fmt"
	"strings"
)

func main() {
	// create slice of slices with big X

	X := [][]string{
		[]string{"X", "_", "X"},
		[]string{"_", "X", "_"},
		[]string{"X", "_", "X"},
	}

	fmt.Println(X)

	// naive print
	for _, v := range X {
		fmt.Println(v)
	}

	// nice one
	for _, v := range X {
		fmt.Printf("%s\n", strings.Join(v, " "))
	}

}
