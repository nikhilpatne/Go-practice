package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("Initial Array ", a)

	// Set value to one of index
	a[3] = 70
	fmt.Println("Array after setting value", a)
	fmt.Println("length of an array", len(a))

	//  Declare and assign values to array

	b := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Array b is", b)

	//  Another way to create array

	myarray := make([]string, len(b))
	fmt.Println("Another array is", myarray)

	// Two dimensional Array

	var twoDimensionalArray [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			twoDimensionalArray[i][j] = i + j
		}
	}
	fmt.Println("Two dimensional", twoDimensionalArray)

	// Print Pattern

	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf(" * ")
		}
		fmt.Println()
	}
}
