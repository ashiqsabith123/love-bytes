package functions

import (
	"context"

	"github.com/ashiqsabith123/api-gateway/pkg/models/request"
	client "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/client/interface"
	auth "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/functions/interfaces"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
)

type AuthFunctions struct {
	client client.AuthClient
}

func NewAuthFunctions(client client.AuthClient) auth.AuthFunctions {
	return &AuthFunctions{client: client}

}

func (A *AuthFunctions) SignUp(data request.SignupReq) {

	client := A.client.GetClient()

	client.Signup(context.TODO(), &pb.SignUpReq{
		Fullname: data.FullName,
		Phone:    data.Phone,
		Username: data.Username,
		Password: data.Password,
	})
}
