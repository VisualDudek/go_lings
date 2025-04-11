package main

// jakim cudem tutaj działą przewidywalnie a w `buffer_cah_async.go` można
// przekazać do buffer channel aż 4 elementy ???

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)

	go func() {
		defer close(ch)
		ch <- "one"
		fmt.Println("sent one")
		ch <- "two"
		fmt.Println("sent two")
		ch <- "three"
		fmt.Println("sent three")
		ch <- "four"
		fmt.Println("sent four")
	}()

	time.Sleep(time.Second)

	// Start receiving
	// for i := 0; i < 4; i++ {
	// 	msg := <-ch
	// 	fmt.Println("received:", msg)
	// }

	for msg := range ch {
		fmt.Println("received:", msg)
	}
}
