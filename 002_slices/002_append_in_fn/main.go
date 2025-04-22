package main

import "fmt"

func appendToSlice(s []int, num int) {
	s = append(s, num)
}

func main() {
	s := []int{1, 2, 3}

	fmt.Printf("Before append: %#v\n", s)
	fmt.Println("calling appendToSlice")
	appendToSlice(s, 4)
	// Dlaczego 4 nie została dodana do slice s ?
	fmt.Printf("After append: %#v\n", s)

}
