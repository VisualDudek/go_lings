package main

// Learn: Closing channels with close() and using range to iterate over channels.
// Range on a channel continues until the channel is closed.
// Only the sender should close a channel.

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// idiomatic way to receive values until channel is closed
	for i := range c {
		fmt.Println(i)
	}
}
