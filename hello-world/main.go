package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"hello-world/handler"
)

func main() {
	lambda.Start(handler.Handle)
}
