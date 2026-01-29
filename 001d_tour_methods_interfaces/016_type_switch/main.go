// Package main demonstrates type switches.
//
// What you can learn:
// - `any` is a type alias for `interface{}`.
// - Type switch syntax: switch v := i.(type) { case Type1: ... case Type2: ... }
// - In each case, v has the concrete type of that case, not interface{}.
// - The default case handles types that don't match any explicit case.
package main

import "fmt"

// TAKE: `any` is type alias for `interface{}`
func do(i any) {
	// TAKE: type switch syntax, really weird looking at first
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
