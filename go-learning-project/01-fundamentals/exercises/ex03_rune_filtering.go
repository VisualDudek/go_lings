package main

// Exercise 3: Rune Filtering and String Transformation
//
// Task: Filter and transform strings based on rune properties
//
// 1. RemoveSpaces: Remove all whitespace runes from a string
// 2. KeepLetters: Keep only letter runes
// 3. AddSpaces: Add spaces between each rune (for readability)

// RemoveSpaces removes all space, tab, and newline runes
func RemoveSpaces(s string) string {
	// TODO: Implement this
	// Hint: Use a loop over []rune, filter out space runes
	// Space runes: ' ' (32), '\t' (9), '\n' (10), '\r' (13)
	return ""
}

// KeepLetters keeps only ASCII letters (a-z, A-Z) and returns them
// Note: This simplified version only handles ASCII
func KeepLetters(s string) string {
	// TODO: Implement this
	// Hint: Check if rune is in range 'a'-'z' or 'A'-'Z'
	// Ranges: 'a'=97, 'z'=122, 'A'=65, 'Z'=90
	return ""
}

// AddSpaces inserts a space between each rune
func AddSpaces(s string) string {
	// TODO: Implement this
	// Hint: Convert to []rune, join with spaces between each rune
	// Example: "Hi🎉" -> "H i 🎉"
	return ""
}

// Example usage (uncomment to test)
/*
func main() {
	// Test RemoveSpaces
	str1 := "Hello   World\t\nTest"
	fmt.Printf("Original: %q\n", str1)
	fmt.Printf("No spaces: %q\n", RemoveSpaces(str1))
	
	// Test KeepLetters
	str2 := "H3ll0, W0r1d! 🎉"
	fmt.Printf("\nOriginal: %s\n", str2)
	fmt.Printf("Only letters: %s\n", KeepLetters(str2))
	
	// Test AddSpaces
	str3 := "Hi🎉Go"
	fmt.Printf("\nOriginal: %s\n", str3)
	fmt.Printf("With spaces: %s\n", AddSpaces(str3))
}
*/
