package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	// append another slice
	x := append(s, []int{4, 5, 6}...)

	fmt.Println(x)
}
