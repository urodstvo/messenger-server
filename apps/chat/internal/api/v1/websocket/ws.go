package websocket

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/olahol/melody"
	"github.com/urodstvo/messenger-server/apps/chat/internal/helpers"
	"github.com/urodstvo/messenger-server/apps/chat/internal/types"
	"github.com/urodstvo/messenger-server/libs/config"
	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
	"github.com/urodstvo/messenger-server/libs/logger"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Chat struct {
	manager  *melody.Melody
	config   config.Config
	db       *gorm.DB
	logger   logger.Logger
	authGrpc proto.AuthServiceClient
}

type Opts struct {
	fx.In

	Config   config.Config
	DB       *gorm.DB
	Logger   logger.Logger
	AuthGrpc proto.AuthServiceClient
}

func New(opts Opts) *Chat {
	m := melody.New()
	m.Config.MaxMessageSize = 1024 * 1024 * 10

	chat := &Chat{
		manager:  m,
		db:       opts.DB,
		logger:   opts.Logger,
		config:   opts.Config,
		authGrpc: opts.AuthGrpc,
	}

	chat.manager.HandleConnect(
		func(session *melody.Session) {
			if err := helpers.CheckAuthentificateUser(session, opts.Config.JWTSecret); err != nil {
				opts.Logger.Error("cannot check user by api key", slog.Any("err", err))
				session.Close()
				return
			}

			if err := helpers.CheckChatId(opts.DB, session); err != nil {
				opts.Logger.Error("cannot check chat id", slog.Any("err", err))
				session.Close()
				return
			}

			if err := helpers.CheckAvailabilityForConnect(opts.DB, session); err != nil {
				opts.Logger.Error("cannot connect to chat", slog.Any("err", err))
				session.Close()
				return
			}

			session.Write([]byte(`{"eventName":"connected to chat service"}`))
		},
	)

	chat.manager.HandleMessage(
		func(session *melody.Session, msg []byte) {
			chat.handleMessage(session, msg)
		},
	)

	chat.manager.HandleDisconnect(
		func(session *melody.Session) {},
	)

	http.HandleFunc("/chat", chat.HandleRequest)

	return chat
}

func (c *Chat) HandleRequest(w http.ResponseWriter, r *http.Request) {
	_ = c.manager.HandleRequest(w, r)
}

func (c *Chat) SendEvent(channelId, eventName string, data any) error {
	message := &types.WebSocketMessage{
		EventName: eventName,
		Data:      data,
		CreatedAt: time.Now().UTC().String(),
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		c.logger.Error("cannot process message", slog.Any("err", err))
		return err
	}

	err = c.manager.BroadcastFilter(
		bytes, func(session *melody.Session) bool {
			socketChatId, ok := session.Get("chatId")
			return ok && socketChatId.(string) == channelId
		},
	)

	if err != nil {
		c.logger.Error("cannot broadcast message", slog.Any("err", err))
		return err
	}

	return nil
}

func (c *Chat) handleMessage(session *melody.Session, msg []byte) {

	data := &types.WebSocketMessage{
		CreatedAt: time.Now().UTC().String(),
	}
	err := json.Unmarshal(msg, data)
	if err != nil {
		c.logger.Error(err.Error())
		return
	}

	switch data.EventName {
	case "send_message":
		if err = c.SendMessage(session, data.Data); err != nil {
			c.logger.Error("cannot send message", slog.Any("err", err))
			return
		}
	case "edit_message":
		if err = c.EditMessage(session, data.Data); err != nil {
			c.logger.Error("cannot edit message", slog.Any("err", err))
			return
		}
	case "read_message":
		if err = c.ReadMessage(session, data.Data); err != nil {
			c.logger.Error("cannot read message", slog.Any("err", err))
			return
		}
	case "delete_message":
		if err = c.DeleteMessage(session, data.Data); err != nil {
			c.logger.Error("cannot delete message", slog.Any("err", err))
			return
		}
	default:
		return
	}
}
