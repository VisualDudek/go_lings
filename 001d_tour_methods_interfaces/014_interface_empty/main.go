// Package main demonstrates the empty interface.
//
// What you can learn:
// - The empty interface interface{} can hold values of any type.
// - Empty interfaces are used when type flexibility is needed, like in generic functions.
// - Runtime type information is preserved for value types stored in empty interfaces.
// - Any type, built-in or custom, satisfies the empty interface.
package main

import "fmt"

func main() {
	// TALE: empty interface can hold any value of any type
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
