package gotchas

// MUST use var, dwa wyjątki (1) global var (2) fn. args
import "fmt"

var gvar int // you can have unused global vars

func unused_vars() {
	var one int   // error, unused variable
	two := 2      // error, unused variable
	var three int // error, even though it's assigned 3 on the next line
	three = 3     // you MUST use var 'three' to make compiler happy

	func(unused string) { // lambda ???
		fmt.Println("Unused arg. No compile error")
	}("what?") // zapamiętaj taką możliwość wywołania fn.
}
