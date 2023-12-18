package client

import (
	"fmt"
	"log"

	"github.com/ashiqsabith123/api-gateway/pkg/config"
	intefaces "github.com/ashiqsabith123/api-gateway/pkg/services/match-svc/client/interface"

	"github.com/ashiqsabith123/love-bytes-proto/match/pb"
	"google.golang.org/grpc"
)

type MatchClient struct {
	config config.Config
}

var Conn *grpc.ClientConn
var err error

func NewMatchClient(config config.Config) intefaces.MatchClient {
	client := &MatchClient{config: config}
	client.InitMatchClient()
	return &MatchClient{}
}

func (A *MatchClient) InitMatchClient() {

	// credentials, err := helper.GetCertificate("pkg/services/auth-svc/cert/ca-cert.pem", "pkg/services/auth-svc/cert/client-cert.pem", "pkg/services/auth-svc/cert/client-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	Conn, err = grpc.Dial(A.config.PORTS.MatchSvcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect the auth server:", err)
	}

	fmt.Println("Match service connected at port ", A.config.PORTS.MatchSvcPort)

}

func (A *MatchClient) GetClient() pb.MatchServiceClient {
	return pb.NewMatchServiceClient(Conn)
}
