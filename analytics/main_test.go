package main

import "testing"

type testRequestContainer struct {
	req Request
}

func (container testRequestContainer) GetRequest() *Request {
	return &container.req
}

func request(method, contentType string) RequestContainer {
	return testRequestContainer{req: Request{
		HTTPMethod: method,
		Headers:    map[string]string{"Content-Type": contentType}},
	}
}

func TestHandler(t *testing.T) {
	cases := []struct {
		method, contentType string
		statusCode          int
		body                string
	}{
		{"GET", "application/json", 405, ""},
		{"POST", "text/plain", 400, "Bad Request. Content-Type not acceptable: text/plain"},
		{"POST", "application/json", 200, "Hi Renato"},
	}
	for _, c := range cases {
		resp, err := main_handler(request(c.method, c.contentType))
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != c.statusCode {
			t.Errorf("statusCode=%d, want %d", resp.StatusCode, c.statusCode)
		}
		if resp.Body != c.body {
			t.Errorf("statusCode=%q, want %q", resp.Body, c.body)
		}
	}
}
