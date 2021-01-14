package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Print("Today")
	case today + 1:
		fmt.Print("Tommorow")
	default:
		fmt.Print("Saturday is Too Far \n")

	}
}
