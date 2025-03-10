package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader // chyba nie do końca rozumiem czym jest `io.Reader`
}

// to nie jest łątwy przykład
// 1. w pierwszy podejściu nie odczytałem danych z poprzedniej warstwy
func (r rot13Reader) Read(b []byte) (int, error) {

	// Read data from the underlying io.Reader
	n, err := r.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := range b {
		switch { // mój autorski if-else
		case b[i] >= 'A' && b[i] <= 'Z':
			b[i] = 'A' + (b[i]-'A'+13)%26

		case b[i] >= 'a' && b[i] <= 'z':
			b[i] = 'a' + (b[i]-'a'+13)%26
		}
	}

	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // zakładam że `io.Copy` wywołuje Reader interface ?
}
