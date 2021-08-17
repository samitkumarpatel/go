# golang
This repos contain my learning and reference meterials as code.

- [Golang Tutorial](https://golang.org/doc/tutorial/) - module, package and etc..]
- [Developing and Publish a module](https://golang.org/doc/modules/developing)
- [Create Module](https://golang.org/doc/tutorial/create-module) 
- [Call Module](https://golang.org/doc/tutorial/call-module-code)
- [go package](https://pkg.go.dev/)


## 00-variable

- lower case first letter for package scope
- upper case first letter to export
- no private scope. but you can declare a variable within a block and that variable can only be access with in that block.

```go
package main
import ( "fmt" )

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
    i := 20 // This will error out, bcoz i is already declare

}
```
## 01-primitives
The types in Go can be classified as follows −
- Boolean types
- Numeric types
- String types
- Derived types
The more details can be found in [tutorialspoint](https://www.tutorialspoint.com/go/go_data_types.htm) and well explained in this [video](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=3425s)


## 02-constants

## 03-Array-n-Slices
