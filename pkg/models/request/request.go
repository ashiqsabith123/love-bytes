package request

type SignupReq struct {
	FullName string `json:"fullname"`
	Phone    int64  `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type OtpReq struct {
	Phone string `json:"phone"`
}
