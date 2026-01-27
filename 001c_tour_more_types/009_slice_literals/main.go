// Slice literals: creating slices with values, and that the type after [] can be any type including anonymous structs.
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// TAKE: at first this is mind bending BUT in fact what follows `[]`
	// is just the type of the slice elements. Instead of `int` or `bool`
	// you can have `struct {}` type
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
