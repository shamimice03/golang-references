package main

import "fmt"

func zero(p *int) {

	fmt.Println("Before change")
	fmt.Println(p)  // address of x
	fmt.Println(*p) // value of x

	*p = 0 // change value of x to 0

	fmt.Println("After change")
	fmt.Println(p)  // same address as x
	fmt.Println(*p) // value of x changed to 0

}

func main() {
	x := 5
	fmt.Println("value of x is:", x)
	zero(&x) //passing address of x -> &x

	fmt.Println("value of x changed to:", x)
}
