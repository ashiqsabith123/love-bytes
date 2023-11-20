package functions

import (
	"context"
	"fmt"

	"github.com/ashiqsabith123/api-gateway/pkg/helper"
	"github.com/ashiqsabith123/api-gateway/pkg/models/request"
	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
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

func (A *AuthFunctions) SendOtp(data request.OtpReq) (responce.Response, bool) {

	if !helper.IsValidPhoneNumber(data.Phone) {

		response := helper.CreateResponse(400, "Phone number is not in proper format", "Invalid phone number")
		return response, false
	}

	resp, err := clients.SendOtp(context.TODO(), &pb.OtpReq{
		Phone: data.Phone,
	})

	fmt.Println("resp:", resp)

	response := helper.CreateResponse(resp.Code, resp.Message, resp.Error)

	if err != nil {
		return response, false
	}

	return response, true

}
