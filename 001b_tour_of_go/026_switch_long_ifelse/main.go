// Learn: switch without a tag, each case is a boolean expression.
package main

import "time"

func main() {
	t := time.Now()

	// TAKE: switch with no condition, each case is a boolean expression
	switch {
	case t.Hour() < 12:
		println("Good morning!")
	case t.Hour() < 17:
		println("Good afternoon.")
	default:
		println("Good evening.")
	}
}
