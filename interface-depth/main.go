package main

import (
	"errors"
	"fmt"
)

type Blog struct {
	Post string `json:"post"`
}

var EOF = errors.New("EOF")

var pos = 0

// Implementations must not retain p.
type Reader interface {
	Read(p []byte) (n int, err error)
}

func ReadAll(r Reader) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {

		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == EOF {
				err = nil
			}
			return b, err
		}
	}
}

func (b *Blog) Read(p []byte) (int, error) {
	if pos+1 <= len(b.Post) {
		n := copy(p, []byte(b.Post[pos:pos+1]))
		pos++
		return n, nil
	}

	return 0, EOF
}

func main() {

	b := &Blog{
		Post: "Tod",
	}

	output, err := ReadAll(b)

	fmt.Println(string(output))
	fmt.Println(err)

}
