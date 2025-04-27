package Extensions

import "time"

//Timestamp
type Timestamp struct {
}

//NewTimestamp timestamp new construct.
func NewTimestamp() *Timestamp {
	return &Timestamp{}
}

//Now Current timestamp yyyy-mm-dd h:m:s default format.
func (t Timestamp) Now(Format ...string) interface{} {
	timeDate := time.Now()
	FormatLayout := "2006-01-02 15:04:05"
	if len(Format) > 0 {
		FormatLayout = Format[0]
	}
	return timeDate.Format(FormatLayout)
}
