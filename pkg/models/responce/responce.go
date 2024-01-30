package responce

type TokenResp struct {
	UserFound bool   ` json:"userfound"`
	Token     string ` json:"token,omitempty"`
}

type MatchedUsersResponse struct {
	MatchedUsers []*MatchedUsers `json:"matchedUsers,omitempty"`
}



type MatchedUsers struct {
	UserID     int32     `json:"userID,omitempty"`
	Name       string    `json:"name,omitempty"`
	Age        int32     `json:"age,omitempty"`
	Place      string    `json:"place,omitempty"`
	MatchScore int32     `json:"matchScore,omitempty"`
	UserImages []*Images `json:"userImages,omitempty"`
}

type Images struct {
	ImageId string `json:"imageId,omitempty"`
}

type Response struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorReposonce(code int, message string, err string) Response {

	return Response{
		Code:    code,
		Message: message,
		Error:   err,
	}
}

func SuccessResponse(code int, message string, data ...interface{}) Response {

	return Response{
		Code:    code,
		Message: message,
		Error:   nil,
		Data:    data,
	}
} // ignore: avoid_single_cascade_in_expression_statements
