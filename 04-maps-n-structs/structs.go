package main

import (
	"fmt"
)

type Department struct {
	id   int
	name string
}

type Employee struct {
	Department
	name   string
	id     int
	salary float32
}

func main() {
	employeeOne := Employee{
		Department: Department{id: 10, name: "HR"},
		name:       "Samit",
		id:         1001,
		salary:     60000,
	}

	fmt.Println(employeeOne)
	fmt.Println(employeeOne.name)
	fmt.Println(employeeOne.salary)

	//example of a map with Struct
	var employee map[string]Employee
	employee = make(map[string]Employee)
	employee["samit"] = Employee{
		Department: Department{id: 10, name: "HR"},
		name:       "Samit",
		id:         1001,
		salary:     60000,
	}

	fmt.Println(employee)
}
