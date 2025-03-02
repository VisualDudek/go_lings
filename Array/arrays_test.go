package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArrays(t *testing.T) {
	// Test array of 10
	var nums [10]int

	if len(nums) != 10 {
		t.Errorf("Expected array lenght 10, got %d", len(nums))
	}

	if reflect.TypeOf(nums) != reflect.TypeOf([10]int{}) {
		t.Errorf("Expected array of type [10]int, got %T", nums)
	}

	if reflect.TypeOf(nums).Kind() != reflect.Array || reflect.TypeOf(nums).Len() != 10 {
		for _, v := range nums {
			if reflect.TypeOf(v).Kind() != reflect.Int {
				t.Errorf("Expected element of type int, got %T", v)
			}
		}
		t.Errorf("Expected array of type [10]int, got %T", nums)
	}

	// Test array of 4 string
	abc := [4]string{"A", "B", "C", "D"}

	if abc[2] != "C" && len(abc) != 4 {
		t.Errorf("Expected array lenght 4, got %d and 3rd element 'C', got %s", len(abc), abc[2])
	}

	// Print abc array itself
	fmt.Printf("Array abc: %v\n", abc)

	a := [2]int{1, 2}
	b := [2]int{3, 4}
	// Concatenate two arrays
	combined := append(a[:], b[:]...)

	fmt.Printf("Combined array: %v\n", combined)

	if !reflect.DeepEqual(combined, []int{1, 2, 3, 4}) { // fix me
		t.Errorf("Expected combined array [1, 2, 3, 4], got %v", combined)
	}

}
