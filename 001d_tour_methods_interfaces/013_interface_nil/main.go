// Package main demonstrates calling methods on a nil interface.
//
// What you can learn:
// - A nil interface has no type information and cannot dispatch method calls.
// - Calling a method on a nil interface causes a runtime panic.
// - A nil interface is safe to pass as an argument, but unsafe to use if methods are called.
package main

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	// TAKE: An uninitialized interface variable is nil (both type and value are absent).
	describe(i) // will NOT panic
	i.M()       // will panic
}
