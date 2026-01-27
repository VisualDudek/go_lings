package main

func main() {
	pow := make([]int, 10)
	for i := range pow {
		// TAKE: bit shift operator req. the right operand to be an uint
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		println(value)
	}
}
