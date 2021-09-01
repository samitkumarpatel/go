package main

import (
	"fmt"
)

func main() {

	//example-1
	myMap := map[string]int{
		"Odisha":     10000,
		"WestBangol": 21000,
	}

	//example-2
	x := map[string]int{}
	x["a"] = 10

	//example-3
	var myOtherMap map[string]string
	myOtherMap = make(map[string]string)
	//or
	//myOtherMap = make(map[string]string, 10)
	myOtherMap["a"] = "apple"
	myOtherMap["b"] = "ball"

	fmt.Println(myMap)
	fmt.Println(x)
	fmt.Println(myOtherMap)

	myMap["Maharastra"] = 5000000
	fmt.Println(myMap)

	//delete and evaluate
	delete(myMap, "Odisha")
	_, ok := myMap["Odisha"]
	fmt.Println(ok)
}
