package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{1, 2, 3}
	fmt.Printf("Slice s1 Length: %d, and Capacity: %d\n", len(s1), cap(s1))
	fmt.Printf("s1: Underlying array address: %p\n\n", unsafe.Pointer(&s1[0]))

	s2 := append(s1, 4) // wydaje Ci się że robisz append na s1 a tak naprawdę dostajesz
	// inną wiekszą tablice pod spodem
	// zazwyczaj nie zaboli bo pewnie będziesz przypisaywał do s1 a nie tak
	// jak tutaj do s2 ALE można sobie strzelić w stope
	fmt.Printf("Slice s2 Length: %d, and Capacity: %d\n", len(s2), cap(s2))
	fmt.Printf("s2: Underlying array address: %p\n", unsafe.Pointer(&s2[0]))

	s2[0] = 100

	fmt.Println(s1)
	fmt.Println(s2)
}
