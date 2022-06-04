package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type realRequest struct {
	request events.APIGatewayProxyRequest
}

func (req *realRequest) GetRequest() *Request {
	return &Request{
		HTTPMethod: req.request.HTTPMethod,
		Headers:    req.request.Headers,
	}
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	response, err := main_handler(&realRequest{request: request})
	if err != nil {
		return nil, err
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      response.StatusCode,
		Headers:         response.Headers,
		Body:            response.Body,
		IsBase64Encoded: request.IsBase64Encoded,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
