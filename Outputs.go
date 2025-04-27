package System

type (
	ApiOutput struct {
		store ApiOutputStore
	}
	ApiOutputStore struct {
		Success    bool        `json:"success"`
		Data       interface{} `json:"data,omitempty"`
		Message    interface{} `json:"message,omitempty"`
		Errors     interface{} `json:"errors,omitempty"`
		Meta       interface{} `json:"meta,omitempty"`
		StatusCode interface{} `json:"status,omitempty"`
		PagingInfo interface{} `json:"pagingInfo,omitempty"`
		IsAuth     bool        `json:"isAuth"`
	}
)

func NewApiOutput(store ApiOutputStore) *ApiOutput {
	return &ApiOutput{store: store}
}

//SetSuccess
func (c *ApiOutput) SetSuccess(success bool) {
	c.store.Success = success
}

//SetData
func (c *ApiOutput) SetData(data interface{}) {
	c.store.Data = data
}

//SetMessage
func (c *ApiOutput) SetMessage(message interface{}) {
	c.store.Message = message
}

//SetStatusCode
func (c *ApiOutput) SetStatusCode(statusCode interface{}) {
	c.store.StatusCode = statusCode
}

//SetPageHeadMeta
func (c *ApiOutput) SetPageHeadMeta(pageHeadMeta interface{}) {
	c.store.Meta = pageHeadMeta
}

//SetPagingInfo
func (c *ApiOutput) SetPagingInfo(pagingInfo interface{}) {
	c.store.PagingInfo = pagingInfo
}

//SetAuth
func (c *ApiOutput) SetAuth(isAuth bool) {
	c.store.IsAuth = isAuth
}

//SetErrors Set output errors.
func (c *ApiOutput) SetErrors(errors interface{}) {
	c.store.Errors = errors
}

func (c ApiOutput) Get() interface{} {
	return c.store
}
