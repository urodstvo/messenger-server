package helpers

import (
	"errors"

	"github.com/olahol/melody"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CheckChatId(db *gorm.DB, session *melody.Session) error {
	chatId := session.Request.URL.Query().Get("chatId")
	if chatId == "" {
		session.Close()
		return errors.New("no chatId")
	}

	dbChat := &models.Chat{}
	if err := db.Where(`"id" = ?`, chatId).First(dbChat).Error; err != nil {
		zap.S().Errorf(chatId, err)
		session.Close()
		return errors.New("no chat with this id")
	}

	session.Set("chatId", chatId)

	return nil
}
