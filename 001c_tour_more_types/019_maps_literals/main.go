package main

import "fmt"

type Point struct {
	X int
	Y int
}

var m = map[string]Point{
	"origin": {X: 0, Y: 0}, // same as origin: Point{X: 0, Y: 0}
	"one":    {X: 1, Y: 1},
	"two":    {X: 2, Y: 4},
}

// TAKE: if top-level type is just a type name, you can omit it from the elements of the literal.

func main() {
	fmt.Println(m)
}
