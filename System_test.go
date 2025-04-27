package System

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExampleRun() {
	httpTest := httptest.NewServer(http.HandlerFunc(func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		Run(ApiController{}, "AboutController", httpResponseWriter, httpRequest)
	}))
	defer httpTest.Close()
}

func TestRun(t *testing.T) {
	ExampleRun()
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExampleRun()
	}
}
