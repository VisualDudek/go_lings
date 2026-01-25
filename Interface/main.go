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

type Adder interface {
	Add(x int)
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

	var a Adder
	a = &v
	a.Add(1)

	// --- The following line would cause a compile error ---
	// a = v // Fail bc Add has pointer receiver and you are trying to pass value

}
