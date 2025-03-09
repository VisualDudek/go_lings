package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s []int

	fmt.Println(s, " capacity: ", cap(s))

	// for i := 0; i < 10; i++ {
	// ^^^ modernize
	for i := range 10 {
		s = append(s, i)

		fmt.Println(s, " capacity: ", cap(s), " Underlying array address: ", unsafe.SliceData(s))
	}
	// can lead to inefficient memory usage

}
