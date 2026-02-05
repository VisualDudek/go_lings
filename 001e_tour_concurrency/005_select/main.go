package main

// Learn: The select statement waits on multiple channel operations.
// It executes whichever case is ready first.
// Use select to coordinate communication between multiple channels.

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // send x to channel c
			x, y = y, x+y
		case <-quit: // receive from channel quit
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// e.g. of anonymous goroutine
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
