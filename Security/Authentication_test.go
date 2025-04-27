package Security

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthentication_Context(t *testing.T) {
	t.Run("When Context with HttpRequest", func(t *testing.T) {
		class := Authentication{}

		httpTest := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, httpRequest *http.Request) {
			httpRequest.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJkYXRhIjoiMSIsImV4cCI6MTU4ODg4MjAxMTczNzIxOTAwMSwiaXNzIjoidnVlZ2hvc3QifQ.DuWl1Yd7sF1kCbFBgO6RyWU4YS7TfPJmYY4BXRg2jlw")
			r := class.Context(httpRequest)
			fmt.Println(r)

		}))

		_, err := http.PostForm(httpTest.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		httpTest.Close()
	})
}

func TestAuthentication_SetAuthToken(t *testing.T) {
	m := Authentication{}

	t.Run("When userID is nil value", func(t *testing.T) {
		authToken := m.SetAuthToken(nil)
		if authToken != "" {
			t.Fatal()
		}
	})
	t.Run("When userID is int value as 1", func(t *testing.T) {
		authToken := m.SetAuthToken(1)
		if authToken == "" {
			t.Fatal()
		}
	})
	t.Run("When userID is string as  1", func(t *testing.T) {
		authToken := m.SetAuthToken("1")
		if authToken == "" {
			t.Fatal()
		}
	})
	t.Run("When userID has err", func(t *testing.T) {
		authToken := m.SetAuthToken("@#092@!1=38")
		if authToken == "" {
			t.Fatal()
		}
	})
}

func TestAuthentication_GetAuthToken(t *testing.T) {
	class := Authentication{}

	type getAuthTokenTest struct {
		value  string
		result interface{}
	}

	currentTestList := map[string]getAuthTokenTest{
		"when is valid token": {
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJkYXRhIjoiMSIsImV4cCI6MTYwMjY5MTAyNTEzODYxNTAwMSwiaXNzIjoidnVlZ2hvc3QifQ.ITeVaHwGs7qc2I6tcp8LZXvR0HT8xNsRbINHqerhJa0",
			"1",
		},
		"when is not valid token": {
			"eyJhbGciOiJI2UzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJkYXRhIjoiMSIsImV4cCI6MTYwMjY5MTAyNTEzODYxNTAwMSwiaXNzIjoidnVlZ2hvc3QifQ.ITeVaHwGs7qc2I6tcp8LZXvR0HT8xNsRbINHqerhJa0",
			"0",
		},
	}

	for name, test := range currentTestList {
		t.Run(name, func(t *testing.T) {
			result := class.GetAuthToken(test.value)
			if result != test.result {
				t.Fatal()
			}
		})
	}
}
