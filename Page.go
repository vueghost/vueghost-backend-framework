package System

import (
	"compress/gzip"
	"fmt"
	"net/http"
)

type (
	Page struct {
		Controller
		view View
	}
	PageHeadMetaTags struct {
		Title         interface{} `db:"title" json:"title,omitempty"`
		Description   interface{} `db:"description" json:"description,omitempty"`
		Keywords      interface{} `db:"keywords" json:"keywords,omitempty"`
		OGImage       interface{} `json:"ogImage,omitempty"`
		OGUrl         interface{} `json:"ogUrl,omitempty"`
		OGType        interface{} `json:"ogType,omitempty"`
		OGTitle       interface{} `json:"ogTitle,omitempty"`
		OGDescription interface{} `json:"ogDescription,omitempty"`
		VGID          interface{} `json:"vgID,omitempty"`
		VGType        interface{} `json:"vgType,omitempty"`
	}
)

func (p *Page) Context(HttpResponseWriter http.ResponseWriter, HttpRequest *http.Request) {
	HttpResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.httpRequest = HttpRequest
	p.httpResponseWriter = HttpResponseWriter
	p.HttpRequestHeader.Context(p.httpRequest)

	//@Inputs & outputs.
	p.Input.Context(HttpResponseWriter, HttpRequest)
}

func (p *Page) View(viewName string, data interface{}) {
	p.view.View(viewName, data)
}

func (p *Page) ViewRender(viewName string, data interface{}) string {
	return p.view.ViewRender(viewName, data)
}

func (p *Page) Finalize() {
	p.httpResponseWriter.Header().Set("Content-Encoding", "gzip")
	gzipNewWriter := gzip.NewWriter(p.httpResponseWriter)
	_ = p.view.viewTemplate.Execute(gzipNewWriter, p.view.data)
	_ = gzipNewWriter.Close()
}

//throwError error exception of controller handler.
func (p Page) throwError(err error) {
	_ = fmt.Errorf(
		"Page error %v",
		err,
	)
}
