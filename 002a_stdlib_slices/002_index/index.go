package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 1, 1, 2}

	idx := slices.Index(s, 2)

	fmt.Println("Index of 2 number in slice: ", idx)
}
