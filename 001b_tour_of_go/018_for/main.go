// This code demonstrates traditional and modernized Go for loop syntax.
// Learn: for loop with init/condition/post (for i := 0; i < 10; i++),
// range-based iteration (for i := range 10), loop variable scope, and accumulating values.
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(i, sum)
	}

	fmt.Printf("\n-----for loop modernized-----\n\n")
	// for loop modernized
	sum2 := 0
	for i := range 10 {
		sum2 += i
		fmt.Println(i, sum2)
	}
}
