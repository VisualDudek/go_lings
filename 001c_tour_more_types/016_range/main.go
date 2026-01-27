package main

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		println("2 **", i, "=", v)
	}
}
