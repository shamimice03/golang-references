package main

import (
	"fmt"
	"lib" // package import from moduleB
)

func main() {

	fmt.Println("Hello From Module A")
	lib.HelloFromModuleB() //calling function from moduleB
}
