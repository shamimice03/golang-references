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


### Ref:
- https://www.youtube.com/watch?v=sTFJtxJXkaY&t=10s
- https://www.golang-book.com/books/intro/8
