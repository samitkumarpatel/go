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
```go
package main

import ("fmt")

func main() {
    grades := [3]int{97,85,93} // array of integer can hold 3 element
    fmt.Printf("Grades: %v",grades)
    //also can be done like below 
    grades := [...]int{97,85,93} // array of integer can hold 3 element
    fmt.Printf("Grades: %v",grades)

    var students [3]string
    fmt.Printf("students: %v",students)
    students[0] = "Samit"
    students[1] = "Hello"
    students[2] = "Hi"
    fmt.Printf("students: %v",students)
    fmt.Printf("students: %v",students[1])
    fmt.Printf("Size Of Students: %v",len(students))

    //array of array
    var identityMatrix [3][3]int = [3][3]int{ [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1} }
    fmt.Println(identityMatrix)
    //simple and readability way
    var identityMatrix [3][3]int
    identityMatrix[0] = [3]int{1,0,0}
    identityMatrix[1] = [3]int{0,1,0}
    identityMatrix[2] = [3]int{0,0,1}
    fmt.Println(identityMatrix)

    //array copy
    a := [...]int{1,2,3}
    b := a
    b[1] = 5
    fmt.Println(a) // [1 2 3]
    fmt.Println(b) // [1 5 3]
    // but if when assign array a to array be , if we give the pointer details the value would be same for both
    a := [...]int{1,2,3}
    b := &a
    b[1] = 5
    fmt.Println(a) // [1 5 3]
    fmt.Println(b) // &[1 5 3]

}

```

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

In slice predefine size is not needed. 

```go
package main

import ("fmt")

func main() {
    a := []int{1, 2, 3}
    b := a
    b[1] = 5
    fmt.Println(a) // [1 5 3]
    fmt.Println(b) // [1 5 3]

    a := []int{1,2,3,4,5,6,7,8}
    b := a[:] //slice of all element
    c := a[3:] //slice from 4th element to end
    d := a[:6] //slice first 6 elelment
    e := a[3:6] //slice the 4th , 5th and 6th elements
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)

    //buildin make function
    a := make([]int,3)
    fmt.Println(a) //[0 0 0 ]
    fmt.Printf("Length: %v", len(a)) //3
    fmt.Printf("Capacity: %v", cap(a)) //3 , cap = capacity

    a := make([]int,3,100) // Length 3 and Capacity is 100
    //we can also add avalue to Slice like below with append method
    a = append(a,1)
    a = append(a,2,3,4,5)

    //push and pop
    b := a[1:]
    fmt.Println(b) //[2 3 45]
    b := a[:len(a)-1] //pop
    
}

```