// implemnt MyInt with custom printing
package main

import (
	"fmt"
	"reflect"
)

type MyInt int

func (i MyInt) String() string {
	return fmt.Sprintf("Custom printing %d", i)
}

func main() {
	var num MyInt = 6

	fmt.Println(num)
	fmt.Printf("Addition still works, add 4 -> %d\n", num+4)

	// ciekawe jakim typem jest num + 4 ? int czy MyInt ?
	fmt.Println(reflect.TypeOf(num + 4))
}
