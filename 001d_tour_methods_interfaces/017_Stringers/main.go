package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// TAKE: Stringer is a type that can describe itself as a string
/* defined by the fmt package

type Stringer interface {
	String() string
}

*/

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
