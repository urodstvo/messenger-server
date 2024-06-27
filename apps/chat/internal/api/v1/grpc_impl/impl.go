package implementation

import (
	"context"
	"fmt"
	"net"

	"github.com/urodstvo/messenger-server/libs/grpc/constants"
	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	"github.com/urodstvo/messenger-server/libs/logger"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Implementation struct {
	proto.UnimplementedChatServiceServer

	db     *gorm.DB
	logger logger.Logger
}

type Opts struct {
	fx.In
	lc fx.Lifecycle

	db     *gorm.DB
	Logger logger.Logger
}

func New(opts Opts) (proto.ChatServiceServer, error) {
	impl := &Implementation{
		db:     opts.db,
		logger: opts.Logger,
	}

	grpcServer := grpc.NewServer()

	opts.lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", constants.CHAT_SERVER_PORT))
				if err != nil {
					return err
				}
				proto.RegisterChatServiceServer(grpcServer, impl)

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
