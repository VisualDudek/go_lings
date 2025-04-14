package main

import (
	"fmt"
	"time"
)

// App exit with active goroutines
// The app will not wait for all your goroutines to complete by it self

func main() {
	workerCount := 2

	for i := range workerCount {
		go doit(i)
	}
	time.Sleep(time.Second * 1)
	fmt.Println("all done")
}

func doit(workerId int) {
	fmt.Printf("[%v] is running\n", workerId)
	time.Sleep(time.Second * 3)
	fmt.Printf("[%v] is done\n", workerId) // will never be printed
}
