package handler

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

func Handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Hello duniya walo!! How are you doing?")
	log.Printf(
		"Headers: %v",
		marshal(request.Headers),
	)

	log.Printf(
		"Pathparams: %v",
		marshal(request.PathParameters),
	)

	log.Printf(
		"Queryparams: %v",
		marshal(request.QueryStringParameters),
	)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello World",
	}, nil
}

func marshal(obj map[string]string) string {
	res, _ := json.Marshal(obj)
	return string(res)
}
