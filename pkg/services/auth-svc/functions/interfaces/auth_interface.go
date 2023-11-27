package intefaces

import (
	"github.com/ashiqsabith123/api-gateway/pkg/models/request"
	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
)

type AuthFunctions interface {
	VerifyOtpAndSignUp(data request.OtpSignupReq) (responce.Response, bool)
	SendOtp(data request.OtpReq) (responce.Response, bool)
}
