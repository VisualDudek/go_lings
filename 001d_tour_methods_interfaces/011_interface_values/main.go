// Package main demonstrates interface values in Go.
//
// What you can learn:
// - An interface variable holds a concrete type and value, not the interface type itself.
// - Using %T format verb shows the concrete runtime type stored in an interface variable.
// - Different concrete types satisfying the same interface can be assigned to the same interface variable.
// - Interface variables can hold values from different types that implement the interface methods.
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	println(t.S)
}

type F float64

func (f F) M() {
	println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	i = &T{"hello"}
	// TAKE: will show that i holds a *T, not interface type I
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}
