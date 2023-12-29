package routes

import (
	"github.com/ashiqsabith123/api-gateway/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func MatchRoutes(user *gin.RouterGroup, matchHandler *handler.MatchHandler) {
	user.POST("/upload/photos", matchHandler.UploadPhotos)
	user.POST("/save/prefrences", matchHandler.SaveUserPrefrences)
}
