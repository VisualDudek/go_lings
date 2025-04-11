package main

import (
	"fmt"
	"time"
)

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
