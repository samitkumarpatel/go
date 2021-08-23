package main

import (
	"fmt"
)

func main() {

	var myArray = [3]int{1, 2, 3}
	fmt.Println(myArray)
	fmt.Println(myArray[0])
	fmt.Println(myArray[1])

	var student [3]string
	student[0] = "samit"
	fmt.Println(student)

	grade := [...]string{"a", "b", "c"}
	fmt.Println(grade)

	a := [2]int{1, 2}
	b := a
	b[0] = 0
	fmt.Println(a)
	fmt.Println(b)

	c := [2]string{"Hello", "World"}
	d := &c
	d[0] = "Hi"
	fmt.Println(c)
	fmt.Println(d)

}
