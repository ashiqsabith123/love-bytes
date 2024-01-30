package client

import (
	logs "github.com/ashiqsabith123/love-bytes-proto/log"

	"github.com/ashiqsabith123/api-gateway/pkg/config"
	interfaces "github.com/ashiqsabith123/api-gateway/pkg/services/notification-svc/client/interface"
	"github.com/ashiqsabith123/love-bytes-proto/notifications/pb"
	"google.golang.org/grpc"
)

type NotificationClient struct {
	config config.Config
}

var Conn *grpc.ClientConn
var err error

func NewNotificationClient(config config.Config) interfaces.NotificationClient {
	client := &NotificationClient{config: config}
	client.InitNotificationClient()
	return &NotificationClient{}
}

func (A *NotificationClient) InitNotificationClient() {

	// credentials, err := helper.GetCertificate("pkg/services/auth-svc/cert/ca-cert.pem", "pkg/services/auth-svc/cert/client-cert.pem", "pkg/services/auth-svc/cert/client-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	Conn, err = grpc.Dial(A.config.PORTS.NotificationSvcPort, grpc.WithInsecure())
	if err != nil {
		logs.ErrLog.Println("Could not connect the notification server:", err)
	}

	logs.GenLog.Println("Notification service connected at port ", A.config.PORTS.NotificationSvcPort)

}

func (A *NotificationClient) GetClient() pb.NotificationServiceClient {
	return pb.NewNotificationServiceClient(Conn)
}
