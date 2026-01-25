// Package main demonstrates Go's type inference with short variable declaration.
// Learn: Types are inferred from assigned values; numeric constants infer to int, float64, or complex128 based on literal syntax.
package main

import "fmt"

func main() {
	var i int
	j := i
	a := 42
	b := 3.14
	g := 0.8 + 0.5i

	fmt.Println("Value:", i, j, a, b, g)
	fmt.Printf("Types: %T %T %T %T %T\n", i, j, a, b, g)
}
