/*
Package main demonstrates multiple return values in Go.

Learn:
- Functions can return multiple values by listing types in the return signature: (string, string)
- Multiple return values are unpacked when calling the function: a, b := swap(...)
- Go uses named return value syntax (type1, type2) without parameter names
*/
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
