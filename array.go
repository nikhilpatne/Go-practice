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

	// MultiDimensional array
	sample := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
	fmt.Printf("Overall Dimension of the array: %d*%d*%d\n", len(sample), len(sample[0]), len(sample[0][0]))

	for _, first := range sample {
		for _, second := range first {
			for _, value := range second {
				fmt.Println(value)
			}
		}
	}

	//  Two Dimensional Slices.

	twoDimensionalSlices := make([][]int, 2)

	twoDimensionalSlices[0] = []int{1, 3, 4}
	twoDimensionalSlices[1] = []int{3, 5, 2}

	fmt.Println("SLice: ", twoDimensionalSlices)

	fmt.Println("length: ", len(twoDimensionalSlices), "*", len(twoDimensionalSlices[0]))

}
