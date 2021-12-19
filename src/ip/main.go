package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == http.MethodGet {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       request.Headers["x-nf-client-connection-ip"],
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       `bad request`,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
