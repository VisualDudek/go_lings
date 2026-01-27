package main

import "fmt"

func main() {
	// TAKE: no underlying array; zero length and capacity.
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}
