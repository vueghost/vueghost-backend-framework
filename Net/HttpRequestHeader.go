package Net

import (
	"net/http"
)

//HttpRequestHeader The HTTP headers that may be specified in a client request.
type HttpRequestHeader struct {
	httpRequest *http.Request
	//UserAgent The User-Agent header, which specifies information about the client agent.
	UserAgent interface{}
}

func (c *HttpRequestHeader) Context(HttpRequest *http.Request) {
	c.httpRequest = HttpRequest
}
