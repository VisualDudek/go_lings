// Package main demonstrates that methods can be defined on any named type, not just structs.
// This enables extending built-in types (like float64) with custom behavior while maintaining type safety.
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// TAKE: you can declare methods on non-struct types too (!!!)
// TAKE: You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)
	fmt.Println(f.Abs())
}
