package main

// Functions can accept pointer parameters using the * operator.
// Use the & operator to pass the address of a value when calling a function that expects a pointer.
// This allows the function to modify the original value.

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

func Scale(p *Point, factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{5, 8}
	fmt.Println(p)
	// TAKE: when passing pointer to function remember to use & operator
	Scale(&p, 10)
	fmt.Println(p)
}
