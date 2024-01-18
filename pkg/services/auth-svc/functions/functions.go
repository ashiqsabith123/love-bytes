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
	"github.com/jinzhu/copier"

	"google.golang.org/protobuf/proto"
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

func (A *AuthFunctions) SendOtp(otp request.OtpReq) (responce.Response, bool) {

	if !helper.IsValidPhoneNumber(otp.Phone) {
		response := helper.CreateResponse(400, "Phone number is not in proper format", "Invalid phone number", nil)
		return response, false
	}

	resp, _ := clients.SendOtp(context.TODO(), &pb.OtpReq{
		Phone: otp.Phone,
	})

	if resp.Error != nil {
		response := helper.CreateResponse(resp.Code, resp.Message, string(resp.Error.Value), nil)
		return response, false
	}

	response := helper.CreateResponse(resp.Code, resp.Message, nil, nil)

	return response, true

}

func (A *AuthFunctions) VerifyOtpAndAuth(verifyOtp request.VerifyOtpReq) (responce.Response, bool) {

	if !helper.IsValidPhoneNumber(verifyOtp.Phone) {
		response := helper.CreateResponse(400, "Phone number is not in proper format", "Invalid phone number", nil)
		return response, false
	}

	err := helper.Validator(verifyOtp)

	if err != nil {
		response := helper.CreateResponse(400, "Data is not in proper format", "Invalid fields", nil)
		return response, false
	}

	resp, err := clients.VerifyOtpAndAuth(context.TODO(), &pb.VerifyOtpReq{
		Phone: verifyOtp.Phone,
		Otp:   verifyOtp.Otp,
	})

	if err != nil {
		fmt.Println("Errr", err)
	}

	if resp != nil {
		if resp.Error != nil {
			response := helper.CreateResponse(resp.Code, resp.Message, string(resp.Error.Value), nil)
			return response, false
		}
	}

	var tokenData pb.TokenResp

	if err := proto.Unmarshal(resp.Data.Value, &tokenData); err != nil {
		response := helper.CreateResponse(resp.Code, resp.Message, "Error unmarshaling data", nil)
		return response, false
	}

	var TokenResp responce.TokenResp

	copier.Copy(&TokenResp, &tokenData)

	response := helper.CreateResponse(resp.Code, resp.Message, nil, TokenResp)

	return response, true
}

func (A *AuthFunctions) SaveUserDetails(ctx context.Context, userDetails request.UserDetails) (responce.Response, bool) {

	err := helper.Validator(userDetails)

	if err != nil {
		response := helper.CreateResponse(400, "Data is not in proper format", "Invalid fields"+err.Error(), nil)
		return response, false
	}

	resp, _ := clients.SaveUserDetais(ctx, &pb.UserDetailsReq{
		UserID:      helper.GetUserID(ctx),
		Fullname:    userDetails.Fullname,
		Email:       userDetails.Email,
		Location:    userDetails.Location,
		Dateofbirth: userDetails.Dateofbirth,
		Gender:      userDetails.Gender,
	})

	if resp.Error != nil {
		response := helper.CreateResponse(resp.Code, resp.Message, string(resp.Error.Value), nil)
		return response, false
	}

	response := helper.CreateResponse(resp.Code, resp.Message, nil, nil)

	return response, true

	
}
