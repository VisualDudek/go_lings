package main

import "fmt"

type Vector struct {
	X, Y int
}

func (v *Vector) Abs() int {
	return v.X + v.Y
}

func (v *Vector) Add(x int) {
	v.X += x
	v.Y += x
}

func main() {
	v := Vector{
		X: 3,
		Y: 4,
	}

	fmt.Printf("Vector v: %+v \n", v)

	fmt.Printf("Use method Abs on Vector v: %v \n", v.Abs())

	v.Add(1)
	fmt.Printf("Vector v after used method Add: %+v \n", v)
}
