package main

func main() {
	a := make([]int, 5)

	a = append(a, 1)
	// co miałeś na myśli a co się stało ?
	// jakie by default zachowanie Go spowodowało "błąd" ?
	// jak to naprawić ?
	_ = a

}
