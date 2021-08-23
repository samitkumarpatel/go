package main

import (
	"fmt"
)

func main() {
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(x)

	y := x
	y[0] = 0
	fmt.Println(x)
	fmt.Println(y)
}
