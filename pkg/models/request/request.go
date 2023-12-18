package request

type VerifyOtpReq struct {
	Phone string `json:"phone" validate:"required"`
	Otp   string `json:"otp" validate:"required,min=6"`
}

type OtpReq struct {
	Phone string `json:"phone"`
}

type UserDetails struct {
	Fullname    string `json:"fullname" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Dateofbirth string `json:"dob" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
}
