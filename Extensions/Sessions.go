package Extensions

//Session the user session token.
type Session struct {
	ID interface{} `json:"id"`
}

//NewSession session construct.
func NewSession(ID interface{}) *Session {
	return &Session{ID: ID}
}

//Set set new session id.
func (s *Session) Set(sessionID interface{}) {
	s.ID = sessionID
}
