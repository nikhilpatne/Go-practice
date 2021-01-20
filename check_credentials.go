package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	_, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("AKIAYZVWFK5MZKOX7XE", "cpy66QWVAf7zhF4ZdHv0O9BfYKD2UK/w4mKwlBpV", ""),
	})

	if err != nil {
		fmt.Println(err)
	}
}
