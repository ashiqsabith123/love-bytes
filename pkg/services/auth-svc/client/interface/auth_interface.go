package interfaces

import "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/pb"

type AuthClient interface {
	InitAuthClient()
	GetClient() pb.AuthServiceClient
}
