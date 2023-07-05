### Pointer
#### Without - Pointer
```go
func zero(x int) {
	x = 0
}
func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
}
```
#### Explanation:
```
x = 5 
Sending the copy of x to the zero func
Original value remains same inside the main func
```
#### With - Pointer
```go
package main

import "fmt"

func zero(p *int) {

	fmt.Println("Before change")
	fmt.Println(p)
	fmt.Println(*p)

	*p = 0

	fmt.Println("After change")
	fmt.Println(p)
	fmt.Println(*p)

}

func main() {
	x := 5
	fmt.Println("value of x is:", x)
	zero(&x) //passing address of x -> &x

	fmt.Println("value of x changed to:", x)
}
```

```go
Output:
value of x is: 5
Before change
0xc000016098
5
After change
0xc000016098
0
value of x changed to: 0
```

#### Explanation:
```
x = 5 // current value of x = 5
&x //address of x
zero(&x) //passing address of x to zero(x) function

zero(p *int) // p = currently holding address of x
*p // dereferencing, currently holding value of x
*p = 0 // change the value of x to zero
```
#### Another Example:
```go
package main
import "fmt"
/* For exercises uncomment the imports below */
// import "strconv"
// import "encoding/json"

func main() {

    // a normal variable  whose address the pointer will store 
    var intData = 20 
    
    //declaration of a pointer 
    var intPointer *int

    //intPointer now points towards intData    
    intPointer = &intData 

    fmt.Println("what intData stores:", intData)
    fmt.Println("address of intData:", &intData)
    fmt.Println("what intPointer stores:", intPointer)


    //this updates the value of intData using dereferncing operator 
    *intPointer = 30

    fmt.Println("what intData now stores:", intData)
}

// Output:
// what intData stores: 20
// address of intData: 0xc00010c000
// what intPointer stores: 0xc00010c000
// what intData now stores: 30
```

### Ref:
- https://www.youtube.com/watch?v=sTFJtxJXkaY
- https://www.golang-book.com/books/intro/8
- https://www.educative.io/answers/what-are-pointers-in-golang