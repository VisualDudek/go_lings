package main

// Exercise 1: Basic Rune Manipulation
//
// Task: Complete the following functions to manipulate runes
//
// 1. CountRunes: Count the number of runes in a string (not bytes)
// 2. ReverseRunes: Reverse a string by rune (handles multi-byte characters)
// 3. ToUpperRune: Convert a single rune to uppercase (basic ASCII)

// CountRunes returns the number of runes in a string
func CountRunes(s string) int {
	// TODO: Implement this
	// Hint: Convert string to []rune and check length
	return 0
}

// ReverseRunes reverses a string by runes
func ReverseRunes(s string) string {
	// TODO: Implement this
	// Hint: Convert to []rune, reverse the slice, convert back to string
	return ""
}

// ToUpperRune converts a single ASCII rune to uppercase
func ToUpperRune(r rune) rune {
	// TODO: Implement this
	// Hint: Check if rune is lowercase letter (a-z), add 32 to uppercase
	// Ranges: 'a'=97, 'z'=122, 'A'=65, 'Z'=90
	return r
}

// Example usage (uncomment to test)
/*
func main() {
	// Test CountRunes
	str1 := "Hello, 世界!"
	fmt.Printf("String: %s\n", str1)
	fmt.Printf("Byte count: %d\n", len(str1))
	fmt.Printf("Rune count: %d\n", CountRunes(str1))

	// Test ReverseRunes
	str2 := "Hello"
	fmt.Printf("\nOriginal: %s\n", str2)
	fmt.Printf("Reversed: %s\n", ReverseRunes(str2))

	// Test ToUpperRune
	fmt.Printf("\nRune 'a' -> %c\n", ToUpperRune('a'))
	fmt.Printf("Rune 'z' -> %c\n", ToUpperRune('z'))
	fmt.Printf("Rune 'A' -> %c (no change)\n", ToUpperRune('A'))
	fmt.Printf("Rune '5' -> %c (no change)\n", ToUpperRune('5'))
}
*/
