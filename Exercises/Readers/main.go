package main

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

//import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (MyReader) Read(b []byte) (int, error) { // ok ale można lepie
	b[0] = 'A'
	return 1, nil
}

func (MyReader) Read(b []byte) (int, error) {
	for i := range b { // TAK JEST LEPIEJ
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	// reader.Validate(MyReader{}) NIE ZADZIAŁA BEZ import, co to jest ?
}
