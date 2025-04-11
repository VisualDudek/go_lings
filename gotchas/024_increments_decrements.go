package gotchas

import "fmt"

// go does not support the PREFIX version of the inc. and dec. operations
// ++i this is error

// ALSO cannot use it in expresions
// data[i++] error

func increment() {
	data := []int{1, 2, 3}
	i := 0
	i++
	fmt.Println(data[i])
}
