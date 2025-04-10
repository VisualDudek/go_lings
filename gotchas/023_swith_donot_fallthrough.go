package gotchas

import "fmt"

func SwitchFallthroughFail() {
	isSpace := func(ch byte) bool {
		switch ch {
		case ' ': //error bc will NOT FALLTHROUGH
		case '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace('\t')) // prints true (ok)
	fmt.Println(isSpace(' '))  // prints false (not ok)

}

func SwitchFallthroughWorks() {
	isSpace := func(ch byte) bool {
		switch ch {
		case ' ', '\t':
			return true
		}
		return false
	}

	fmt.Println(isSpace('\t')) // prints true (ok)
	fmt.Println(isSpace(' '))  // prints false (not ok)
}
