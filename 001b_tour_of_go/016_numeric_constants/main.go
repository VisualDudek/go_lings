// Package main demonstrates Go's untyped numeric constants.
// Learn: Untyped constants can hold very large values without overflow; they take on the type required by their context.
package main

import (
	"fmt"
	"math"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needFloat(x float64) float64 {
	return x
}

func main() {
	// An untyped constant takes the type needed by its context.

	// fmt.Println("Big:", Big)
	fmt.Println("Small:", Small)
	fmt.Println("float64 Big:", needFloat(Big))
	fmt.Println("Max float64:", math.MaxFloat64)
}
