package functions

import (
	"fmt"

	"github.com/ashiqsabith123/api-gateway/pkg/models/request"
	auth "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/functions/interface"
)

type AuthFunctions struct {
}

func NewAuthFunctions() auth.AuthFunctions {
	return &AuthFunctions{}

}

func (A *AuthFunctions) SignUp(data request.SignupReq) {

	fmt.Println(data)

}
