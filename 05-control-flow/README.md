## Control Flow

- if statement

```go
package main

import ("fmt")

func main() {
    //example-1
    if true {
        fmt.Println("The test is True")
    }

    //example-2
    statePopulation := map[string]int {
        "Odisha" : 100000,
        "Maharastra" : 20000,
        "Punjab" : 1500000
    }
    if pop,ok := statePopulation["Maharastra"]; ok {
        fmt.Println(pop)
    }

    //example-3
    number := 10
    guess := 30
    if number < guess {
        fmt.Print("")
    }
    if number > guess {
        fmt.Print("")
    }
    if number == guess {
        fmt.Print("")
    }
    //example-4
    fmt.Println(number<=guess, number>=guess, number==guess)
    
    //example-5
    if number < 1 || guess > 100 {
        fmt.Print("")
    }
    if number < 1 && guess > 100 {
        fmt.Print("")
    }
    //example-6
    fmt.Print(!true)

    //if - else if - else
    if true {
        fmt.Print("")
    } else if 1 > 10 {
        fmt.Print("")
    } else {
        fmt.Print("")
    }

}

```

- switch statement
    - simple cases
    - cases with multiple tests
    - Falling through
    - Type switches
```go
package main

import ("fmt")

func main() {
    switch 2 {
        case 1, 5, 10:
            fmt.Println("one, five or ten")
        case 2:
            fmt.Println("two")
        default:
            fmt.Println("not one or two")
    }

    switch i:= 2+3;i {
        case 1,5,10:
            fmt.Println("")
        case 2,4,6:
            fmt.Println("")
        default:
            fmt.Println("")
    }

    //tag list syntex
    i := 10
    switch {
        case i <= 10:
            fmt.Println("")
            fallthrough
        case i <= 20:
            fmt.Println("")
            break
        default:
            fmt.Println("")
    }

    //type switch
    var i interface{} = 1 //can assigh to anything
    switch i.(type) {
        case int:
            fmt.Println("")
        case string:
            fmt.Println("")
        case float64:
            fmt.Println("")
        default:
            fmt.Println("")
    }
}

```