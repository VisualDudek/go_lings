package main

import "fmt"

// the full slice expression protects against `append`
// try fix it using full slice expression

func main() {
	s := []int{1, 2, 3, 4}
	y := s[:2]

	fmt.Println(cap(s), cap(y))
	y = append(y, 5)
	fmt.Println("s: ", s)
	fmt.Println("y: ", y)

}
