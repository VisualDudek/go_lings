// Learn: wrapping an io.Reader with a custom reader type;
// implementing io.Reader interface with transformation logic;
// using io.Copy.
package main

import (
	"io"
	"os"
	"strings"
)

// TAKE: jeśli struct rot13Reader ma pole r typu io.Reader
// to musisz je baranie przeczytać metodą Read()
type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rr.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
		case 'A' <= b[i] && b[i] <= 'Z':
			b[i] = 'A' + (b[i]-'A'+13)%26
		case 'a' <= b[i] && b[i] <= 'z':
			b[i] = 'a' + (b[i]-'a'+13)%26
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
