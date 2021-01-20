package main

import (
	"fmt"
	"net/http"
)

func main() {

	req, err := http.NewRequest("GET", "https://api.github.com/users/sakar97/repos", nil)
	if err != nil {
		fmt.Println("Error occured", err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error occured", err.Error())
	}
	fmt.Printf("%T", resp)

}
