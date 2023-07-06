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

	// Way to use interface

	p := Adder(&n)
	k := Adder(&s)

	fmt.Printf("Type of p: %T\n", p)
	fmt.Printf("Type of k: %T\n", k)

	// Another way to user Interface
	fmt.Println("Another way to use interface:")
	// Interface type variable, assigned value a struct type
	var a Addition
	fmt.Printf("Type of a: %T\n", a)

	a = &n
	fmt.Printf("Type of a: %T\n", a)
	fmt.Println(a.add())
	// Dynamically assigned another struct type
	a = &s
	fmt.Printf("Type of a: %T\n", a)
	fmt.Println(a.add())

	// Type Assertion
	// If we want to access length() method through the interface

	// a.length()  -> error

	// type assertion:  interface_var.(Type).method
	// Interface var: a
	// Type: *StringContent
	// Method: Length

	fmt.Println(a.(*StringContent).length())

	// Type Assertion Check
	// New Interface type var with Number Struct assigned
	var demo Addition = &n

	dummy, ok := demo.(*Number)
	fmt.Printf("Check %v\nok: %v\n", dummy, ok)

	fmt.Println("Subtract:", dummy.subtract())

}
