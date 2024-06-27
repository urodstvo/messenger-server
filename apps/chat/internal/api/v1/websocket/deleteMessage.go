package websocket

import (
	"errors"
	"log/slog"

	"github.com/olahol/melody"
	"github.com/urodstvo/messenger-server/apps/chat/internal/types"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
)

func (c *Chat) DeleteMessage(session *melody.Session, data any) error {
	userId, ok := session.Get("userId")
	if userId == "" || !ok {
		return errors.New("no user id")
	}

	chatId, ok := session.Get("chatId")
	if chatId == "" || !ok {
		return errors.New("no chat id")
	}

	message, ok := data.(types.DeleteMessage)
	if !ok {
		c.logger.Error("cannot process message")
		return errors.New("cannot process message")
	}

	if err := c.db.Delete(&models.Message{}, message.MessageID).Error; err != nil {
		c.logger.Error("cannot delete message", slog.Any("err", err))
		return err
	}

	if err := c.SendEvent(chatId.(string), "deleted_message", message); err != nil {
		c.logger.Error("cannot send message", slog.Any("err", err))
		return err
	}

	return nil
}
