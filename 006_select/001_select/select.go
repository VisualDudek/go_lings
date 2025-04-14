package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // czy to się wykona jedynie gdy receiver będzie ready?
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := range 10 {
			_ = i
			fmt.Println(<-c) // czy dobrze to sobie interpretuje jako
			// ustawienie receiver na ready ???
		}
		quit <- 0 // kanał do komunikacji
	}()
	fibonacci(c, quit)
}
