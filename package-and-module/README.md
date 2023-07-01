### Project Structure
```
├── README.md
├── breaking
│   └── breaking.go
├── go.mod
├── main.go
└── news.go

```

### Create a Module (`go.mod`)
```
go mod init news.com/events
```
- `go mod init <name>`: Typically `github` repo name is used. (`github.com/goapp`)
- Files in the same directory must be under the same package, such as: `main.go` and `news.go` in the same package (main package)

##### `main.go`
```go
package main

import "news.com/events/breaking"

func main() {
	todaysNews()  // function calling from same package (news.go)
}
```
##### `news.go`
```go
package main

import "fmt"

func todaysNews() {
	fmt.Println("Today is a great day")
}

```

### Create Another Package
- Create a new directory `breaking` 
- Create a new file `breaking/breaking.go` (matching filename and directory is a good practice)

##### `breaking.go`
```go
package breaking // new package name

import "fmt"

func BreakingNews() {   // function name(or others) must start with capital to export that in another package
	fmt.Println("Breaking News")
}
```

### Use the new package(breaking) inside the `main` package
##### `main.go`
```go
package main

import "news.com/events/breaking" // `news.com/events/` is the module name. Its a virtual path, created while we initilized `go mod init` command.

func main() {
	todaysNews()            // function calling from same package (news.go)
	breaking.BreakingNews() // function calling from different package (breaking)
}
```