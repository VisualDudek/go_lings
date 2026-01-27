// Slice length and capacity: using len() for current length
// and cap() for available capacity;
// slicing changes length and capacity while keeping the underlying array.
package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	// TAKE: when slicing from left idx, capacity changes too.
	s = s[2:]
	printSlice(s)
}
