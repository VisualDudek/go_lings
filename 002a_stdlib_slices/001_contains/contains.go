package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3, 4}

	if slices.Contains(s, 3) {
		fmt.Println("s contain 3")
	}
}
