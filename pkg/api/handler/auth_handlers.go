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

func (A *AuthHandler) SendOtp(C *gin.Context) {

	var otpReq request.OtpReq

	if err := C.ShouldBindJSON(&otpReq); err != nil {
		resp := responce.ErrorReposonce(http.StatusBadRequest, "Invalid number", err.Error())
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp, ok := A.functions.SendOtp(otpReq)

	if !ok {

		C.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	C.JSON(resp.Code, resp)

}

func (A *AuthHandler) VerifyOtpAndAuth(C *gin.Context) {
	var verifyOtpReq request.VerifyOtpReq

	if err := C.ShouldBindJSON(&verifyOtpReq); err != nil {
		resp := responce.ErrorReposonce(http.StatusBadRequest, "Invalid request", err.Error())
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp, ok := A.functions.VerifyOtpAndAuth(verifyOtpReq)

	if !ok {
		C.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	C.JSON(resp.Code, resp)
}

func (A *AuthHandler) UserDetails(C *gin.Context) {

	_, ok := C.Get("userID")

	if !ok {
		resp := responce.ErrorReposonce(http.StatusBadRequest, "Invalid request", "User id not found")
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var userDetails request.UserDetails

	if err := C.ShouldBindJSON(&userDetails); err != nil {
		resp := responce.ErrorReposonce(http.StatusBadRequest, "Invalid request", err.Error())
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp, ok := A.functions.SaveUserDetails(C, userDetails)

	if !ok {
		C.AbortWithStatusJSON(resp.Code, resp)
		return
	}

	C.JSON(resp.Code, resp)

}
