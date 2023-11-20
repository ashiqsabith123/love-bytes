package responce

type Response struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorReposonce(code int, message string, err string, data interface{}) Response {

	return Response{
		Code:    code,
		Message: message,
		Error:   err,
		Data:    data,
	}
}

func SuccessResponse(code int, message string, data ...interface{}) Response {

	return Response{
		Code:    code,
		Message: message,
		Error:   nil,
		Data:    data,
	}
}
