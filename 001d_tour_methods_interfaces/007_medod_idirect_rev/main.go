package main

// Methods with value receivers can be called on pointers; Go automatically dereferences them.
// This complements the reverse case where pointer receivers can be called on values.

import "math"

type Point struct {
	X, Y float64
}

// NOTE: value receiver
func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p := Point{3, 4}
	_ = p.Abs()

	v := &Point{6, 8}
	// TAKE: can call method with value receiver on pointer type, Go will dereference automatically
	_ = v.Abs()
}
