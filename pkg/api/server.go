package server

import (
	"github.com/ashiqsabith123/api-gateway/pkg/api/handler"
	"github.com/ashiqsabith123/api-gateway/pkg/api/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(authHandler *handler.AuthHandler, matchHanlder *handler.MatchHandler) *Server {

	server := gin.Default()

	user := server.Group("/user")

	routes.AuthRoutes(user, authHandler)
	routes.MatchRoutes(user, matchHanlder)

	return &Server{
		engine: server,
	}

}

func (s *Server) Start() {

	s.engine.Run(":8081")
}
