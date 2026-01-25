// Package main demonstrates Go's basic types and factored variable declarations.
// Learn: Grouping var declarations with block syntax, basic types (bool, uint64, complex128),
// bit shift operators for numeric calculations, and type introspection with fmt.Printf.
package main

import (
	"fmt"
	"math/cmplx"
)

// TAKE: you can also factor into block `var` declarations
var (
	ToBe bool = false
	// ASK: interesting how wise bit shift is calculated first and then put into uint64
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
