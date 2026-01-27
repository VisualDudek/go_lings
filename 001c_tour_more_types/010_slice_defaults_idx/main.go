// Slice indices: omitting the low index defaults to 0, omitting the high index defaults to len(s), and omitting both gives the whole slice.
package main

import "fmt"

// TAKE: slicing with default low and high indices can be skipped.
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	// TAKE: these slice expr. are equivalent:
	// s[0:10]
	// s[:10]
	// s[0:]
	// s[:]
}
