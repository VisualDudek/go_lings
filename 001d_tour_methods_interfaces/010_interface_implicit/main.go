package main

type I interface {
	M()
}

type T struct {
	S string
}

// TAKE: T implements the interface I, implicitly
func (t T) M() {
	println(t.S)
}

func main() {
	// TAKE: NOTE below syntax: no explicit declaration of intent to implement interface I
	var i I = T{"hello"}
	i.M()
}
