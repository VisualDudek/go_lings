// Learn: defers pushed in a loop execute in LIFO order, stack behavior.
package main

import "fmt"

func main() {
	fmt.Println("Counting")

	for i := range 10 {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
