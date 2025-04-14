package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	for i := range 3 {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	// get first result
	fmt.Println(<-ch)
	close(ch) // not ok (you still have other senders)

	time.Sleep(time.Second * 2)
}
