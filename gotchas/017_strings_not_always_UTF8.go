package gotchas

// String values are not required to be UTF8 text. They can contain arbitrary
//bytes. The only time stings are UTF8 ws when string literal are used.
// BUT Even than thay can include other data using excape sequences.

// USE `ValidString()` to know if you have valid UTF8 text

import (
	"fmt"
	"unicode/utf8"
)

func StringNotAlwaysUTF8() {

	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1))

	data2 := "A\xtec"
	fmt.Println(utf8.ValidString(data2)) // prints: false
}
