// Package main demonstrates runes in Go.
// Learn: Runes are int32 values representing Unicode code points; strings are UTF-8 encoded byte sequences;
// rune literals use single quotes; range over strings yields runes, not bytes; handle Unicode properly with runes.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Rune is an alias for int32, representing a single Unicode character
	var r rune = 'A'
	fmt.Printf("Rune 'A': value=%d, type=%T\n", r, r)

	// String vs rune: strings are UTF-8 encoded bytes, runes are code points
	str := "Hello"
	fmt.Println("\n--- Iterating over bytes vs runes ---")
	fmt.Printf("String length (bytes): %d\n", len(str))
	fmt.Printf("String length (runes): %d\n", utf8.RuneCountInString(str))

	// Range over string yields runes (code points)
	fmt.Println("\nRange over string (runes):")
	for i, r := range str {
		fmt.Printf("  Index %d: rune=%q (%U), decimal=%d\n", i, r, r, r)
	}

	// Unicode characters: runes handle multi-byte UTF-8 sequences
	text := "Hello, 世界! 🎉"
	fmt.Printf("\n--- Unicode string ---\n")
	fmt.Printf("String: %s\n", text)
	fmt.Printf("Byte length: %d\n", len(text))
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(text))

	fmt.Println("\nRange over Unicode string:")
	for i, r := range text {
		fmt.Printf("  Byte index %d: %q (%U)\n", i, r, r)
	}

	// Working with rune slices
	fmt.Println("\n--- Converting string to rune slice ---")
	runes := []rune(text)
	fmt.Printf("Rune slice length: %d\n", len(runes))
	for i, r := range runes {
		fmt.Printf("  Rune %d: %q (%U)\n", i, r, r)
	}

	// Rune literals and escape sequences
	fmt.Println("\n--- Rune literals and escapes ---")
	fmt.Printf("Rune literal 'A': %q\n", 'A')
	fmt.Printf("Unicode escape \\u0041 (A): %q\n", '\u0041')
	fmt.Printf("Unicode escape \\U0001F389 (🎉): %q\n", '\U0001F389')

	// Common gotchas
	fmt.Println("\n--- Common gotchas ---")
	emoji := "😀"
	fmt.Printf("Emoji string: %s\n", emoji)
	fmt.Printf("Byte length: %d (emoji is 4 bytes in UTF-8)\n", len(emoji))
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(emoji))
	// Accessing by index gives bytes, not runes
	fmt.Printf("emoji[0] (first byte): %d (not the full rune!)\n", emoji[0])

	// Converting between strings and rune slices
	fmt.Println("\n--- Converting between strings and rune slices ---")
	original := "Go"
	runeSlice := []rune(original)
	runeSlice[0] = 'C'            // Modify rune slice
	modified := string(runeSlice) // Convert back to string
	fmt.Printf("Original: %s, Modified: %s\n", original, modified)
}
