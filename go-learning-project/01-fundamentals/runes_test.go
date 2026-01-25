package main

import (
	"testing"
	"unicode/utf8"
)

// Test basic rune functionality
func TestRuneBasics(t *testing.T) {
	var r rune = 'A'
	if r != 65 {
		t.Errorf("Rune 'A' should be 65, got %d", r)
	}
}

// Test string vs rune length
func TestStringVsRuneLength(t *testing.T) {
	tests := []struct {
		str     string
		byteLen int
		runeLen int
	}{
		{"Hello", 5, 5},
		{"Hi🎉", 6, 3},
		{"", 0, 0},
	}

	for _, tt := range tests {
		if len(tt.str) != tt.byteLen {
			t.Errorf("%q: byte length = %d, want %d", tt.str, len(tt.str), tt.byteLen)
		}
		if utf8.RuneCountInString(tt.str) != tt.runeLen {
			t.Errorf("%q: rune count = %d, want %d", tt.str, utf8.RuneCountInString(tt.str), tt.runeLen)
		}
	}
}

// Test CountRunes
func TestCountRunes(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello", 5},
		{"Hi🎉", 3},
		{"", 0},
	}

	for _, tt := range tests {
		result := CountRunes(tt.input)
		if result != tt.expected {
			t.Errorf("CountRunes(%q) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

// Test ReverseRunes
func TestReverseRunes(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "olleH"},
		{"Hi🎉", "🎉iH"},
		{"a", "a"},
		{"", ""},
	}

	for _, tt := range tests {
		result := ReverseRunes(tt.input)
		if result != tt.expected {
			t.Errorf("ReverseRunes(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

// Test ToUpperRune
func TestToUpperRune(t *testing.T) {
	tests := []struct {
		input    rune
		expected rune
	}{
		{'a', 'A'},
		{'z', 'Z'},
		{'A', 'A'},
		{'5', '5'},
	}

	for _, tt := range tests {
		result := ToUpperRune(tt.input)
		if result != tt.expected {
			t.Errorf("ToUpperRune(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

// Test TakeFirst
func TestTakeFirst(t *testing.T) {
	tests := []struct {
		input    string
		n        int
		expected string
	}{
		{"Hello", 3, "Hel"},
		{"Hi🎉", 2, "Hi"},
		{"", 5, ""},
	}

	for _, tt := range tests {
		result := TakeFirst(tt.input, tt.n)
		if result != tt.expected {
			t.Errorf("TakeFirst(%q, %d) = %q, want %q", tt.input, tt.n, result, tt.expected)
		}
	}
}

// Test DropFirst
func TestDropFirst(t *testing.T) {
	tests := []struct {
		input    string
		n        int
		expected string
	}{
		{"Hello", 3, "lo"},
		{"Hi🎉", 2, "🎉"},
		{"", 5, ""},
	}

	for _, tt := range tests {
		result := DropFirst(tt.input, tt.n)
		if result != tt.expected {
			t.Errorf("DropFirst(%q, %d) = %q, want %q", tt.input, tt.n, result, tt.expected)
		}
	}
}

// Test RemoveSpaces
func TestRemoveSpaces(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "HelloWorld"},
		{"a b c", "abc"},
		{"", ""},
	}

	for _, tt := range tests {
		result := RemoveSpaces(tt.input)
		if result != tt.expected {
			t.Errorf("RemoveSpaces(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

// Test KeepLetters
func TestKeepLetters(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"H3ll0", "Hll"},
		{"Hello, World!", "HelloWorld"},
		{"", ""},
	}

	for _, tt := range tests {
		result := KeepLetters(tt.input)
		if result != tt.expected {
			t.Errorf("KeepLetters(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

// Test AddSpaces
func TestAddSpaces(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hi", "H i"},
		{"Hi🎉", "H i 🎉"},
		{"a", "a"},
		{"", ""},
	}

	for _, tt := range tests {
		result := AddSpaces(tt.input)
		if result != tt.expected {
			t.Errorf("AddSpaces(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}
