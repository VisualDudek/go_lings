package gotchas

import "fmt"

// jeśli chcesz zmienić tekst to zamień go na slice of runes

func stringsAreImmutable() {
	x := "text"

	xrunes := []rune(x)

	xrunes[0] = 'T'

	fmt.Println(string(xrunes)) // prints Text
}
