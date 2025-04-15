package main

import "fmt"

func generate() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {

	for v := range generate() { // tutaj kluczowe żeby załapać że idiomatyczne jest `v := range ch`
		fmt.Println(v)
	}
}
