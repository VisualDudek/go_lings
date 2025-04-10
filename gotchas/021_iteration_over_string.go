package gotchas

import "fmt"

// The `for range` clauses with string varivables will try to interpret
// the data as UTF8 text. For any byte
// sequences it doesn't understand it will return 0xfffd runes
// (aka unicode replacement characters) instead of
// the actual data. If you have arbitrary (non-UTF8 text) data
// stored in your string variables, make sure to
// convert them to byte slices to get all stored data as is.

func IterationOverString() {
	data := "A\xfe\x02\xff\x04"
	for _, v := range data {
		fmt.Printf("%#x ", v)
	}
	// prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)

	fmt.Println()
	for _, v := range []byte(data) {
		fmt.Printf("%#x ", v)
	}
	// prints: 0x41 0xfe 0x2 0xff 0x4 (good)
}
