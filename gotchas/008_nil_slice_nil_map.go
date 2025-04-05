package gotchas

func nilSlice() {
	var s []int // nil slice

	s = append(s, 1) // it is on to add items to "nil" slice
}

func nilMap() {
	var m map[string]int
	m["one"] = 1 // error
}

// works: m := make(map[string]int)
