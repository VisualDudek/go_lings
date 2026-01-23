// TODO: Does Go support variadic functions?
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("3 + 5 =", add(3, 5))
}
