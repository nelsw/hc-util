package aws

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

var headers = map[string]string{"Access-Control-Allow-Origin": "*"}

type Headers map[string]string
type Object interface{}

type ResponseBuilder interface {
	Head(Headers) ResponseBuilder
	Code(int) ResponseBuilder
	Text(string) ResponseBuilder
	Toke(string) ResponseBuilder
	Data(Object) ResponseBuilder
	Build() (events.APIGatewayProxyResponse, error)
}

type responseBuilder struct {
	statusCode int
	headers    Headers
	body       string
	cookie     string
	object     Object
}

func (r *responseBuilder) Head(headers Headers) ResponseBuilder {
	r.headers = headers
	return r
}

func (r *responseBuilder) Code(statusCode int) ResponseBuilder {
	r.statusCode = statusCode
	return r
}

func (r *responseBuilder) Text(body string) ResponseBuilder {
	r.body = body
	return r
}

func (r *responseBuilder) Data(object Object) ResponseBuilder {
	r.object = object
	return r
}

func (r *responseBuilder) Toke(cookie string) ResponseBuilder {
	r.cookie = cookie
	return r
}

func (r *responseBuilder) Build() (events.APIGatewayProxyResponse, error) {
	if r.cookie != "" {
		r.headers["Set-Cookie"] = string(r.cookie)
	}
	if r.object != nil {
		// json package is resilient enough to marshal any non nil value, gl.
		bytes, _ := json.Marshal(r.object)
		r.body = string(bytes)
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      int(r.statusCode),
		Headers:         r.headers,
		Body:            string(r.body),
		IsBase64Encoded: false,
	}, nil
}

func New() ResponseBuilder { return &responseBuilder{headers: headers} }
