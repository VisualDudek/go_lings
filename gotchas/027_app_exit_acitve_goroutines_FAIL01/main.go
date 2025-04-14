package main

import (
	"fmt"
	"sync"
	"time"
)

// App exit with active goroutines
// The app will not wait for all your goroutines to complete by it self

// SOLUTION: WaitGroup ALE po co chan i dlaczego powoduje deadlock ???

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	workerCount := 2

	for i := range workerCount {
		wg.Add(1)
		go doit(i, done, wg)
	}

	close(done)
	wg.Wait()
	fmt.Println("all done")
}

func doit(workerId int, done <-chan struct{}, wg sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()
	time.Sleep(time.Second * 1) // czy to jest non-blocking ?
	<-done
	fmt.Printf("[%v] is done\n", workerId)
}
