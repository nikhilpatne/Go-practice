package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// if len(os.Args) != 2 {
	// 	exitErrorf("pair name required\nUsage: %s key_pair_name",
	// 		filepath.Base(os.Args[0]))
	// }
	pairName := "hhh"
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-east-2")},
	// )

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("AKIAYZVWFK5MZKOX7XE", "cpy66QWVAf7zhF4ZdHv0O9BfYKD2UK/w4mKwlBpV", ""),
	})

	// fmt.Println(sess)
	// Create an EC2 service client.
	svc := ec2.New(sess)

	fmt.Println(svc)

	result, err := svc.CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName: aws.String(pairName),
	})
	if err != nil {

		fmt.Println("OOOOOOOOOOOOps")
		// if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "InvalidKeyPair.Duplicate" {
		// 	exitErrorf("Keypair %q already exists.", pairName)
		// }
		// exitErrorf("Unable to create key pair: %s, %v.", pairName, err)
	}

	fmt.Printf("Created key pair %q %s\n%s\n",
		*result.KeyName, *result.KeyFingerprint,
		*result.KeyMaterial)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

// sess, err := session.NewSession(&aws.Config{
//     Region:      aws.String("us-west-2"),
//     Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
// })
