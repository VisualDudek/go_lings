package main

//lint:file-ignore * This is a learning exercise demonstrating go vet usecase

// Run go vet

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex //nolint
	value int
}

func main() {
	// Bug 1: Printf format mismatch
	name := "Alice"
	age := 30
	fmt.Printf("Name: %d, Age: %s\n", name, age) // %d expects int, %s expects string

	// Bug 2: Copying a mutex
	c1 := Counter{value: 10}
	c2 := c1 // This copies the mutex by value - dangerous!
	fmt.Println(c2.value)

	// Bug 3: Unreachable code
	return
	fmt.Println("This will never print") // Unreachable
}
