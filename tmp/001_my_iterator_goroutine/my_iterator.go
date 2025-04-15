package main

import "fmt"

func iter(n int, c chan int) {
	for {
		c <- n
		n += 1
	}
}

func main() {
	fmt.Println("hello")

	c := make(chan int)

	go iter(1, c)

	fmt.Println(<-c, " iter")
	fmt.Println(<-c, " iter")
	fmt.Println(<-c, " iter")
}
