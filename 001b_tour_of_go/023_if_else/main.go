// This code demonstrates if-else statement with short variable declaration.
// Learn: if-else syntax, variable scoping in if/else blocks (v only available in if/else),
// else clause execution when condition is false.
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// TAKE: this will be printed before fmt.Println in main
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20), // expect output: "27 >= 20" will be printed before fmt.Println in main
	)
}
