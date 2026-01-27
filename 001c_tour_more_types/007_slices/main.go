// Slices: dynamic-length sequences declared without fixed size,
// slice literals, and slicing syntax with [low:high] indices.
package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
