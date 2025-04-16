package main

import "fmt"

// Send and rec. operation on `nil` channel block forever.

func main() {
	var ch chan int

	for i := range 3 {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	fmt.Print("result:", <-ch) // deadlock
}
