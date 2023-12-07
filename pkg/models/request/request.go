package request

type OtpSignupReq struct {
	FullName string `json:"fullname"`
	Phone    string  `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
}

type OtpReq struct {
	Phone string `json:"phone"`
}
