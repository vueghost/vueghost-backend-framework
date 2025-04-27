package System

import (
	"VGBackendFramework/Helpers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	router *mux.Router
}

func NewRouter() *Router {
	r := Router{}
	r.router = mux.NewRouter().StrictSlash(true)
	return &r
}

//ApiRoute
func (r *Router) ApiRoute(uri string, apiController interface{}, method string) {
	uriPath := r.getURI(uri)
	r.router.HandleFunc(uriPath, func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		Run(apiController, method, httpResponseWriter, httpRequest)
	})
}

//PageRoute
func (r *Router) PageRoute(uri string, page interface{}, method string) {
	uriPath := r.getURI(uri)
	r.router.HandleFunc(uriPath, func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		Run(page, method, httpResponseWriter, httpRequest)
	})
}

//ControllerRoute
func (r *Router) ControllerRoute(uri string, controller Controller, method string) {

}

//GetRouter
func (r *Router) GetRouter() *mux.Router {
	return r.router
}

func (r Router) Execute() *mux.Router {
	return r.router
}

func (r Router) getURI(uri string) string {
	uriMD5 := Helpers.EncodeMD5(uri)
	return fmt.Sprintf("/%s", uriMD5)
}
