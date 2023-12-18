package intefaces

import (
	"context"

	"github.com/ashiqsabith123/api-gateway/pkg/models/request"
	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
)

type AuthFunctions interface {
	SendOtp(data request.OtpReq) (responce.Response, bool)
	VerifyOtpAndAuth(data request.VerifyOtpReq) (responce.Response, bool)
	SaveUserDetails(ctx context.Context, userDetails request.UserDetails) (responce.Response, bool)
}
