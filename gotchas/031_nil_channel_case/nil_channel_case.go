package main

import (
	"fmt"
	"time"
)

// Send and rec. operation on `nil` channel block forever.
// can be used to dynamically enable and disable case block (???)

func main() {
	inch := make(chan int)
	outch := make(chan int)

	go func() {
		var in <-chan int = inch // Receive-only channel, initially set to inch
		var out chan<- int       // Send-only channel, initially nil
		var val int              // Stores value between receiving and sending

		for {
			select {
			case out <- val: // Try to send value to out channel
				out = nil // After sending, disable this case by setting out to nil
				in = inch // Re-enable receiving from inch
			case val = <-in: // Try to receive value from in channel
				out = outch // After receiving, enable sending to outch
				in = nil    // Disable receiving by setting in to nil
			}
		}
	}()

	go func() {
		for r := range outch {
			fmt.Println("result:", r)
		}
	}()

	time.Sleep(0)
	inch <- 1
	inch <- 2
	time.Sleep(time.Second * 3)

}
