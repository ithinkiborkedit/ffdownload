package src

import (
	"io"
	"net/http"
)

type RequestBuilder struct {
	Url    string
	Method string
	Header http.Header
	Body   io.Reader
}

func NewRequestBuilder(Url, Method string) *RequestBuilder {
	return &RequestBuilder{
		Url:    Url,
		Method: Method,
		Header: make(http.Header),
	}
}

func (rb *RequestBuilder) SetHeader(Key, Value string) *RequestBuilder {
	rb.Header.Set(Key, Value)
	return rb
}

func (rb *RequestBuilder) SetBody(Body io.Reader) *RequestBuilder {
	rb.Body = Body
	return rb
}

func (rb *RequestBuilder) Build() (*http.Request, error) {
	req, err := http.NewRequest(rb.Method, rb.Url, rb.Body)
	if err != nil {
		return nil, err
	}
	req.Header = rb.Header
	return req, nil
}
