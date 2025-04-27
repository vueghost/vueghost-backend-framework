package System

import (
	"Framework/Net"
	"fmt"
	"net/http"
)

type Controller struct {
	httpResponseWriter http.ResponseWriter
	httpRequest        *http.Request
	HttpRequestHeader  Net.HttpRequestHeader
	Input              Net.HttpRequestInputs
}

//Context controller context.
func (c *Controller) Context(HttpResponseWriter http.ResponseWriter, HttpRequest *http.Request) {
	c.httpRequest = HttpRequest
	c.httpResponseWriter = HttpResponseWriter
	c.HttpRequestHeader.Context(c.httpRequest)
}

//Constructor The controller constructor.
func (c *Controller) Constructor() {}

//Finalize The controller destructor.
func (c *Controller) Finalize() {}

//AboutController some information about the controller.
func (c *Controller) AboutController() {
	fmt.Println("This just about controller")
}

//GetHttpRequest Return httpRequest.
func (c *Controller) GetHttpRequest() *http.Request {
	return c.httpRequest
}

//throwError error exception of controller handler.
func (c Controller) throwError(err error) {
	_ = fmt.Errorf(
		"Controller error %v",
		err,
	)
}
