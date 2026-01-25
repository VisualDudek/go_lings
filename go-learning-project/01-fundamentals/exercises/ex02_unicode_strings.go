package main

// import "unicode/utf8"

// Exercise 2: Unicode String Processing
//
// Task: Work with Unicode strings correctly using runes
//
// 1. TakeFirst: Take first N runes from a string
// 2. DropFirst: Drop first N runes from a string
// 3. EachRuneIndex: Find the byte index of the Nth rune

// TakeFirst returns the first n runes of a string
func TakeFirst(s string, n int) string {
	// TODO: Implement this
	// Hint: Convert to []rune, slice it, convert back
	// Handle case where n > length
	return ""
}

// DropFirst skips the first n runes and returns the rest
func DropFirst(s string, n int) string {
	// TODO: Implement this
	// Hint: Convert to []rune, slice from n onwards, convert back
	// Handle case where n >= length
	return ""
}

// EachRuneIndex returns a slice of byte indices for each rune start position
func EachRuneIndex(s string) []int {
	// TODO: Implement this
	// Hint: Use utf8.DecodeRuneInString to iterate and track byte positions
	// Example: "Hi🎉" should return [0, 2, 3] (positions where each rune starts)
	return nil
}

// Example usage (uncomment to test)
/*
func main() {
	str := "Hello, 世界! 🎉"

	// Test TakeFirst
	fmt.Printf("Original: %s\n", str)
	fmt.Printf("First 5 runes: %s\n", TakeFirst(str, 5))
	fmt.Printf("First 10 runes: %s\n", TakeFirst(str, 10))
	fmt.Printf("First 3 runes: %s\n", TakeFirst(str, 3))

	// Test DropFirst
	fmt.Printf("\nDrop first 5 runes: %s\n", DropFirst(str, 5))
	fmt.Printf("Drop first 10 runes: %s\n", DropFirst(str, 10))

	// Test EachRuneIndex
	str2 := "Hi🎉"
	fmt.Printf("\nString: %s\n", str2)
	fmt.Printf("Rune indices: %v\n", EachRuneIndex(str2))
}
*/
