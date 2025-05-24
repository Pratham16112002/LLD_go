package builder

import "net/http"

type RequestBuilder interface {
	Method(method string) RequestBuilder
	URL(url string) RequestBuilder
	Header(key, value string) RequestBuilder
	Build() (*http.Request, error)
}

type HTTPRequestBuilder struct {
	method  string
	url     string
	headers map[string]string
}

func NewHTTPRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		headers: make(map[string]string),
	}
}

func (rb *HTTPRequestBuilder) Method(method string) RequestBuilder {
	rb.method = method
	return rb
}

func (rb *HTTPRequestBuilder) URL(url string) RequestBuilder {
	rb.url = url
	return rb
}

func (rb *HTTPRequestBuilder) Header(key, value string) RequestBuilder {
	rb.headers[key] = value
	return rb
}

func (rb *HTTPRequestBuilder) Build() (*http.Request, error) {
	req, err := http.NewRequest(rb.method, rb.url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range rb.headers {
		req.Header.Set(key, value)
	}
	return req, nil
}
