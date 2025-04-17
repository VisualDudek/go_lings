package main

// dlaczego nie blokuje po 2 ? tylko po 4 ?

import (
	"fmt"
	"time"
)

func producer(ch chan string) {
	for i := 1; i <= 6; i++ {
		msg := fmt.Sprintf("Message %d", i)
		fmt.Println(time.Now().Format("15:04:05.0000000"), "Sending:", msg)
		ch <- msg // this won't block unless the buffer is full
		fmt.Println(time.Now().Format("15:04:05.0000000"), "Sent:", msg)
	}
	close(ch)
}

func consumer(ch chan string) {
	for msg := range ch {
		fmt.Println(time.Now().Format("15:04:05.0000000"), "Received:", msg)
		time.Sleep(time.Second * 2) // simulate slow processing
	}
}

func main() {
	ch := make(chan string, 4) // buffered channel with capacity 2
	fmt.Println("Capacity of ch: ", cap(ch))

	go producer(ch)
	consumer(ch)
}
