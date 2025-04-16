package main

import "fmt"

type data struct {
	num   int
	key   *string
	items map[string]bool
}

func (d *data) pmethod() {
	d.num = 7
}

func (d data) vmethod() {
	d.num = 8 // ooo nawet static type checker? pokazuje że
	// to sie nie zadzieje :)
	*d.key = "v.key"
	d.items["vmethod"] = true
}

func main() {
	key := "key.1"
	d := data{1, &key, make(map[string]bool)}

	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)

	d.pmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)

	d.vmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
}
