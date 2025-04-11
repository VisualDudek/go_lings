package main

import (
	"fmt"
	"time"
)

// Both sender and receiver must be ready at the same time for the message
// to be passed.
// KLUCZOWE DO ZROZUMIENIA:
// if the producer tries to send but no one is ready to receive, IT WAITS.
// if the consumer tires to receive but no message is there yet, IT WAITS.
// THAT'S WHY WE SAY THE COMMUNICATION IS SYNCHRONOUS! NOT ASYNC

func producer(ch chan string) {
	for i := range 5 {
		msg := fmt.Sprintf("Message %d", i)
		ch <- msg
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func consumer(ch chan string) {
	for msg := range ch {
		fmt.Println("Received: ", msg)
	}
}

func main() {
	ch := make(chan string)

	go producer(ch)
	consumer(ch)

}
