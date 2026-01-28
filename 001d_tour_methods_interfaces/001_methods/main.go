package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// TAKE: methods syntax is a func with receiver, instead of func name, and method name with return type
func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p := Point{3, 4}
	fmt.Println(p.Abs())
}
