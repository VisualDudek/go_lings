// Package main demonstrates Go's explicit type conversion.
// Learn: Go requires explicit type conversion T(v) between types; no implicit coercion like in C.
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("Value:", x, f, z)
	fmt.Printf("Types: %T %T %T\n", x, f, z)
}
