// Package main demonstrates Go's constants.
// Learn: Constants can be declared with const keyword; they are package-level or block-scoped with values fixed at compile time.
// Constants cannot be declared using := syntax.
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "Hello, World"

	fmt.Println(World)
	fmt.Println("Pi:", Pi)
}
