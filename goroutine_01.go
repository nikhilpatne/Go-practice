package main

import (
	"fmt"
	"time"
)

func main() {
	go count("test") // this works in background, async
	go count("test1")

	fmt.Scanln()
}

func count(test string) {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println(test)
	}
}

      