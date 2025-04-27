package System

type ErrorHandler struct {
	Message        string      `db:"message" json:"message,omitempty"`
	Type           string      `db:"type" json:"type,omitempty"`
	Code           interface{} `db:"code" json:"code,omitempty"`
	ErrorSubcode   interface{} `db:"error_subcode" json:"errorSubcode,omitempty"`
	ErrorUserTitle interface{} `db:"error_user_title" json:"errorUserTitle,omitempty"`
	ErrorUserMsg   interface{} `db:"error_user_msg" json:"errorUserMsg,omitempty"`
}
