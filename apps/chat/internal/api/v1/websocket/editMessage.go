package websocket

import (
	"errors"
	"log/slog"

	"github.com/olahol/melody"
	"github.com/urodstvo/messenger-server/apps/chat/internal/types"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"gorm.io/gorm/clause"
)

func (c *Chat) EditMessage(session *melody.Session, data any) error {
	userId, ok := session.Get("userId")
	if userId == "" || !ok {
		return errors.New("no user id")
	}

	chatId, ok := session.Get("chatId")
	if chatId == "" || !ok {
		return errors.New("no chat id")
	}

	message, ok := data.(types.EditMessage)
	if !ok {
		c.logger.Error("cannot process message")
		return errors.New("cannot process message")
	}

	var dbMessage models.Message
	if err := c.db.Model(&dbMessage).Clauses(clause.Returning{}).Where("id = ?", message.MessageID).Update("text", message.Text).Error; err != nil {
		c.logger.Error("cannot update message", slog.Any("err", err))
		return err
	}

	if err := c.SendEvent(chatId.(string), "edited_message", dbMessage); err != nil {
		c.logger.Error("cannot send message", slog.Any("err", err))
		return err
	}

	return nil
}
