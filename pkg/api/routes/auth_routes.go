package routes

import (
	"github.com/ashiqsabith123/api-gateway/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(server *gin.Engine, authHandler *handler.AuthHandler) {

	user := server.Group("/user")
	{
		user.POST("/signup", authHandler.Signup)
	}

}
