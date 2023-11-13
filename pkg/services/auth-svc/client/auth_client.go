package client

import (
	"fmt"
	"log"

	"github.com/ashiqsabith123/api-gateway/pkg/config"
	"github.com/ashiqsabith123/api-gateway/pkg/helper"
	intefaces "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/client/interface"

	"github.com/ashiqsabith123/love-bytes-proto/auth/pb"
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

	credentials := helper.GetCertificate("pkg/services/auth-svc/cert/ca-cert.pem", "pkg/services/auth-svc/cert/client-cert.pem", "pkg/services/auth-svc/cert/client-key.pem")
	Conn, err = grpc.Dial(A.config.PORTS.AuthSvcPort, grpc.WithTransportCredentials(credentials))
	if err != nil {
		log.Fatal("Could not connect the auth server:", err)
	}

	fmt.Println("Auth service connected at port ", A.config.PORTS.AuthSvcPort)

}

func (A *AuthClient) GetClient() pb.AuthServiceClient {
	return pb.NewAuthServiceClient(Conn)
}
