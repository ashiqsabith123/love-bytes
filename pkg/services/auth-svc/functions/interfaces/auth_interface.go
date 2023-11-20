package intefaces

import (
	"github.com/ashiqsabith123/api-gateway/pkg/models/request"
	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
)

type AuthFunctions interface {
	SignUp(data request.SignupReq)
	SendOtp(data request.OtpReq) (responce.Response, bool)
}
