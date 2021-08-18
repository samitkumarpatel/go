package main

import (
	"fmt"
)

var (
	I             = "I'm available Globally!"
	name   string = "Samit Kumar Patel"
	id            = 1001
	salary        = 6000.01
)

func main() {
	i := "Smitika"
	j := 10
	k := 10.4
	var m string
	m = "Hello"
	I = "What?"
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", m, m)

	fmt.Printf("%v, %T\n", I, I)
	fmt.Printf("%v, %T\n", id, id)
	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", salary, salary)
}
