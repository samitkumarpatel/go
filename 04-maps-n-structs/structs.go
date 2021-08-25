package main

import (
	"fmt"
)

type Employee struct {
	name   string
	id     int
	salary float32
}

func main() {
	employeeOne := Employee{
		name:   "Samit",
		id:     1001,
		salary: 60000,
	}

	fmt.Println(employeeOne)
	fmt.Println(employeeOne.name)
	fmt.Println(employeeOne.salary)
}
