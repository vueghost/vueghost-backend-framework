package Security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type (
	//JwtToken is an abbreviation for JSON Web Token, which is a compact URL-safe means of representing claims to be transferred between two parties. The claims in a JWT are encoded as a JSON object that is digitally signed using JSON Web Signature (JWS).
	JwtToken      struct{}
	JwtTokenClaim struct {
		Foo  string `json:"foo,omitempty"`
		Data string `json:"data,omitempty"`
		jwt.StandardClaims
	}
)

//Get decode Jwt token data.
func (j JwtToken) Get(token string) (data interface{}, success bool) {
	mapClaims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, mapClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignKey), nil
	})

	// Check Validation of the token.
	if err == nil && jwtToken.Valid {
		return mapClaims["data"], true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "hat's not even a token", false
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return "token is expired", false
		} else {
			return fmt.Sprintf("Couldn't handle this token: %s", err), false
		}
	} else {
		return fmt.Sprintf("Couldn't handle this token: %s", err), false
	}
}

//Set encode data to Jwt token.
func (j JwtToken) Set(data string) (string, error) {
	privateSigningKey := []byte(SignKey)

	mTokenClaimsStruct := JwtTokenClaim{
		"bar",
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1).UnixNano(),
			Issuer:    "vueghost",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mTokenClaimsStruct)
	return token.SignedString(privateSigningKey)
}
