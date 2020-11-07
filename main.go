package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func helloHandler() (string, error) {
	return "hello", nil
}

func main() {
	lambda.Start(helloHandler)
}
