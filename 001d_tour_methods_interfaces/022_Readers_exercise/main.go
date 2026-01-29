// Learn: implementing io.Reader interface with custom type;
// using copy() to fill byte slice in Read method.
package main

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	n = copy(b, "A")
	return
}

func main() {
	var r MyReader
	b := make([]byte, 8)

	n, err := r.Read(b)
	if err != nil {
		panic(err)
	}

	println(n)
	println(string(b[:]))
}
