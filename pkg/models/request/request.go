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

type UserPreferences struct {
	Height        string `json:"height" validate:"required"`
	MaritalStatus string `json:"marital_status" validate:"required"`
	Faith         string `json:"faith" validate:"required"`
	MotherTongue  string `json:"mother_tongue" validate:"required"`
	SmokeStatus   string `json:"smoke_status" validate:"required"`
	AlcoholStatus string `json:"alcohol_status" validate:"required"`
	SettleStatus  string `json:"settle_status" validate:"required"`
	Hobbies       string `json:"hobbies" validate:"required"`
	TeaPerson     string `json:"tea_person" validate:"required"`
	LoveLanguage  string `json:"love_language" validate:"required"`
}
