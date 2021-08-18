package main

import (
	"fmt"
)

func main() {
	const a int = 42
	const b = "Hello, "
	const (
		c      = 35.6
		d      = "Samit"
		e bool = false
	)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)
	fmt.Printf("%v %T\n", e, e)

	const (
		f = iota
		g = iota
		h = iota
	)
	fmt.Printf("%v %T\n", f, f)
	fmt.Printf("%v %T\n", g, g)
	fmt.Printf("%v %T\n", h, h)

	fmt.Printf("---------------\n") //output = 0
	const (
		_ = iota
		i
		j
		k
	)
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v %T\n", k, k)
}

//output

/*
42 int
Hello,  string
35.6 float64
Samit string
false bool
0 int
1 int
2 int
---------------
1 int
2 int
3 int

*/
