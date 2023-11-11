package handler

import (
	"net/http"

	request "github.com/ashiqsabith123/api-gateway/pkg/models/request"
	responce "github.com/ashiqsabith123/api-gateway/pkg/models/responce"
	auth "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/functions/interface"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	functions auth.AuthFunctions
}

func NewAuthHandler(authFunc auth.AuthFunctions) *AuthHandler {
	return &AuthHandler{functions: authFunc}
}

func (A *AuthHandler) Signup(C *gin.Context) {

	var signupreq request.SignupReq

	if err := C.ShouldBindJSON(&signupreq); err != nil {
		resp := responce.ErrorReposonce(http.StatusNotAcceptable, "All fileds required", err)
		C.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
	}

	A.functions.SignUp(signupreq)

}
