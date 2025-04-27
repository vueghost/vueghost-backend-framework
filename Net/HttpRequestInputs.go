package Net

import (
	"fmt"
	"html"
	"net/http"
)

//HttpRequestInputs Class struct Pre-processes global input data for security.
type HttpRequestInputs struct {
	httpResponseWriter http.ResponseWriter
	httpRequest        *http.Request
}

func (c *HttpRequestInputs) Context(HttpResponseWriter http.ResponseWriter, HttpRequest *http.Request) {
	c.httpResponseWriter = HttpResponseWriter
	c.httpRequest = HttpRequest
}

//Post Fetch an item from the POST array.
func (c *HttpRequestInputs) Post(key string, escape ...bool) interface{} {
	var value string
	if c.httpRequest.Method == http.MethodPost {
		value = c.httpRequest.PostFormValue(key)
		return c.escapeValue(value, escape...)
	}
	return value
}

//Get Fetch an item from the GET array.
func (c *HttpRequestInputs) Get(key string, escape ...bool) interface{} {
	var value string
	if c.httpRequest.Method == http.MethodGet {
		value = c.httpRequest.FormValue(key)
		return c.escapeValue(value, escape...)
	}
	return value
}

//FormValue
func (c *HttpRequestInputs) FormValue(key string, escape ...bool) interface{} {
	var value string
	value = c.httpRequest.FormValue(key)
	return c.escapeValue(value, escape...)
}

//Cookie
func (c *HttpRequestInputs) Cookie() {

}

//File
func (c *HttpRequestInputs) File(key string) {

}

//Photo
func (c *HttpRequestInputs) Photo() {}

//escapeValue
func (c HttpRequestInputs) escapeValue(value string, escape ...bool) interface{} {
	if len(escape) > 0 {
		if escape[0] {
			return html.EscapeString(value)
		} else {
			return value
		}
	} else {
		return html.EscapeString(value)
	}
}

//@throwError HttpRequestInputs throw error handler.
func (c HttpRequestInputs) throwError(err error) {
	_ = fmt.Errorf(
		"HttpRequestInputs Error %v",
		err,
	)
}
