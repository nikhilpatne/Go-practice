package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	myChannel := make(chan int)

	go sum(a, myChannel)
	go sum(a, myChannel)

	result1, result2 := <-myChannel, <-myChannel

	fmt.Println(result1, result2)

	fmt.Println("Total time needed", time.Since(now))
}

func sum(a []int, myChannel chan int) {

	time.Sleep(time.Second * 1)
	num := 0
	for _, value := range a {
		num += value
	}

	myChannel <- num
}
