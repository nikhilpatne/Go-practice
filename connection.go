package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("AKIAYZVWFK5MZKOX7XER", "cpy66QWVAf7zhF4ZdHv0O9BfYKD2UK/w4mKwlBpV", ""),
	})

	svc := ec2.New(sess)
	_, err := svc.DescribeKeyPairs(nil)
	if err != nil {
		fmt.Printf("Errrror")
	}
}
