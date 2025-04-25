package main

import "fmt"

// clear to zero value

func main() {
	s := []int{1, 2, 3}

	fmt.Println(s)

	clear(s)

	fmt.Println(s)
}
