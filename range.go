package main

import "fmt"

func main() {

	// range on array
	myArray := []int{1, 2, 3, 4}

	for index, value := range myArray {
		fmt.Printf("index is %d and values id %d \n", index, value)
	}

	// sum of elements of array
	sum := 0
	for _, value := range myArray {
		sum += value
	}
	fmt.Printf("sum of elemets is %d \n", sum)

	// sum of elements of array using function
	result := sum_of_elements(myArray)
	fmt.Printf("sum of array elemets using func %d \n", result)
	// Range on map

	myMap := map[string]string{"name": "nikhil", "surname": "patne"}
	for key, value := range myMap {
		fmt.Printf("key -> %s value -> %s\n", key, value)
	}

	//  To print only values

	for _, value := range myMap {
		fmt.Printf("Values are %s \n", value)
	}

}

func sum_of_elements(myArray []int) int {
	sum := 0
	for _, value := range myArray {
		sum += value
	}
	return sum
}
