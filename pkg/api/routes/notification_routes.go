package routes

import (
	"github.com/ashiqsabith123/api-gateway/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func NotificationRoutes(user *gin.RouterGroup, notificationHandler *handler.NotificationHandler) {

	notification := user.Group("/notification")
	{
		notification.POST("/all", notificationHandler.GetAllNotifications)
		notification.POST("/save/fcm",)
	}
}
