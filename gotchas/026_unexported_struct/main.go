package main

import (
	"encoding/json"
	"fmt"
)

// Unexported Struct Fields are not encoded

type MyData struct {
	One int
	two string
}

func main() {
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in)

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) // two is now visible

	var out MyData // two was not exported so when you decode
	// you will end up with zero values
	json.Unmarshal(encoded, &out)

	fmt.Printf("%#v\n", out)

}
