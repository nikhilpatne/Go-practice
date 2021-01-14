package main

import "fmt"

func main() {

	// var a [10]int
	a := make([]int, 10)

	for i := 0; i < 10; i++ {
		a[i] = i * 2
	}

	fmt.Println("My Array", a)

	a = append(a, 20, 22, 23)
	fmt.Println("My Array after append", a)

	// copy array

	b := make([]int, len(a))
	copy(b, a)
	fmt.Println("Copied array b", b)

	// Slicing

	slice1 := b[:]
	fmt.Println("Slicing of slice1", slice1)

	slice2 := b[2:5] // started with 3rd element ends with 5th element.
	fmt.Println("Slicing of slice2", slice2)

	slice3 := b[:5] // First 5th element.
	fmt.Println("Slicing of slice3", slice3)

	slice4 := b[5:] // started with 6th element.......
	fmt.Println("Slicing of slice4", slice4)

}
