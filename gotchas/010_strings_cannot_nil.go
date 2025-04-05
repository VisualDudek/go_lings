package gotchas

func stringsCannotBeNilFails() {
	var x string = nil // error

	if x == nil { // error
		x = "default"
	}
}

func stringsCannotBeNilWorks() {
	var x string // defaults to "" (zero value)

	if x == "" {
		x = "default"
	}
}
