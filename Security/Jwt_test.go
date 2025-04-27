package Security

import "testing"

var JwtTokens = map[string]bool{
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJkYXRhIjoiMSIsImV4cCI6MTU4ODg4MjAxMTczNzIxOTAwMSwiaXNzIjoidnVlZ2hvc3QifQ.DuWl1Yd7sF1kCbFBgO6RyWU4YS7TfPJmYY4BXRg2jlw": true,
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJkYXRhIjoiMSIsImV4cCI6MTU4ODg4MjAxMTczNzIxOTAwMSwiaXNzIjoidnVlZ2hvc3QifQ.DuWl1Yd7sF1kCbFBgO6RyWU4YS7TfPJmYY4BXRg2jl2": false,
}

func TestJwtToken_Get(t *testing.T) {
	a := JwtToken{}

	for token, r := range JwtTokens {
		data, result := a.Get(token)
		if r == result {
			t.Log(data)
		} else {
			t.Fatalf("Must be %v and return %v", r, result)
		}
	}
}
