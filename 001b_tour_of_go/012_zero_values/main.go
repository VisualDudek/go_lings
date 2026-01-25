// Package main demonstrates Go's zero values for basic types.
// Learn: Uninitialized variables have zero values: int=0, float64=0.0, bool=false, string="".
package main

import "fmt"

func main() {
	var (
		i int
		f float64
		b bool
		s string
	)

	fmt.Printf("i: %v \nf: %v \nb: %v \ns: %q\n", i, f, b, s)
}
