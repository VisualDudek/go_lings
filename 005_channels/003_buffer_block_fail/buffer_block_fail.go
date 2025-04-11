package main

func main() {
	ch := make(chan string, 2)
	ch <- "A"
	ch <- "B"
	ch <- "C" // blocks here! if no receiver
}
