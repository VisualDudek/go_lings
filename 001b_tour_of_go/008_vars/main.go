package main

import "fmt"

// TAKE: Variables can be declared without an initial value; they get the zero value of their type.
// For example, the zero value of an int is 0, and for a bool it is false.
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
