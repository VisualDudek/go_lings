// Structs: defining struct types with fields, creating instances with positional values, and printing struct values.
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Printf("%#v\n", v)
}
