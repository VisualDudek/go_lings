package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{1, 2, 3}
	fmt.Printf("Slice s1 Length: %d, and Capacity: %d\n", len(s1), cap(s1))
	fmt.Printf("s1: Underlying array address: %p\n", unsafe.Pointer(&s1[0]))

	// Function Signature:
	// func SliceData(slice []T) *T
	// return A pointer to the first element of the slice's underlying array
	ptr := unsafe.SliceData(s1)
	fmt.Printf("s1: Underlying array address using unsafe.SliceData: %p\n\n", ptr)

	s2 := append(s1, 4) // wydaje Ci się że robisz append na s1 a tak naprawdę dostajesz
	// inną wiekszą tablice pod spodem
	// zazwyczaj nie zaboli bo pewnie będziesz przypisaywał do s1 a nie tak
	// jak tutaj do s2 ALE można sobie strzelić w stope
	fmt.Printf("Slice s2 Length: %d, and Capacity: %d\n", len(s2), cap(s2))
	fmt.Printf("s2: Underlying array address: %p\n", unsafe.Pointer(&s2[0]))

	s2[0] = 100

	fmt.Println(s1)
	fmt.Println(s2)

	// Albo w drugą stronę celowo chcesz "rozmnożyć" slice dlatego dajesz s2 ALE s1
	// miało capacity na append i okazuje się że zmieniłeś wartości w s1 i s2 a
	// chciałeś tylko w s2
}
