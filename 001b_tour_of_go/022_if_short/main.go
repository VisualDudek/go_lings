// This code demonstrates if statement with short variable declaration.
// Learn: variables declared in if statement (if v := ...) are scoped only to the if block, pattern for conditional computation.
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// TAKE: v variable only scoped to this if statement
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
