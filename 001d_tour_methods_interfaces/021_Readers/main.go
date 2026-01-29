// Learn: using io.Reader interface;
// reading from strings.NewReader into a byte slice; detecting EOF.
package main

import (
	"fmt"
	"io"
	"strings"
)

// TAKE: NOTE: io.Reader interface has a Read method with pointer receiver
// func (T) Read(b []byte) (n int, err error)

func main() {
	// TAKE: using strings.NewReader to create an io.Reader
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		// TAKE: Read implements the io.Reader interface using pointer receiver
		// TAKE: check definition of Read method
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		// TAKE: check definition of io.EOF
		if err == io.EOF {
			break
		}
	}
}
