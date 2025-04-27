package System

import (
	"Framework/Extensions"
	"Framework/Security"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//ApiController API controller is a class which can be created under the Controllers folder or any other folder under your project's root folder
type ApiController struct {
	Controller
	Output         *ApiOutput
	Session        Extensions.Session
	authentication Security.Authentication
}

//Context
func (c *ApiController) Context(HttpResponseWriter http.ResponseWriter, HttpRequest *http.Request) {
	HttpResponseWriter.Header().Set("Content-Type", "application/x-javascript; charset=utf-8")
	c.httpRequest = HttpRequest
	c.httpResponseWriter = HttpResponseWriter
	c.HttpRequestHeader.Context(c.httpRequest)

	//@Inputs & outputs.
	c.Input.Context(HttpResponseWriter, HttpRequest)
	c.Output = NewApiOutput(ApiOutputStore{})
	c.Output.SetSuccess(false)
	c.Output.SetAuth(true)

	//@Authentication
	c.authentication = Security.Authentication{}
	c.Session.ID = c.authentication.Context(HttpRequest)

}

//InvalidInputs
func (c *ApiController) InvalidInputs(errorHandler ...ErrorHandler) {
	if len(errorHandler) > 0 {
		c.Output.SetErrors(errorHandler)
	} else {
		c.Output.SetErrors(ErrorHandler{
			Code:    100,
			Type:    "ValidationException",
			Message: "Invalid inputs",
		})
	}
}

//IsAuthorized
func (c *ApiController) IsAuthorized() bool {
	if c.Session.ID == "0" ||
		c.Session.ID == 0 {
		c.Output.SetAuth(false)
		c.Output.SetErrors(ErrorHandler{
			Message: "Invalid authentication session.",
			Type:    "OAuthException",
			Code:    468,
		})
		return false
	}
	c.Output.SetAuth(true)
	return true
}

//SetAuthToken authentication set new auth token.
func (c *ApiController) SetAuthToken(userID interface{}) string {
	return c.authentication.SetAuthToken(userID)
}

//ActAsPage Acting like page
func (c *ApiController) ActAsPage() {
	c.Output.SetSuccess(true)
	c.Output.SetPageHeadMeta(PageHeadMetaTags{
		Title:       "Vueghost",
		Description: "",
	})
	c.Output.SetStatusCode(200)
}

//Finalize
func (c *ApiController) Finalize() {
	apiOUTPUT := c.Output.Get()
	if c.httpResponseWriter != nil {
		c.httpResponseWriter.Header().Set("Content-Encoding", "gzip")
		gzipNewWriter := gzip.NewWriter(c.httpResponseWriter)
		jsonEncode := json.NewEncoder(gzipNewWriter)
		jsonEncode.SetEscapeHTML(true)
		_ = jsonEncode.Encode(&apiOUTPUT)
		_ = gzipNewWriter.Close()
		return
	}
	jsonEncode := json.NewEncoder(os.Stdout)
	_ = jsonEncode.Encode(apiOUTPUT)
}

//@throwError ApiController throw error handler.
func (c ApiController) throwError(err error) {
	_ = fmt.Errorf(
		"ApiController Error %v",
		err,
	)
}
