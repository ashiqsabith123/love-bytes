package routes

import (
	"github.com/ashiqsabith123/api-gateway/pkg/api/handler"
	"github.com/ashiqsabith123/api-gateway/pkg/api/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(user *gin.RouterGroup, authHandler *handler.AuthHandler) {
	user.POST("/sendotp", authHandler.SendOtp)
	user.POST("/verifyotp", authHandler.VerifyOtpAndAuth)

	user.Use(middlewares.Authenticate)
	{
		user.POST("/details", authHandler.UserDetails)
	}

}
