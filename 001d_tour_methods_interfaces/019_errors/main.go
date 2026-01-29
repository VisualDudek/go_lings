// Package main demonstrates custom error types.
//
// What you can learn:
// - The error interface requires a single Error() method that returns a string.
// - Custom error types are created by implementing the Error() method.
// - Error types can hold additional context and metadata about failures.
// - Custom errors provide type-specific information while satisfying the error interface.
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// TAKE: best practice is to implement the Error() method on pointer receivers not values
// bc. it avoids copying and when you use value receiver it will never be nil (!!!)
func (e *MyError) Error() string {
	return e.What + " at " + e.When.String()
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
	}
}
