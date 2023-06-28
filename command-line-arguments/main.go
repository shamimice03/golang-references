package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Args[0] --> The name of the command that invoked it
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	// for loop iterates over the range values
	for index, key := range os.Args[1:] {
		fmt.Println(index, " ", key)
	}

}


//Input-->  go run main.go test1 test2

//Outputs:

// /tmp/go-build3363658004/b001/exe/main
// test1
// test2
// 0   test1
// 1   test2