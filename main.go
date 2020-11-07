package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func response(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}

func helloHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return response("Hello World", http.StatusOK), nil
}

func main() {
	lambda.Start(helloHandler)
}
