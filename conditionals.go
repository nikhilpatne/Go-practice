package main

import "fmt"

func main() {
	num1, num2, num3 := 12, 34, 67

	if num1 >= num2 && num1 <= num3 {
		fmt.Printf("%d is the largets number \n", num1)
	} else if num2 >= num1 && num2 >= num3 {
		fmt.Printf("%d is the largets number \n", num2)
	} else {
		fmt.Printf("%d is the largest number \n", num3)
	}
}
