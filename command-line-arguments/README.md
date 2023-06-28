```go
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

```
