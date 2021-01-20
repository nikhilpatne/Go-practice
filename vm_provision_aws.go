package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

func MakeInstance(svc ec2iface.EC2API, instance_name string, image string, size string) (*ec2.Reservation, error) {
	result, err := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String(image),
		InstanceType: aws.String(size),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})
	if err != nil {
		return nil, err
	}

	_, err = svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{result.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String(instance_name),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("AKIAYZVWFK5MZKOX7XER", "cpy66QWVAf7zhF4ZdHv0O9BfYKD2UK/w4mKwlBpV", ""),
	})

	svc := ec2.New(sess)

	result, err := MakeInstance(svc, "nikhil-web", "ami-0a91cd140a1fc148a", "t2.micro")
	if err != nil {
		fmt.Println("Got an error creating an instance with tag")
		return
	}

	fmt.Println("Created tagged instance with ID " + *result.Instances[0].InstanceId)
}

// snippet-end:[ec2.go.create_instance_with_tag]
