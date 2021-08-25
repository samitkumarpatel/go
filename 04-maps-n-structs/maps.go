package main

import (
	"fmt"
)

func main() {

	myMap := map[string]int{
		"Odisha":     10000,
		"WestBangol": 21000,
	}

	fmt.Println(myMap)
	myMap["Maharastra"] = 5000000
	fmt.Println(myMap)

	//delete and evaluate
	delete(myMap, "Odisha")
	_, ok := myMap["Odisha"]
	fmt.Println(ok)

}
