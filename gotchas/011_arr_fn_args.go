package gotchas

import "fmt"

// When you pass arrays to functions the functions
//reference the same memory location, so they can update the original data.
//Arrays in Go are values, so when
//you pass arrays to functions the functions get a copy of the original array data.
//This can be a problem if you are trying to update the array data.

func arrayFunctionArgs001() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) // prints: [7, 2, 3]
	}(x)

	fmt.Println(x) // prints: [1, 2, 3] not ok if you need [7, 2, 3]
}

// if you need to update the original array use array pointer types.
func arrayFunctionArgs002() {
	x := [3]int{1, 2, 3}

	func(arr *[3]int) {
		(*arr)[0] = 7    // uwaga na składnie, nawiady do dereferencowania
		fmt.Println(arr) // prints: &[7, 2, 3]
	}(&x)

	fmt.Println(x) // prints: [7, 2, 3]  updated values
}

// Another option is to use slices.
// Even though your function gets a copy of the slice variable
// it still references the original data.
func arrayFunctionArgs003() {
	x := []int{1, 2, 3}

	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr) // prints: [7, 2, 3]
	}(x)

	fmt.Println(x) // prints: [7, 2, 3]  updated values
}
