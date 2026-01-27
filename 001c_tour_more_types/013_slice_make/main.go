// make() builtin: creating slices with make(type, length)
// and make(type, length, capacity) to allocate underlying arrays.
package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	// TAKE: allocate a zeroed array and return a slice that refers to it.
	// set len at 5
	s := make([]int, 5)
	printSlice(s)

	s = make([]int, 0, 5)
	printSlice(s)

	s = s[:2]
	printSlice(s)

	s = s[2:5]
	printSlice(s)
}
