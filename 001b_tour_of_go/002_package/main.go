package main

import (
	"fmt"
	"math/rand" // "math/rand" refers to the "rand" package within the "math" standard library
)

func main() {
	fmt.Println("Welcome to the Tour of Go!")
	fmt.Println("A random number:", rand.Intn(100))
}
