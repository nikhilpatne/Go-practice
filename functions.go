package main

import "fmt"

func main() {
	// a, b := 12, 15

	var a, b int = 15, 17
	//  Addition of two numbers
	result := add(a, b)
	fmt.Printf("Addition of two number is %d\n", result)

	//  Greater number
	fmt.Printf("Max number is %d \n", max(a, b))

	// Swapping of names
	num1, num2 := swap(a, b)
	fmt.Printf("Numbers after swap %d and %d \n", num1, num2)

}

func add(a, b int) int {
	return a + b
}

func max(a, b int) int {
	var max int

	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

func swap(a, b int) (int, int) {
	return b, a
}
