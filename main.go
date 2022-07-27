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
				return events.APIGatewayProxyResponse{Body: "hey ..!", StatusCode: 200}, nil
			case "POST":
				name := req.QueryStringParameters["name"]
				return events.APIGatewayProxyResponse{Body: "Nice to meet you" + name, StatusCode: 200}, nil
			default:
				return events.APIGatewayProxyResponse{StatusCode: http.StatusMethodNotAllowed}, nil
		}
}