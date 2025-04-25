package main

import (
	"fmt"
)

func main() {
	// allocate each []int inside the [][]int

	s := make([][]int, 3)

	for idx := range s {
		s[idx] = make([]int, 3)
	}

	fmt.Println(s)

	// naive print
	for _, v := range s {
		fmt.Println(v)
	}

}
