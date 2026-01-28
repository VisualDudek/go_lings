// cSpell:words Abser
package main

// Interfaces define method sets that types satisfy implicitly.
// Interface variables hold a concrete value and its type. Receiver type matters:
// types with pointer receivers satisfy the interface at the pointer level only.
// the deal is that interface variable holds two things: the actual value and the type information about that value which is different from the value itself.

import "math"

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	// TAKE: interface variable can hold any value that implements the interface
	var a Abser
	f := MyFloat(-math.Sqrt2)
	p := Point{3, 4}

	// TAKE: put one type that implements interface into interface variable
	a = f  // a MyFloat implements Abser
	a = &p // a *Point implements Abser

	// TAKE: important detail: does method have pointer receiver or value receiver
	a = p // error: Point does not implement Abser ONLY *Point does (!!!)

	_ = a.Abs()
}
