package main

// Go automatically dereferences and takes addresses when calling methods.
// Methods with pointer receivers can be called on values (Go takes the address),
// and methods with value receivers can be called on pointers (Go dereferences).

import "fmt"

type Point struct {
	X, Y float64
}

func (p *Point) Scale(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{3, 4}
	fmt.Println(p)
	// TAKE: can call method with pointer receiver on value type, Go will take address automatically
	p.Scale(10)
	fmt.Println(p)

	v := &Point{6, 8}
	fmt.Println(v)
	// TAKE: also works when calling on pointer type
	v.Scale(5)
	fmt.Println(*v)
}
