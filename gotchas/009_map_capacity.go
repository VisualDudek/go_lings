package gotchas

// nie możesz użyć cap() on maps. Implementacja hash map-y powoduje że cap nie ma sensu

func mapCap() {
	m := make(map[string]int, 99)
	cap(m) // error
}
