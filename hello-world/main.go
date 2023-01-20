package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handler() (events.APIGatewayProxyResponse, error) {
	log.Printf("Hello duniya walo!! How are you doing?")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello World",
	}, nil
}

func main() {
	lambda.Start(handler)
}
