package app

import (
	"github.com/urodstvo/messenger-server/libs/config"
	"github.com/urodstvo/messenger-server/libs/logger"
	"go.uber.org/fx"
)

func CreateBaseApp(appName string) fx.Option {
	return fx.Options(
		fx.Provide(
			config.NewFx,
			logger.NewFx(
				logger.LoggerOptions{
					Service: appName,
				},
			),
			newGorm,
			newRedis,
		),
	)
}
