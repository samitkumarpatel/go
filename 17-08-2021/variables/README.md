### Variable Defination in Go
var variable_list optional_data_type;
```
var  i, j, k int;
var  c, ch byte;
var  f, salary float32;
d =  42; // int
```

### Static Type Declaration in Go
A static type variable declaration provides assurance to the compiler that there is one variable available with the given type and name so that the compiler can proceed for further compilation without requiring the complete detail of the variable. A variable declaration has its meaning at the time of compilation only

```go
...

func main() {
   var x float64
   x = 20.0
   fmt.Println(x)
   fmt.Printf("x is of type %T\n", x)
}
```

### Dynamic Type Declaration / Type Inference in Go
A dynamic type variable declaration requires the compiler to interpret the type of the variable based on the value passed to it
```go
....
func main() {
   var x float64 = 20.0

   y := 42 
   fmt.Println(x)
   fmt.Println(y)
   fmt.Printf("x is of type %T\n", x)
   fmt.Printf("y is of type %T\n", y)	
}
```

### Mixed Variable Declaration in Go
Variables of different types can be declared in one go using type inference.
```go
...
func main() {
   var a, b, c = 3, 4, "foo"  
	
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)
}
```