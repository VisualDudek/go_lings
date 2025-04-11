package main

import (
	"fmt"
	"time"
)

func worker(ch chan rune) {
	defer close(ch)
	for _, v := range []rune{'A', 'B', 'C', 'D', 'E', 'F'} {
		fmt.Println("Sendig letter: ", v)
		ch <- v
	}

}

func consumer(ch chan rune) {
	for msg := range ch {
		time.Sleep(time.Second) // simulate slow processing
		fmt.Println("Received:", msg)
	}
}

func main() {
	ch := make(chan rune, 1)

	go worker(ch)
	consumer(ch)
}
