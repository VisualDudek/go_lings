package main

import "fmt"

type MyErr int

func (e MyErr) Error() string {
	return fmt.Sprint("oj zaraz wpadnę w recursiv call: ", e) // recursive call
	// BC Interface Priority
	// The fmt package checks for error implementations before stringers
	// 1. error interface: First priority for formating
	// 2. Stringer interface: Secondarty priority
	//
	// fmt.Sprint(e) -> checks if 'e' implements 'error' -> calls 'e.Error()'
	//
	// SOLUTION
	//
	// return fmt.Sprint("cast z error na int: ", int(e)) // recursive call
}

func main() {
	err := MyErr(5)
	fmt.Println(err)
}
