package responce

type ErrorResp struct {
	Code    int
	Message string
	Error   error
}

func ErrorReposonce(code int, message string, err error) ErrorResp {

	return ErrorResp{
		Code:    code,
		Message: message,
		Error:   err,
	}
}
