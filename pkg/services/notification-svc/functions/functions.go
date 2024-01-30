package functions

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/ashiqsabith123/api-gateway/pkg/helper"
	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
	client "github.com/ashiqsabith123/api-gateway/pkg/services/notification-svc/client/interface"
	interfaces "github.com/ashiqsabith123/api-gateway/pkg/services/notification-svc/functions/interface"
	"github.com/ashiqsabith123/api-gateway/pkg/utils/graphql/models"
	"github.com/ashiqsabith123/love-bytes-proto/notifications/pb"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
)

type NotificationFunctions struct {
	client client.NotificationClient
}

var clients pb.NotificationServiceClient

func NewNotificationFunctions(client client.NotificationClient) interfaces.NotificationFunctions {
	notifiFunc := &NotificationFunctions{client: client}
	clients = notifiFunc.client.GetClient()
	return &NotificationFunctions{}

}

func (N *NotificationFunctions) GetAllNotifications(ctx context.Context, filter ...string) (responce.Response, bool) {
	var resp *pb.NotificationResponce

	for i := 0; i < 5; i++ {
		resp, _ = clients.GetAllNotifiacation(context.Background(), &pb.GetNotificationRequest{
			UserID: uint32(helper.GetUserID(ctx)),
			Day:    filter[0],
			Type:   filter[1],
		})

		if resp != nil {
			break

		}

		time.Sleep(1 * time.Second)
	}

	if resp != nil {
		if resp.Error != nil {
			response := helper.CreateResponse(resp.Code, resp.Message, string(resp.Error.Value), nil)
			return response, false
		}

		var notifications pb.AllNotifications

		if err := proto.Unmarshal(resp.Data.Value, &notifications); err != nil {
			response := helper.CreateResponse(resp.Code, resp.Message, "Error unmarshaling data", nil)
			return response, false
		}

		var notifi []models.Notifications

		copier.Copy(&notifi, &notifications.Allnotification)

		response := helper.CreateResponse(resp.Code, resp.Message, nil, notifi)
		return response, true
	}

	response := helper.CreateResponse(http.StatusInternalServerError, "Service retunrned nill", errors.New("server error").Error(), nil)

	return response, false

}


func (N *NotificationFunctions)SaveFCMTokens(ctx context.Context, token string){

}
