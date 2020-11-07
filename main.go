package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func response(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}

func helloHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(map[string]string{"message": "Hello, World!"})
	if err != nil {
		return response("", http.StatusInternalServerError), err
	}
	return response(string(body), http.StatusOK), nil
}

func main() {
	lambda.Start(helloHandler)
}
