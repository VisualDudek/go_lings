package gotchas

import "fmt"

type info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}

// fails:
func noShortDeclarationToSetFieldValues() {
	var data info

	data.result, err := work() // error
	fmt.Println("info: %+v\n", data)
}

// works:
func yesShortDeclarationToSetFieldValues() {
	var data info

	var err error
	data.result, err = work() // predeclare and use std assignment operator `=`
	fmt.Println("info: %+v\n", data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("info: %+v\n", data) // prints: info: {result:13}
}

// does short decl. works without Field Values ? YEP
func tryShortDeclarationToSetFieldValues() {

	data, err := work() // działa bez assignment do Field Value
	fmt.Println("info: %+v\n", data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("info: %+v\n", data) // prints: info: {result:13}
}
