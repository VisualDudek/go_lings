package main

// Use pointer receivers even for read-only methods on large structs to avoid copying overhead.
// This is a performance optimization consideration when the struct is expensive to copy.

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// NOTE: pointer receiver even if method does not modify the struct, avoid copying
func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p := Point{3, 4}
	_ = p.Abs()

	fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)
}
