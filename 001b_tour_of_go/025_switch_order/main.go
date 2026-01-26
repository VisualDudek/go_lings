// Learn: switching on a constant Weekday, comparing computed expressions, and that Weekday supports int arithmetic.
package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()

	fmt.Printf("type=%T, value=%v\n", today, today)

	// TAKE: twisted switch statement with value first, then expressions to match
	switch time.Saturday {
	case today:
		fmt.Println("Today.")
		// TAKE: how can you add int to Weekday type?
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

}
