package gotchas

import "fmt"

// The most reliable way to know if a given map record exists
// is to check the second value returned by the map access operation

func nonExistingKeyMap() {
	d := map[string]string{"one": "a", "two": "", "three": "c"}

	if _, ok := d["two"]; !ok {
		fmt.Println("no entry")
	}
}
