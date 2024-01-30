package handler

import (
	"errors"
	"net/http"

	responce "github.com/ashiqsabith123/api-gateway/pkg/models/responce"
	notification "github.com/ashiqsabith123/api-gateway/pkg/services/notification-svc/functions/interface"
	"github.com/ashiqsabith123/api-gateway/pkg/utils/graphql/models"
	Const "github.com/ashiqsabith123/love-bytes-proto/constants"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type NotificationHandler struct {
	functions notification.NotificationFunctions
}

func NewNotificationHandler(mathcFunc notification.NotificationFunctions) *NotificationHandler {
	return &NotificationHandler{functions: mathcFunc}
}

func (N *NotificationHandler) GetAllNotifications(C *gin.Context) {

	_, ok := C.Get("userID")

	if !ok {
		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, Const.USER_ID_NOT_FOUND)
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var NotificationQueryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Notificationquery",
		Fields: graphql.Fields{
			"notifications": &graphql.Field{
				Type: graphql.NewList(models.NotificationType),
				Args: graphql.FieldConfigArgument{
					"day": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					day, _ := p.Args["day"].(string)

					notificationType, _ := p.Args["type"].(string)

					resp, ok := N.functions.GetAllNotifications(C, day, notificationType)

					if !ok {
						return nil, errors.New("server error retuned nill data")
					}

					return resp.Data, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: NotificationQueryType,
	})

	if err != nil {
		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, err.Error())
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	handler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	handler.ServeHTTP(C.Writer, C.Request)
}

func (N *NotificationHandler) SaveFCMToken(C *gin.Context) {
	_, ok := C.Get("userID")

	if !ok {
		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, Const.USER_ID_NOT_FOUND)
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var token string
	
	if err := C.ShouldBindJSON(&token); err != nil {
		resp := responce.ErrorReposonce(http.StatusBadRequest, Const.INVALID_REQUEST, err.Error())
		C.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

}
