// This code demonstrates Go's switch statement with short variable declaration.
// Learn: switch syntax with init statement (switch os := ...), case matching, default clause, no fallthrough between cases.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on:")
	// TAKE: no fallthrough between cases by default
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
