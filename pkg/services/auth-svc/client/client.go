package client

import (
	"github.com/ashiqsabith123/api-gateway/pkg/config"
	intefaces "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/client/interface"

	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
	logs "github.com/ashiqsabith123/love-bytes-proto/log"
	"google.golang.org/grpc"
)

type AuthClient struct {
	config config.Config
}

var Conn *grpc.ClientConn
var err error

func NewAuthClient(config config.Config) intefaces.AuthClient {
	client := &AuthClient{config: config}
	client.InitAuthClient()
	return &AuthClient{}
}

func (A *AuthClient) InitAuthClient() {

	// credentials, err := helper.GetCertificate("pkg/services/auth-svc/cert/ca-cert.pem", "pkg/services/auth-svc/cert/client-cert.pem", "pkg/services/auth-svc/cert/client-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	Conn, err = grpc.Dial(A.config.PORTS.AuthSvcPort, grpc.WithInsecure())
	if err != nil {
		logs.ErrLog.Fatal("Could not connect the auth server:", err)
	}

	logs.GenLog.Println("Auth service connected at port ", A.config.PORTS.AuthSvcPort)

}

func (A *AuthClient) GetClient() pb.AuthServiceClient {
	return pb.NewAuthServiceClient(Conn)
}
