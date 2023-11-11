//go:build wireinject
// +build wireinject

package di

import (
	server "github.com/ashiqsabith123/api-gateway/pkg/api"
	"github.com/ashiqsabith123/api-gateway/pkg/api/handler"
	"github.com/ashiqsabith123/api-gateway/pkg/config"
	client "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/client"
	auth "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/functions"
	"github.com/google/wire"
)

func InitializeApi(config config.Config) *server.Server {
	wire.Build(
		client.NewAuthClient,
		auth.NewAuthFunctions,
		handler.NewAuthHandler,
		server.NewServer,
	)

	return &server.Server{}
}
