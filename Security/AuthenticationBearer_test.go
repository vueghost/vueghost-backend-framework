package Security

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticationBearer_GetToken(t *testing.T) {
	authBearer := AuthenticationBearer{}
	TestList := map[string]interface{}{
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJkYXRhIjoiMSIsImV4cCI6MTU4ODg4MjAxMTczNzIxOTAwMSwiaXNzIjoidnVlZ2hvc3QifQ.DuWl1Yd7sF1kCbFBgO6RyWU4YS7TfPJmYY4BXRg2jlw": 1,
		"": nil,
	}

	for token, _ := range TestList {
		httpTest := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, httpRequest *http.Request) {
			httpRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			authBearerToken := authBearer.GetToken(httpRequest)
			if authBearerToken != token {
				t.Fatal()
			}
		}))

		_, err := http.PostForm(httpTest.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		httpTest.Close()
	}

	httpTest := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, httpRequest *http.Request) {
		authBearerToken := authBearer.GetToken(httpRequest)
		if authBearerToken != "" {
			t.Fatal()
		}
	}))

	_, err := http.PostForm(httpTest.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	httpTest.Close()

}
