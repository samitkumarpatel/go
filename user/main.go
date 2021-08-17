package main

import (
	"fmt"

	"example.com/greetings"
	"example.com/math"
)

func main() {
	message := greetings.Greet("Samit")

	fmt.Println(message)

	fmt.Println(fmt.Sprintf("Addition of 5 +5 = %v", math.Add(5, 5)))

	fmt.Println(fmt.Sprintf("Substraction of 5 - 3 = %v", math.Add(5, 3)))
}
