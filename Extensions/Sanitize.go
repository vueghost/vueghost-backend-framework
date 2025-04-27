package Extensions

import "html/template"

//Sanitize
type Sanitize struct {
}

//NewSanitize
func NewSanitize() *Sanitize {
	return &Sanitize{}
}

//URLQueryEscape does the inverse transformation of QueryEscape, converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits.
func (S Sanitize) URLQueryEscape(input interface{}) interface{} {
	return template.URLQueryEscaper(input)
}
