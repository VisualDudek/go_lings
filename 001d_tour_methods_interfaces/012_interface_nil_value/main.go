// Package main demonstrates interface values holding nil concrete values.
//
// What you can learn:
// - An interface holding a nil concrete value is NOT the same as a nil interface.
// - An interface stores (type, value) pair; if a nil pointer is assigned, type is non-nil but value is nil.
// - Methods can be called on an interface with a nil value; the receiver will be nil.
// - An uninitialized interface variable is nil (both type and value are absent).
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		println("<nil>")
		return
	}
	println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	var t *T // t is nil
	// TAKE: Note that an interface value holds a nil concrete value is itself NOT nil.
	// bc. interface as "tuple" holds (type, value) where type is *T and value is nil
	i = t // NOTE: i is not nil
	describe(i)
	i.M()

	i = &T{"hello"} // not nil
	describe(i)
	i.M()
}
