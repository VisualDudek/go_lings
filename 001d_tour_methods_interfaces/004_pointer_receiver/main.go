// Methods with pointer receivers can modify the original struct,
// while value receivers operate on a copy of the struct.
// Use pointer receivers when the method needs to mutate the receiver.
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// TAKE: pointer receiver allows method to modify the original struct
func (p *Point) Scale(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{3, 4}
	fmt.Println(p)
	p.Scale(10)

	fmt.Println(p)
}
