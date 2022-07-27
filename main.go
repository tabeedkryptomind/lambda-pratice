package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

)


func main() {
	lambda.Start(HandleRequest)
}


func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

		switch req.HTTPMethod {
			case "GET":
				if req.PathParameters["name"] != "" {
					return events.APIGatewayProxyResponse{Body: "hey ..!" + req.Body, StatusCode: 200}, nil
				}
			default:
				return events.APIGatewayProxyResponse{StatusCode: http.StatusMethodNotAllowed}, nil
		}
		return events.APIGatewayProxyResponse{StatusCode: http.StatusMethodNotAllowed}, nil
}