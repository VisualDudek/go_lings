package main

import (
	"fmt"
	"slices" // Go 1.21+
)

func main() {
	nums := []int{1, 2, 5, 6}
	fmt.Println("Before:", nums) // → [1 2 5 6]

	// Insert single element (3) at index 2
	nums = slices.Insert(nums, 2, 3)
	fmt.Println("Insert 3 at pos 2:", nums) // → [1 2 3 5 6]

	// Insert multiple elements (7,8,9) at the end
	nums = slices.Insert(nums, len(nums), 7, 8, 9)
	fmt.Println("Append 7,8,9:", nums) // → [1 2 3 5 6 7 8 9]

	// Insert at the very start
	nums = slices.Insert(nums, 0, 0)
	fmt.Println("Prepend 0:", nums) // → [0 1 2 3 5 6 7 8 9]
}
