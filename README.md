### To-do
- https://www.youtube.com/watch?v=_ok29xwZ11k&list=PLHq1uqvAteVufR04uKxK-oDRZqFenJLii&index=3


### Doc & Refs
- https://go.dev/tour/welcome/1
- https://www.youtube.com/watch?v=50ewcV8PsI4
- https://www.youtube.com/watch?v=TkbhQQS3m_o&t=500s
- Pointers
  - https://www.youtube.com/watch?v=sTFJtxJXkaY&t=8s&ab_channel=JunminLee

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
> go mod init  

> go mod tidy


Here's a breakdown of what each command does:

`go mod init <name>`: 
This command initializes a new Go module with the specified <name>.
A Go module is a collection of related Go packages that are versioned together. 
The <name> parameter typically follows a domain name format, such as github.com/username/projectname. 

For example, if you have a project called "myproject" hosted on GitHub under your username "johnsmith," you would run the command go mod init github.com/johnsmith/myproject to initialize the Go module.

`go mod tidy`: This command is used to update the dependencies listed in the go.mod file. It ensures that the go.mod file accurately reflects the actual dependencies used in your code. The go mod tidy command will add any missing dependencies to the go.mod file and remove any unnecessary dependencies that are not used in your code.


