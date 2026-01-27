// Pointers to structs: accessing struct fields through pointers with implicit dereferencing; Go automatically handles it without (*p).X syntax.
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// TAKE: no need to dereference the pointer to access struct fields; Go automatically handles it.
	p.X = 1e9 // equivalent to (*p).X = 1e9
	fmt.Println(v)
}
