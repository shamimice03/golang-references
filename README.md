### Doc & Refs
- https://go.dev/tour/welcome/1
- https://www.youtube.com/watch?v=50ewcV8PsI4
- https://www.youtube.com/watch?v=TkbhQQS3m_o&t=500s

### Check version
```
> go version
```

### Build The Project
```
> ls
main.go

> go build main.go

> ls
main main.go

## execute file
> ./main
```

### Go commands
> go mod init NAME 

> go mod tidy


Here's a breakdown of what each command does:

`go mod init <name>`: 
This command initializes a new Go module with the specified <name>.
A Go module is a collection of related Go packages that are versioned together. 
The <name> parameter typically follows a domain name format, such as github.com/username/projectname. 
The go mod init command creates a new go.mod file in the current directory, which serves as the module's definition file. 
It records the module's name and its dependencies.

For example, if you have a project called "myproject" hosted on GitHub under your username "johnsmith," you would run the command go mod init github.com/johnsmith/myproject to initialize the Go module.

`go mod tidy`: This command is used to update the dependencies listed in the go.mod file. It ensures that the go.mod file accurately reflects the actual dependencies used in your code. The go mod tidy command will add any missing dependencies to the go.mod file and remove any unnecessary dependencies that are not used in your code.

Running go mod tidy also verifies that the module's dependencies are consistent and compatible with each other. It checks that the requested versions of the dependencies can be resolved and updates the go.sum file, which contains the expected cryptographic checksums for the dependencies.

It's a good practice to run go mod tidy periodically or whenever you make changes to your code or dependencies to keep your go.mod file up to date and maintain a clean dependency graph.

Both of these commands are essential for working with Go modules effectively. go mod init sets up the initial module definition, while go mod tidy ensures that the module's dependencies are accurately represented and properly managed.

