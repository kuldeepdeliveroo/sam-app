package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handler() {
	log.Printf("Hello duniya walo!! How are you doing?")
}

func main() {
	lambda.Start(handler)
}
