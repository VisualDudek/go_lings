package gotchas

import (
	"fmt"
	"unicode/utf8"
)

// `len()` returns the number of bytes instead of the number of characters.
// USE `RuneCountString()`
// BUT single character may span multiple runes.

func StringLenght() {
	data := "♥"
	fmt.Println(len(data)) // prints: 3

	fmt.Println(utf8.RuneCountInString(data)) // prints: 1

	fmt.Println(utf8.RuneCountInString("é")) // prints: 2 (!!!)
}
