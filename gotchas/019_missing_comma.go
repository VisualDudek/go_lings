//go:build ignore
// package gotchas

// missing comma in multi-line Literals
// add comment

func MissingCommaFail() {

	x := []int{
		1,
		2  // error
	}

	_ = x
}


func MissingCommaWorks() {

	x := []int{
		1,
		2, // ok
	}

	_ = x
}