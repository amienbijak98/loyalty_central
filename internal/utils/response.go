package utils

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewSuccessResponse(data interface{}) *Response {
	return NewResponse(200, "Success", data)
}

func NewErrorResponse(err error) *Response {
	return NewResponse(500, err.Error(), nil)
}
