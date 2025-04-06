package gotchas

import "fmt"

// you can build milti-dimensional array/slices
// (1) slices of "independent" slices or
// (2) slices of "shared data" slices

// creating slices of "independent" slices is a two step process.
func slicesOfSlices() {
	x := 2
	y := 4

	table := make([][]int, x)
	for i := range table {
		table[i] = make([]int, y)
	}
}

// Creating a dynamic milti-dimensional array using slices of "shared data"
// slices is a three step process.

func slicesOfSharedData() {
	h, w := 2, 4

	// fist you have to create the data "container" slice that will hold raw data
	raw := make([]int, h*w)
	for i := range raw {
		raw[i] = i
	}

	fmt.Println(raw, &raw[4])
	// prints: [0 1 2 3 4 5 6 7] <ptr_addr_x>

	// create the outer slice
	table := make([][]int, h)

	// initialize each inner slice by reslicing the raw data slice
	for i := range table {
		table[i] = raw[i*w : i*w+w]
	}

	fmt.Println(table, &table[1][0])
	// prints: [[0 1 2 3] [4 5 6 7]] <ptr_addr_x>

}
