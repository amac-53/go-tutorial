package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
    n, err = rot.r.Read(b)
    for i := range b {
		switch {
			case 'A' <= b[i] && b[i] <= 'Z':
				b[i] = (b[i] - 'A' + 13) % 26 + 'A'
			case 'a' <= b[i] && b[i] <= 'z':
				b[i] = (b[i] - 'a' + 13) % 26 + 'a'
		}		
    }
    return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
