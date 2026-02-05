package main

// Learn: Buffered channels with capacity specified in make().
// Sends block only when the buffer is full; receives block when empty.
// Without goroutines to drain the channel, filling the buffer causes deadlock.

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// can end up in a deadlock if buffer is full

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
