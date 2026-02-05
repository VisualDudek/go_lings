package main

// Learn: The default case in select executes when no other cases are ready (non-blocking).
// time.Tick() creates a channel that sends repeatedly at intervals.
// time.After() creates a channel that sends once after a duration.

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}

	for {
		select {
		case <-tick:
			fmt.Println("tick.", elapsed())
		case <-boom:
			fmt.Println("BOOM!", elapsed())
			return
		default:
			fmt.Println("    .", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}
