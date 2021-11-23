//Lambda Function Go
package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest) //start the whole workflow for the lambda requests   : handleRequest = function that i am going to make to handle the logic
	// it will get the GET method and then query the database
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) { //APIGatewayProxyRequest : will contain JSON for query parameters and headers
	if request.HTTPMethod == "GET" {
		var stringResponse string = "Yay a successful repsponse!!"
		APIResponse := events.APIGatewayProxyResponse{Body: stringResponse, StatusCode: 200}
		return APIResponse, nil
	} else {
		err := errors.New("method not allowed")
		APIResponse := events.APIGatewayProxyResponse{Body: "Method Not OK", StatusCode: 502}
		return APIResponse, err
	}
}
