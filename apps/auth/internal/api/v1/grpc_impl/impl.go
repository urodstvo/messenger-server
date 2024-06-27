package implementation

import (
	"context"
	"fmt"
	"net"

	"github.com/urodstvo/messenger-server/apps/auth/internal/email"
	"github.com/urodstvo/messenger-server/libs/config"
	"github.com/urodstvo/messenger-server/libs/grpc/constants"
	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	"github.com/urodstvo/messenger-server/libs/logger"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Implementation struct {
	proto.UnimplementedAuthServiceServer

	cfg    config.Config
	db     *gorm.DB
	logger logger.Logger

	email email.Email
}

type Opts struct {
	fx.In
	lc fx.Lifecycle

	cfg    config.Config
	db     *gorm.DB
	Logger logger.Logger

	Email email.Email
}

func New(opts Opts) (proto.AuthServiceServer, error) {
	impl := &Implementation{
		cfg:    opts.cfg,
		db:     opts.db,
		logger: opts.Logger,
		email:  opts.Email,
	}

	grpcServer := grpc.NewServer()

	opts.lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", constants.AUTH_SERVER_PORT))
				if err != nil {
					return err
				}
				proto.RegisterAuthServiceServer(grpcServer, impl)

				go grpcServer.Serve(lis)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				grpcServer.GracefulStop()
				return nil
			},
		},
	)

	return impl, nil
}
