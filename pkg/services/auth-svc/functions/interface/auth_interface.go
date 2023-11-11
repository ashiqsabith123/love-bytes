package inteface

import "github.com/ashiqsabith123/api-gateway/pkg/models/request"

type AuthFunctions interface {
	SignUp(data request.SignupReq)
}
