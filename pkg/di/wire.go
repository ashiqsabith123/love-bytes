//go:build wireinject
// +build wireinject

package di

import (
	server "github.com/ashiqsabith123/api-gateway/pkg/api"
	"github.com/ashiqsabith123/api-gateway/pkg/api/handler"
	"github.com/ashiqsabith123/api-gateway/pkg/config"
	authclient "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/client"
	auth "github.com/ashiqsabith123/api-gateway/pkg/services/auth-svc/functions"
	matchclient "github.com/ashiqsabith123/api-gateway/pkg/services/match-svc/client"
	match "github.com/ashiqsabith123/api-gateway/pkg/services/match-svc/functions"
	"github.com/google/wire"
)

func InitializeApi(config config.Config) *server.Server {
	wire.Build(
		authclient.NewAuthClient,
		auth.NewAuthFunctions,
		handler.NewAuthHandler,
		matchclient.NewMatchClient,
		match.NewMatchFunctions,
		handler.NewMatchHandler,
		server.NewServer,
	)

	return &server.Server{}
}
