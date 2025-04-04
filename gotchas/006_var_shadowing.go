package gotchas

// NO LONGER included as of Go .12+ USE staticcheck.dev

import "fmt"

func var_shadowing() {
	x := 1
	fmt.Println(x) // prints 1
	{
		fmt.Println(x) // prints 1
		x := 2
		fmt.Println(x) // prints 2
	}
	fmt.Println(x) // prints 1 (bad if you need 2)
}
