package app

import (
	implementation "github.com/urodstvo/messenger-server/apps/auth/internal/api/v1/grpc_impl"
	"github.com/urodstvo/messenger-server/apps/auth/internal/email"
	"github.com/urodstvo/messenger-server/libs/app"
	"github.com/urodstvo/messenger-server/libs/logger"
	"go.uber.org/fx"
)

const service = "Auth-Service"

var App = fx.Module(
	service,
	app.CreateBaseApp(service),
	fx.Provide(
		email.New,
	),
	fx.Invoke(
		implementation.New,
		func(l logger.Logger) {
			l.Info(service + " started")
		},
	),
)
