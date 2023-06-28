package main

import (
	"fmt"
	"strings"
)

func main() {

	a := []string{}
	a = append(a, "One", "two", "three", "four")

	b := strings.Join(a, "-")

	fmt.Println(b)
	// output: One-two-three-four

}


