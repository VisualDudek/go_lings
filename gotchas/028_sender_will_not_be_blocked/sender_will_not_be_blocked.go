package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()

	ch <- "one"
	ch <- "two" // Sender will not be block !!!
	// jeśli receiver nie będzie gotowy do odbioru to payload idze w kosmos
	// rozwiązaniem jest buffer channel
}
