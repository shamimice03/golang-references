package main

import (
	"fmt"
	"io"
)

type Blog struct {
	Post string `json:"post"`
}

func (b *Blog) Read(p []byte) (int, error) {
	//fmt.Println(p)
	n := copy(p, []byte(b.Post))
	fmt.Println(p)
	fmt.Println(n)
	return n, io.EOF
}

func main() {

	b := &Blog{
		Post: "Today is wrere disagdugasdg asdyayd dysaydfad ydysafdyfad ydfsaydfsaydf",
	}

	output, err := io.ReadAll(b)

	fmt.Println(output)
	fmt.Println(err)

}
