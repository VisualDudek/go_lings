package gotchas

import "fmt"

// do not expext the items to be in a certain order

func IterateForMap() {
	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m {
		fmt.Println(k, v)

	}
}
