package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (this *rot13Reader) Read(o []byte) (int, error) {
	b := make([]byte, 100)
	len, err := this.r.Read(b)

	for i:=0;i<len;i++ {
		if ('A' <= b[i] && b[i] <= 'M') || ('a' <= b[i] && b[i] <= 'm') {
			o[i] = b[i] + 13
		} else if ('N' <= b[i] && b[i] <= 'Z') || ('n' <= b[i] && b[i] <= 'z') {
			o[i] = b[i] - 13
		} else {
			o[i] = b[i]
		}
	}
	return len, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

