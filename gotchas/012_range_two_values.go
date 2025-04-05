package gotchas

import "fmt"

// jedynie dla przypomnienia, range zwraca dwie wartości,
//możesz odebrać jedną lub dwie

func rangeTwoValues() {
	x := []string{"a", "b", "c"}

	for v := range x { // pułapka w nazwie zmiennej, lepiej idx
		fmt.Println(v) // prints 0, 1, 2
	}

	for _, v := range x {
		fmt.Println(v) // prints a, b, c
	}
}
