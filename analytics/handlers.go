package main

type RequestContainer interface {
	GetRequest() *Request
}

type Request struct {
	HTTPMethod string
	Headers    map[string]string
}

type Response struct {
	StatusCode      int
	Headers         map[string]string
	Body            string
	IsBase64Encoded bool
}

func main_handler(req RequestContainer) (Response, error) {
	request := req.GetRequest()
	content_type := request.Headers["Content-Type"]
	if request.HTTPMethod == "POST" && content_type == "application/json" {
		return Response{
			StatusCode:      200,
			Headers:         map[string]string{"Content-Type": "text/plain"},
			Body:            "Hi Renato",
			IsBase64Encoded: false,
		}, nil
	}
	if request.HTTPMethod != "POST" {
		return Response{
			StatusCode: 405,
		}, nil
	}
	return Response{
		StatusCode:      400,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            "Bad Request. Content-Type not acceptable: " + content_type,
		IsBase64Encoded: false,
	}, nil
}
