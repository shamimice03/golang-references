### Resources:

-------- Golang --------
--------------------------

golang tutorial (advanced) => https://www.youtube.com/playlist?list=PLujhHB_uAFJws6Vv5q1KDoaQ4YcpS9UOm

- composition in golang (inheritance alternative) => https://www.youtube.com/watch?v=I0Ei70cZtf8&t=122s

- nil pointer dereference => https://ashourics.medium.com/nil-pointer-deference-in-golang-91aaa245145c

- Connect go lang with PostgreSQL => https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/

- using blank identifier(underscore _ ) in import => https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/

- go lang sql drivers list => https://github.com/golang/go/wiki/SQLDrivers

- gorm in k-fis project => https://v1.gorm.io/docs/connecting_to_the_database.html

- MVC pattern golang project structure tutorial (in Japanese) => https://dev.classmethod.jp/articles/go-sample-rest-api/

- golang crud in MVC pattern folder architecture (using gorm.io V2) => https://codewithmukesh.com/blog/implementing-crud-in-golang-rest-api/

- go clean architecture => https://github.com/bxcodec/go-clean-arch

- go clean architecture article => https://medium.easyread.co/golang-clean-archithecture-efd6d7c43047

- golang clean architecture folder structure followed in k-fis => https://medium.easyread.co/trying-clean-architecture-on-golang-2-44d615bf8fdf

- golang swagger implementation = https://github.com/swaggo/http-swagger
  
- swaggo http-swagger comment annotation description => https://ldej.nl/post/generating-swagger-docs-from-go/
  
- gorm model tags(and auto table create)- https://gorm.io/docs/models.html

- golang tutorial japaneese - https://www.youtube.com/watch?v=kPXfMFJ0oIE

- demo project(vuetify, go, postgres) for db connect and crud, clean architecture, swagger, query param =>
- check project-> "vue go postgres crud" => https://github.com/IshmamAbir/CRUD-Go-Postgres-Vuetify3

- multi-modular architecture - < (completed. will be adding here after completing the documentation) >

- Building a modern application with a Golang API backend + a Vue.js SPA frontend using Docker ->
https://juliensalinas.com/en/golang-API-backend-vuejs-SPA-frontend-docker-modern-application/

- Struct, method and interface definition = https://qiita.com/S-Masakatsu/items/6fb8e765cd443e2edd7f

- db diagram generate tools (online) = https://dbdiagram.io/home

- jwt authentication & authorization - https://www.bacancytechnology.com/blog/golang-jwt
- https://www.youtube.com/watch?v=hWmR8YtlFlE&t=307s
- https://github.com/architanayak/golang-jwt-authentication

- Udemy course for reference golang- https://www.udemy.com/course/building-web-applications-with-go-intermediate-level/

Udemy course for golang & vuejs3 - https://www.udemy.com/course/working-with-vue-3-and-go/

blogs on golang + vueJS/vuetify - https://adeshg7.medium.com/vuejs-golang-a-rare-combination-53538b6fb918

dynamic query https://betterprogramming.pub/dynamic-sql-query-with-go-8aeedaa02907

others https://stackoverflow.com/questions/66563805/how-to-insert-a-null-foreign-key-in-gorm



golang resources begineers er jnno.
amr starting e jegula follow korsilm



### About the repository
- Every directory is a separate module
- When creating a directory use <go mod init> to initialized the module.
```
go mod init github.com/shamimice03/golang-references/<DIRECTORY>
go mod tidy
```

### To-do


### Doc & Refs
- https://go.dev/tour/welcome/1
- https://www.youtube.com/watch?v=50ewcV8PsI4
- https://www.youtube.com/watch?v=TkbhQQS3m_o&t=500s
- Pointers
  - https://www.youtube.com/watch?v=sTFJtxJXkaY&t=8s&ab_channel=JunminLee
- Working With Json
  -  https://www.youtube.com/watch?v=_ok29xwZ11k&list=PLHq1uqvAteVufR04uKxK-oDRZqFenJLii&index=3

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




