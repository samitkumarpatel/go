# 00-variable
[https://play.golang.org](https://play.golang.org) - make use of playground env.


### Variable Declaration

```go
package main

import ( "fmt" ) 

//variable declaration and assignment
var k float32 = 42

// group of variables and assignment
var (
    actorName string = ""
    companion string = ""
    doctorNumber int = 3
    season int = 11
)

var I int = 100 // Globally visible

func main() {
    var i int // declare a variable but not initializing to it
    i = 42 //assign
    i = 27 // update previous value of i

    var i int = 42 // can replace like this from above
    i := 42 // shortend way to declare and assign a variable

    var j float32 = 27
    fmt.Printf("%v, %T",i,j) // %v print value and %T print type

    // group of variable
    var (
        eId int = 1001
        name string = "x"
        ...
    )

    var i int = 10
    i := 20 // will error out, bcoz i is already declare

}

```
Go compiler will not like if we have declare a variable but will not used it anywhere.

The error will be - j declared but not used

```go 
package main

import ("fmt")

func main() {

    var i int = 42
    j := 43

    fmt.Printf(i)

}
```


### Redeclaration and Shadowing
```go 

```


### Visibility
- lower case first letter for package scope
- upper case first letter to export
- no private scope. but you can declare a variable within a block and that variable can only be access with in that block.

```go
....
var I int = 100 // visible Globally

func main() {
    var i int = 42 // visible with in this function
}
```
### Naming Conventions
- Pascal or camelCase
- Capitalize acronyms(HTTP, URL)
- As short as reasonable

```go
var myName string = "Samit"
var theHTTP string = "https://x.net"
```

### Type Conversion
- destinationType(variable)
- use strconv package for string

```go
....

func main() {
    var i int = 42
    fmt.Printf("%v , %T\n", i, i)

    var j float32
    j = float32(i) // type conversion
    fmt.Printf("%v , %T\n", j, j)
}

```