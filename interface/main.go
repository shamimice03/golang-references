package main

import (
	"fmt"
	"math"
)

type Addition interface {
	add() interface{}
}

type Number struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type StringContent struct {
	Str1 string `json:"str1"`
	Str2 string `json:"str2"`
}

func (n *Number) add() interface{} {
	return n.Num1 + n.Num2
}

func (n *Number) subtract() int {
	return int(math.Abs(float64(n.Num1) - float64(n.Num2)))
}

func (s *StringContent) add() interface{} {
	return s.Str1 + s.Str2
}

func (s *StringContent) length() int {

	// strings.ToUpper(s.Str1)
	// strings.ToLower(s.Str2)
	return len(s.Str1 + s.Str2)
}

func Adder(a Addition) interface{} {
	fmt.Println("Adding: ", a.add())
	return a.add()
}

func main() {
	n := Number{
		Num1: 2,
		Num2: 3,
	}
	s := StringContent{
		Str1: "Hello ",
		Str2: "World",
	}

	// Without Interface

	// fmt.Println("Adding two number", n.add())
	// fmt.Println("Subtracting two number", n.subtract())
	// fmt.Println("Concate two string", s.add())
	// fmt.Println("Length after concate two string", s.length())

	p := Adder(&n)
	k := Adder(&s)

	fmt.Printf("Type of p: %T\n", p)
	fmt.Printf("Type of k: %T\n", k)
}
