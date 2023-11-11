//go:build ignore
// +build ignore

package di

import (
	server "github.com/ashiqsabith123/api-gateway/pkg/api"
	"github.com/ashiqsabith123/api-gateway/pkg/api/handler"
	auth "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/functions"
	"github.com/google/wire"
)

func InitializeApi() *server.Server {
	wire.Build(
		auth.NewAuthFunctions,
		handler.NewAuthHandler,
		server.NewServer,
	)

	return &server.Server{}
}
