package main

import "fmt"

// TAKE: Cannot use := outside of functions.

func main() {
	// TAKE: Short variable declaration using := infers the type from the initial value.
	// It can only be used inside functions.
	k := 3

	fmt.Println("k:", k)
}
