package main

import "fmt"

// In Go, functions can share parameter types.
// x int, b int is equivalent to:
// x, b int
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("3 + 5 =", add(3, 5))
}
