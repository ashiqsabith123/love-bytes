package server

import (
	"github.com/ashiqsabith123/api-gateway/pkg/api/handler"
	"github.com/ashiqsabith123/api-gateway/pkg/api/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/ashiqsabith123/api-gateway/cmd/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine *gin.Engine
}

// @title Love Bites API
// @version 1.0
// @description Fully fuctional dating app API
// @contact ashiqsabith328@gmail.com
// @host localhost:8081/user
// @securityDefinitions.apikey	BearerTokenAuth
// @in							header
// @name						Authorization
func NewServer(authHandler *handler.AuthHandler, matchHanlder *handler.MatchHandler, notificationHandler *handler.NotificationHandler) *Server {

	server := gin.Default()

	user := server.Group("/user")

	// Add Swagger route
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.AuthRoutes(user, authHandler)
	routes.MatchRoutes(user, matchHanlder)
	routes.NotificationRoutes(user, notificationHandler)

	return &Server{
		engine: server,
	}

}

func (s *Server) Start() {

	s.engine.Run(":8081")
}
