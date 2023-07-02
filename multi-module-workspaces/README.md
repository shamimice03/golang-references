### Multi-Module Workspace in Go
- Using `go.work` file we can work with multiple module in go (from v1.18)

#### Project Structure
```
.
├── README.md
├── go.work
├── moduleA
│   ├── go.mod
│   └── main.go
└── moduleB
    ├── go.mod
    └── lib.go
```
Create Two folders named as `moduleA`  and `moduleB`
- Initialized them using `god mod init NAME`
   - moduleA - `go mod init multimodule`
   - moduleB - `go mod init lib`

#### `main.go`
```go
package main

import (
	"fmt"
	"lib"   // package import from moduleB
)

func main() {

	fmt.Println("Hello From Module A")
	lib.HelloFromModuleB() //calling function from moduleB
}

```

#### Creating multimodule workspace using - `go.work`
- Initialized `go.work`
```
go work init ./moduleA ./moduleB
```

```go
// go.work

go 1.20

use (
	./moduleA
	./moduleB
)

```
- Add a new module named `lib2` to the `go.work` file
```
go work use ./lib2
```

#### Run go module
```go
multi-module-workspaces $ go run multimodule
// Outputs
// Hello From Module A
// Hello From Module B
```

#### References:
- https://go.dev/doc/tutorial/workspaces
- https://www.youtube.com/watch?v=6MBSicd9Ipw&ab_channel=nummer4