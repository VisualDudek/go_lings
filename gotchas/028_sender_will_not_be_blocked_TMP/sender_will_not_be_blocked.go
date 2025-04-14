package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func(ch chan string) {
		for m := range ch {
			fmt.Println("processed:", m)
			time.Sleep(time.Second * 1)
		}
	}(ch)

	ch <- "one"
	ch <- "two" // !!! DLACZEGO teraz jest blocking ??? tylko dlatego że przekazałem
	// channel jako arg ? do gorutyny?
}
