package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", "world"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
