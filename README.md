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

- Immutable , but can be changed
- Replaced by the compiler at compile time
    - value must be calculated at compile time
- Named like variable
    - PascalCase for exported constant
    - camelCase for internal constant
- Typed constants work like immutable variables
    - Can interoperate only with same type
- UnTyped Constants work like immutable variables
    - can interoperate only with similar type
- Enumerated contsnat
    - special symbol iota allows related constants to be created easily
    - iota starts at 0 in each const block and increaments by one
    - Watch out of constant values that matches zero value for variable
- Enumerated expressions
    - Operations that can be determined at compile time are allowed
        - arithmetic
        - bitwise operations
        - bitshifting


## 03-Array-n-Slices
- Array
    - Collection of items with same type
    - Fixed Size
    - Declaration styles
        - a := [3]int{1,2,3}
        - a := [...]int{1,2,3}
        - var a [3]int
    - The index start from 0
    - len function returns size of array
    - Copies refers to different underlying data

- Slices
    - Backed by array
    - Creation style
        - Slice existing array or slice
        - literal style
        - Via Make function
            - a :=make([]int, 10) //create slice with capacity and length == 10
            - a :=make([]int, 10, 100) // crate slice with length == 10 and capacity == 100
        - len function returns of slice
        - cap function returns length of underlying array.
        - append function to add elements to slice
        - copies refers to same underlying array 


## 04-maps-n-structs
Both are collection types.

Topic to cover are :

- maps

```go
package main

import ("fmt")

func main() {
    statePopulation : map[string]int {
        "Odisha": 100000,
        "Maharastra" : 600000,
        "UttarPradesh" : 1000000
    }
    fmt.Println(statePopulation)
    fmt.Println(statePopulation["Maharastra"])
    
    //add
    statePopulation["Kerla"] = 800000
    fmt.Println(statePopulation["Kerla"])
    
    //delete a key from map
    delete(statePopulation,"Odisha")
    
    fmt.Println(statePopulation["Chatisgarh"]) // 0
    pop, ok := statePopulation["Chatisgarh"]
    fmt.Println(pop,ok) //0 false

    _, ok := statePopulation["Chatisgarh"]
    fmt.Println(ok) //0 false

    // len(mapVariableName) will give you the map lenght

    statePopulation := make(map[string]int,10) // define length of the map , but does not work well- TODO - explore more!
    
    // slice to a key of Map
    m := map[[3]int]string{} 

    //map copy. 
    //manipulation of one object , might have impact to the reference map
    sp := statePopulation
    delete(sp,"Odisha")
    fmt.Println(sp) // same result as statePopulation
    fmt.Println(statePopulation) // // same result as sp


}

```
- Structs
   
```go
package main

import ("fmt")

type Doctor struct {
    number int
    actorName string
    companions []string
}

func main() {
    aDoctor := Doctor {
        number : 3,
        actorName: "John",
        companions: []string{
            "A Paul",
            "Jo Grant"
        }
    }
    //also can be done like this , but not recomanded
    anotherDoctor := Doctor {
        4,
        "JDeb",
        []string{
            "A Mith",
            "J Vri"
        }
    }

    fmt.Println(aDoctor)
    fmt.Println(anotherDoctor)
    fmt.Println(aDoctor.actorName)
    fmt.Println(aDoctor.companions[1])

    //anonymous structs
    aMan := struct{name string}{name: "John Doe"}
    fmt.Println(aMan)

}

```


## Template
```go
package main

import ("fmt")

func main() {
    
}

```