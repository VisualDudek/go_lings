package gotchas

// the "nil" identifier can be use as the "zero value" ONLY for interfaces,
//functions, pointers, maps, slices and channels.

func noNilWithoutTypeFail() {
	var x = nil // error

	_ = x
}

func noNilWithoutTypeWorks() {
	var x interface{} = nil

	_ = x
}
