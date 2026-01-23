// TAKE: Peek where math.Pi lives in the math package
// Ctrl+Shift+F10 to peek
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to the Tour of Go!")
	fmt.Printf("The pi is %v\n", math.Pi) // math.pi lowercase 'p' is not exported
}
