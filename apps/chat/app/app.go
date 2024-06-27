package app

import (
	"net/http"

	implementation "github.com/urodstvo/messenger-server/apps/chat/internal/api/v1/grpc_impl"
	"github.com/urodstvo/messenger-server/apps/chat/internal/api/v1/websocket"
	"github.com/urodstvo/messenger-server/libs/app"
	"github.com/urodstvo/messenger-server/libs/config"
	"github.com/urodstvo/messenger-server/libs/grpc/clients"
	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	"github.com/urodstvo/messenger-server/libs/logger"
	"go.uber.org/fx"
)

const service = "Chat-Service"

var App = fx.Module(
	service,
	app.CreateBaseApp(service),
	fx.Provide(
		func(c config.Config) proto.AuthServiceClient {
			return clients.NewAuthAppClient(c.EnvMode == "production")
		},
		websocket.New,
	),
	fx.Invoke(
		implementation.New,
		func() {
			go http.ListenAndServe(":8001", nil)
		},
		func(l logger.Logger) {
			l.Info(service + " started")
		},
	),
)
