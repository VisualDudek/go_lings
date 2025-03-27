package gotchas

func redeclaringVariables() {
	one := 0
	one := 1 //error

	one, two := 1, 2 // ok

	one, two = two, one

}
