package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Rot13Reader struct {
	r io.Reader
}

func rot13Parser(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return 'A' + (b-'A'+13)%26
	}

	if b >= 'a' && b <= 'z' {
		return 'a' + (b-'a'+13)%26
	}

	return b
}

func (rot13 Rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)

	fmt.Printf("n = %v err = %v\n", n, err)
	fmt.Printf("b[:n] = %q\n", b[:n])

	for i := range b[:n] {
		parsedData := rot13Parser(b[i])

		fmt.Printf("b[%d]: %q -> %q\n", i, b[i], parsedData)

		b[i] = parsedData
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")

	r := Rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
