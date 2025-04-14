package main

import (
	"fmt"
	"sync"
	"time"
)

// App exit with active goroutines
// The app will not wait for all your goroutines to complete by it self

// DLACZEGO NIE MOżNA W PROSTY SPOSÓB:
// po co chan w legacy solution ???j

func doit(workerId int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[%v] is running\n", workerId)
	time.Sleep(time.Second * 1) // czy to jest non-blocking ?
	fmt.Printf("[%v] is done\n", workerId)
}

func main() {
	var wg sync.WaitGroup
	workerCount := 2

	for i := range workerCount {
		wg.Add(1)
		go doit(i, &wg)
	}

	wg.Wait()
	fmt.Println("all done")
}
