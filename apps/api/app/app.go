package app

import (
	"github.com/urodstvo/messenger-server/libs/app"
	"github.com/urodstvo/messenger-server/libs/config"
	"github.com/urodstvo/messenger-server/libs/grpc/clients"
	"github.com/urodstvo/messenger-server/libs/logger"
	"go.uber.org/fx"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
)

const service = "API-Service"

var App = fx.Module(
	service,
	app.CreateBaseApp(service),
	fx.Provide(
		func(c config.Config) proto.ChatServiceClient {
			return clients.NewChatAppClient(c.EnvMode == "production")
		},
		func(c config.Config) proto.AuthServiceClient {
			return clients.NewAuthAppClient(c.EnvMode == "production")
		},
		//TODO
	),
	fx.Invoke(
		//TODO
		func(l logger.Logger) {
			l.Info(service + " started")
		},
	),
)
