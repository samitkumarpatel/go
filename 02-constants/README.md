- Naming convention
 like variable 

```go 
package main

import (
    "fmt"
    "math"
    )
const d int16 = 27 // package level constant

func main() {
    const d int = 14 // this will win over package level constant

    const myConst int = 42 // Type constant
    myConst = 27 // ERROR , because ita a const
    
    const myConstOne = math.Sin(1.57) // will error out

    // type constant
    const a int = 20 
    const b string = "foo"
    const c float32 = 3.14
    const d bool = true

    // Untype Constant
    const e = 42 // int type 

    // Enumurated constant 
    const (
        i = iota
        j = iota
        k = iota
    )

    //also work like above
     const (
        i = iota
        j
        k
    )
    const (
        a2 = iota
    )
    fmt.Printf("%v\n", i) //output = 0
    fmt.Printf("%v\n", j) //output = 1
    fmt.Printf("%v\n", k) //output = 2
    fmt.Printf("%v\n", a2) //output = 0

    const (
        _ = iota + 5
        x
        y
        z
    )
    var g int
    fmt.Printf("%v\n", a == x) //false
    fmt.Printf("%v\n", x) //output = 6
}

```

- Type Constants

```go 
const e int = 42 // int type // type Constant

```
- Untype Constants

```go 
const e = 42 // int type // Untype Constant

```

- Enumerated Constants / expression

```go 
const (
        i = iota
        j = iota
        k = iota
    )

    //also work like above
     const (
        i = iota
        j
        k
    )
    const (
        a2 = iota
    )
    fmt.Printf("%v\n", i) //output = 0
    fmt.Printf("%v\n", j) //output = 1
    fmt.Printf("%v\n", k) //output = 2
    fmt.Printf("%v\n", a2) //output = 0

    const (
        _ = iota + 5
        x
        y
        z
    )
    var g int
    fmt.Printf("%v\n", a == x) //false
    fmt.Printf("%v\n", x) //output = 6
```