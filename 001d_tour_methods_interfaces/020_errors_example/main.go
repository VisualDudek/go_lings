package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// TAKE: `fmt.Sprint(e)` detects that `e` implements the `error` interface
	// automatically calls `e.Error()` to get the string representation ... ending
	// in infinite recursion
	return "Cannot sqrt negative number:" + fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		// TAKE: need type conversion to convert x (float64) to ErrNegativeSqrt
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
