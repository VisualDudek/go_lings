package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	for i := range 3 {
		go func(idx int) {
			fmt.Println("sending to ch: ", (idx+1)*2)
			ch <- (idx + 1) * 2 // TO JEST BLOCKING; SENDER is blocking
			fmt.Println("sent to ch: ", (idx+1)*2)
		}(i)
	}

	// get first result
	fmt.Println(<-ch)
	// close(ch) // not ok (you still have other senders)
	time.Sleep(time.Second * 3)
	// checking if sender is blocked ??? should not be, A JEDNAK JEST
	fmt.Println(<-ch)
	time.Sleep(time.Second * 1)
	fmt.Println(<-ch)
	time.Sleep(time.Second * 1)
}
