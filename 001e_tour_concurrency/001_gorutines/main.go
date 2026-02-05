package main

// Learn: How to spawn goroutines with the `go` keyword.
// Goroutines execute concurrently with the main function.
// The main function does not wait for goroutines to complete.

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
