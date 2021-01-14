package main

import (
	"fmt"
)

func main() {

	// Declare map
	myMap := make(map[string]int)

	// Print Initial map

	fmt.Println("Initial Map", myMap)

	// Assign values to map

	myMap["age"] = 22
	myMap["salary"] = 37800

	fmt.Println("After assignig values", myMap)
	fmt.Println("age is", myMap["age"])
	fmt.Println("Length of map", len(myMap))

	// delete map entry

	delete(myMap, "age")
	fmt.Println("After deleting", myMap)
	fmt.Println("age is", myMap["age"])

	// declaring and assignig values

	myMap2 := map[string]string{"name": "nikhil", "description": "software engineer"}
	fmt.Println("myMap2 values", myMap2)

	// --------------------------------------------------------------------------------
	// nested maps

	nestedMaps := make(map[string]map[string]string)

	// Print Initial nested mapted
	fmt.Println("Initial nested Map", nestedMaps)

	// Method2
	data := map[string]map[string]string{
		"a": map[string]string{},
		"b": map[string]string{},
		"c": map[string]string{},
	}

	data["a"]["w"] = "x"
	data["b"]["w"] = "x"
	data["c"]["w"] = "x"
	fmt.Println(data)

	// method3

	data2 := map[string]map[string]string{
		"a": map[string]string{
			"x": "y"},
		"b": map[string]string{
			"x": "y"},
		"c": map[string]string{
			"x": "y"},
	}

	fmt.Println("nested maps method2", data2)

}
