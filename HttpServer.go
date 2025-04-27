package System

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

type (
	//HttpServer
	HttpServer struct {
		Settings   HttpServerSettings
		httpServer http.Server
	}
	//HttpServerSettings
	HttpServerSettings struct {
		Host           string
		Port           int
		Router         *Router
		CorsMiddleware map[string]string
		WriteTimeout   time.Duration
		ReadTimeout    time.Duration
		Environment    int
	}
)

//NewHttpServer
func NewHttpServer(settings HttpServerSettings) *HttpServer {
	return &HttpServer{
		Settings: settings,
		httpServer: http.Server{
			Addr:         fmt.Sprintf("%s:%v", settings.Host, settings.Port),
			WriteTimeout: settings.WriteTimeout,
			ReadTimeout:  settings.ReadTimeout,
		},
	}
}

//getHostAddress
func (s *HttpServer) getHostAddress() string {
	return fmt.Sprintf("%s:%v", s.Settings.Host, s.Settings.Port)
}

//setCorsMiddleware
func (s *HttpServer) setCorsMiddleware(httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		for key, value := range s.Settings.CorsMiddleware {
			httpResponseWriter.Header().Set(key, value)
		}
		if httpRequest.Method != "OPTIONS" {
			httpHandler.ServeHTTP(httpResponseWriter, httpRequest)
		}
	})
}

//Listen
func (s *HttpServer) Listen() {
	log.Print(ACIILogo)
	information := fmt.Sprintf(`
┌── SERVER INFOMRATION
│ › HTTP APIServer • Write Timeout:%[3]v • Read Timeout: %[4]v • IP4 Host: %[1]v • Port: %[2]v • Serving:http://%[1]v:%[2]v
| %[5]s • Number of CPU %[6]v • OS %[7]s`, s.Settings.Host, s.Settings.Port, s.Settings.WriteTimeout, s.Settings.ReadTimeout, runtime.Version(), runtime.NumCPU(), runtime.GOOS)
	log.Print(information)
	http.Handle("/", s.setCorsMiddleware(s.Settings.Router.Execute()))
	log.Fatal(s.httpServer.ListenAndServe())
}
