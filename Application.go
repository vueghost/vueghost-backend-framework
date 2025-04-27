package System

import (
	"time"
)

type (
	//Application
	Application struct {
		appServer     *HttpServer
		Configuration Configure
		Router        *Router
	}
	//ApplicationConfigure
	ApplicationConfigure struct {
		Name string
	}
)

//NewApplication
func NewApplication(config Configure, router *Router) *Application {
	return &Application{
		Configuration: config,
		Router:        router,
	}
}

//Run
func (a *Application) Run() {
	a.appServer = NewHttpServer(HttpServerSettings{
		Host:   a.Configuration.HttpServerHost,
		Port:   a.Configuration.HttpServerPort,
		Router: a.Router,
		CorsMiddleware: map[string]string{
			"Api-version":                      "1.0",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "GET,POST,OPTIONS,UPLOAD",
			"Access-Control-Allow-Headers":     "Origin, Content-Type, X-Auth-Token, X-Requested-With, Content-Type, Authorization",
			"Content-Type":                     "application/json; charset=utf-8",
		},
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Environment:  1,
	})
	a.appServer.Listen()
}
