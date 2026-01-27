// Maps: declaring map types, zero value nil maps, initializing with make, and adding/reading key-value pairs.
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// TAKE: zero value of map is nil; need to use make to initialize before adding key-value pairs.
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
