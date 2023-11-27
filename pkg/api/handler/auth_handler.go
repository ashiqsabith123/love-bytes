package handler

import (
	"net/http"

	request "github.com/ashiqsabith123/api-gateway/pkg/models/request"
	responce "github.com/ashiqsabith123/api-gateway/pkg/models/responce"
	auth "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/functions/interfaces"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	functions auth.AuthFunctions
}

func NewAuthHandler(authFunc auth.AuthFunctions) *AuthHandler {
	return &AuthHandler{functions: authFunc}
}

func (A *AuthHandler) VerifyOtpAndSignup(C *gin.Context) {

	var otpsignupreq request.OtpSignupReq

	if err := C.ShouldBindJSON(&otpsignupreq); err != nil {
		resp := responce.ErrorReposonce(http.StatusNotAcceptable, "All fileds required", err.Error(), nil)
		C.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
	}

	resp, ok := A.functions.VerifyOtpAndSignUp(otpsignupreq)

	if !ok {

		resp := responce.ErrorReposonce(resp.Code, resp.Message, resp.Error.(string), nil)
		C.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp = responce.SuccessResponse(resp.Code, resp.Message)
	C.JSON(resp.Code, resp)

}

func (A *AuthHandler) SendOtp(C *gin.Context) {
	var otpReq request.OtpReq

	if err := C.ShouldBindJSON(&otpReq); err != nil {
		resp := responce.ErrorReposonce(http.StatusBadRequest, "Invalid number", err.Error(), nil)
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp, ok := A.functions.SendOtp(otpReq)

	if !ok {

		resp := responce.ErrorReposonce(resp.Code, resp.Message, resp.Error.(string), nil)
		C.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	resp = responce.SuccessResponse(resp.Code, resp.Message)
	C.JSON(resp.Code, resp)

}
