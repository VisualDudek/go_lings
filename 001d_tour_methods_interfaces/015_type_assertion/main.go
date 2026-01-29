// Package main demonstrates type assertions on interface values.
//
// What you can learn:
// - Type assertion syntax: value := interfaceVar.(ConcreteType) panics if assertion fails.
// - The "comma ok" idiom: value, ok := interfaceVar.(ConcreteType) safely checks type.
// - Always use "comma ok" when you're uncertain about the interface's underlying type.
// - Type assertions extract the concrete value from an interface and verify its type.
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(ok, s)

	f, ok := i.(float64)
	fmt.Println(ok, f)

	// TAKE: and that's why you need to use the "comma ok" idiom when you're not sure about the underlying type
	f = i.(float64) // panic: interface conversion: interface {} is string, not float64

}
