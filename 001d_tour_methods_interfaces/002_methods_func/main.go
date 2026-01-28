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

// TAKE: same same but different: function version of Abs
func Abs(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	p := Point{3, 4}
	fmt.Println(p.Abs())
	fmt.Println(Abs(p))
}
