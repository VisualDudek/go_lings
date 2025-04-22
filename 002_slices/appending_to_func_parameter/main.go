package main

import "fmt"

func modify(s []int) {
	s = append(s, 100) // This does NOT modify the caller's slice
	fmt.Println("slice s inside modify func", s, " Len: ", len(s), " Capacity:", cap(s))
}

// Poprawne modify zwraca zmodyfikowny slice
func properModify(s []int) []int {
	s = append(s, 100)
	return s
}

func main() {
	s := []int{1, 2, 3}

	fmt.Println(s, " Len: ", len(s), " Capacity:", cap(s))
	modify(s) // This will not modify slice
	fmt.Println(s, " Len: ", len(s), " Capacity:", cap(s))

}
