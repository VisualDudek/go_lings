package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5}
	fmt.Println("Len: ", len(s), " Capacity:", cap(s))

	sub := s[:2]
	sub = append(sub, 100) // niby zminiasz tylko w sub ale w sclice nigdy nie jest
	// to na 100% explicite

	fmt.Println(s) // strzał w stopę: [1, 2, 100, 4, 5]
	fmt.Println(sub)
}
