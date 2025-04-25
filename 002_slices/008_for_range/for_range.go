package main

var s = []int{1, 2, 3, 4, 5}

// ważne jest żeby pamiętać że value jest copią a nie referencją !!!

func main() {
	_ = s
	// iteruj po idx i v

	// iteruj tylko po idx (może powodować błędy w logice jak pomylisz z v)

	// iteruj tylko po value

}
