package main

import (
	"strings"
	"unicode/utf8"
)

// Exercise 1: Basic Rune Manipulation - Implementations

// CountRunes returns the number of runes in a string (not bytes)
func CountRunes(s string) int {
	return len([]rune(s))
}

// ReverseRunes reverses a string by runes (handles multi-byte characters correctly)
func ReverseRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ToUpperRune converts a single ASCII rune to uppercase
func ToUpperRune(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

// Exercise 2: Unicode String Processing - Implementations

// TakeFirst returns the first n runes of a string
func TakeFirst(s string, n int) string {
	runes := []rune(s)
	if n > len(runes) {
		n = len(runes)
	}
	if n < 0 {
		n = 0
	}
	return string(runes[:n])
}

// DropFirst skips the first n runes and returns the rest
func DropFirst(s string, n int) string {
	runes := []rune(s)
	if n >= len(runes) {
		return ""
	}
	if n < 0 {
		n = 0
	}
	return string(runes[n:])
}

// EachRuneIndex returns byte indices where each rune starts
func EachRuneIndex(s string) []int {
	var indices []int
	byteIndex := 0
	for byteIndex < len(s) {
		indices = append(indices, byteIndex)
		r, size := utf8.DecodeRuneInString(s[byteIndex:])
		if r == utf8.RuneError {
			break
		}
		byteIndex += size
	}
	return indices
}

// Exercise 3: Rune Filtering - Implementations

// RemoveSpaces removes all whitespace runes
func RemoveSpaces(s string) string {
	return strings.Join(strings.Fields(s), "")
}

// KeepLetters keeps only ASCII letters
func KeepLetters(s string) string {
	var result []rune
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			result = append(result, r)
		}
	}
	return string(result)
}

// AddSpaces inserts a space between each rune
func AddSpaces(s string) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}
	var result strings.Builder
	for i, r := range runes {
		if i > 0 {
			result.WriteRune(' ')
		}
		result.WriteRune(r)
	}
	return result.String()
}
