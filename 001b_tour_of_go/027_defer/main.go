// Learn: defer schedules a call to run after the surrounding function returns.
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
