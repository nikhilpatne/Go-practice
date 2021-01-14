package main

import "fmt"

func main() {
	// infinit_loop()

	// For loop written like while loop
	i := 1
	for i <= 3 {
		fmt.Println("value of i (way1) ", i)
		i++
	}

	// Way2:

	for i := 1; i < 5; i++ {
		fmt.Println("value of i (way2) ", i)
	}

	

}

func infinit_loop() {
	for {
		fmt.Println("Im am in Infinite Loops")
	}
}
