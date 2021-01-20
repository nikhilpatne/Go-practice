package main

import (
	"errors"
	"fmt"
)

func main() {

	var start, end int

	fmt.Printf("Enter start number \n")
	fmt.Scanln(&start)

	fmt.Printf("Enter end number \n")
	fmt.Scanln(&end)

	result, err := sum_of_elements(start, end)

	if err != nil {
		fmt.Println("Error occured  ->", err)
	} else {
		fmt.Printf("%d is the sum of elements \n", result)
	}
}

// sum of array elements
//   :given ->  start and end number

func sum_of_elements(start int, end int) (int, error) {

	if start > end {
		return 0, errors.New("First number should be small")
	}
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum, nil
}

// type error interface {
// 	Error() string
// }

//  error is the built is interface defined by go that i used  *error*
