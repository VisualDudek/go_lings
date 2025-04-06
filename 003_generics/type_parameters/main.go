package main

import "fmt"

// jeśli chcesz funkcje która przyjmuje więcej niż jeden z góry określony typ

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	// jak się nazywa "comparable"? constraint ??? i jakie są inne możliwości?
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a lice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

}
