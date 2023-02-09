package common

type Response struct {
	HasError int         `json:"hasError"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

func NewResponse(hasError int, message string, data interface{}) *Response {
	return &Response{
		HasError: hasError,
		Message:  message,
		Data:     data,
	}
}
