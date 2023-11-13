package functions

import (
	"context"
	"fmt"

	"github.com/ashiqsabith123/api-gateway/pkg/models/request"
	client "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/client/interface"
	auth "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/functions/interfaces"
	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
)

type AuthFunctions struct {
	client client.AuthClient
}

var clients pb.AuthServiceClient

func NewAuthFunctions(client client.AuthClient) auth.AuthFunctions {
	authfunc := &AuthFunctions{client: client}
	clients = authfunc.client.GetClient()
	return &AuthFunctions{}

}

func (A *AuthFunctions) SignUp(data request.SignupReq) {

	resp, err := clients.Signup(context.TODO(), &pb.SignUpReq{
		Fullname: data.FullName,
		Phone:    data.Phone,
		Username: data.Username,
		Password: data.Password,
	})

	fmt.Println(resp, err)
}
