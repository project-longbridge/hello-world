package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Interface implements the generic interface, this should change in the future
type Interface interface{}

func response(body Interface, statusCode int) events.APIGatewayProxyResponse {
	bodyStr, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(bodyStr),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}

func helloHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return response(map[string]string{"message": "Hello, World!"}, http.StatusOK), nil
}

func main() {
	lambda.Start(helloHandler)
}
