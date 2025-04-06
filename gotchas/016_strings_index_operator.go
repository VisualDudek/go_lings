package gotchas

import "fmt"

// The index on a string returns a BYTE VALUE, not a character

func stringsIndexOperatorByte() {
	x := "text"
	fmt.Println(x[0]) //print 116
}

// if you need access specific string "characters" use the `for range` clause.
// r := []rune(x)
