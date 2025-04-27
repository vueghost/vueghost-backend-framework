package Security

import (
	"Framework/Helpers"
	"fmt"

	"net/http"
)

// Authentication class
type Authentication struct {
	authSessionToken string
	authSessionID    int
	authKeyCode      []byte
	authBearer       AuthenticationBearer
	jwtToken         JwtToken
}

//Context the context function.
func (s Authentication) Context(HttpRequest *http.Request) interface{} {
	s.authKeyCode = []byte("vueghost-com-key")

	bearerToken := s.authBearer.GetToken(HttpRequest)
	fmt.Println(bearerToken)
	if !Helpers.IsEmpty(bearerToken) {
		return s.GetAuthToken(bearerToken)
	}
	return "0"
}

//@pubic
//SetAuthToken Create authentication token.
func (s Authentication) SetAuthToken(userID interface{}) string {
	if userID == nil {
		return ""
	}
	formattedUserID := fmt.Sprintf("%v", userID)
	authToken, _ := s.jwtToken.Set(formattedUserID)
	return authToken
}

//GetAuthToken Return authentication token data.
func (s Authentication) GetAuthToken(authToken string) interface{} {
	if jwtTokenData, success := s.jwtToken.Get(authToken); success {
		return jwtTokenData
	}
	return "0"
}
