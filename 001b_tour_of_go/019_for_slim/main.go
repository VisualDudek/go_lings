// This code demonstrates a for loop with only a condition (no init or post statements).
// Learn: for loop as a while-like construct (for condition {}), manual loop variable management outside the loop.
package main

import "fmt"

func main() {
	i := 0
	// equivalent of while loop
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
