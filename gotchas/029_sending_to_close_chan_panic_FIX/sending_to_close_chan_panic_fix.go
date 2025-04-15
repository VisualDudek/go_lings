package main

import (
	"fmt"
	"time"
)

// FIX create special channel to signal the remaining workers that their reuslts
// are no longer needed.

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	for i := range 3 {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, " sent result")
			case <-done:
				fmt.Println(idx, " exiting")
			}
		}(i)
	}

	// get first result AAAA ALE Uważaj żebyś nie pomyślał że jest zachowana
	// kolejność, to są wyniki z trzech a nie z jednej gorutyny więc
	// kilejność może być dowolna
	fmt.Println("result: ", <-ch)
	fmt.Println("result: ", <-ch)
	close(done)

	time.Sleep(time.Second * 2)
}
