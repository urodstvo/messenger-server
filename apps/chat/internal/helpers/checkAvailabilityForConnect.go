package helpers

import (
	"errors"

	"github.com/olahol/melody"
	models "github.com/urodstvo/messenger-server/libs/models/golang"
	"gorm.io/gorm"
)

func CheckAvailabilityForConnect(db *gorm.DB, session *melody.Session) error {
	chatId, ok := session.Get("chatId")
	if chatId == "" || !ok {
		return errors.New("no chat id")
	}

	userId, ok := session.Get("userId")
	if userId == "" || !ok {
		return errors.New("no user id")
	}

	if err := db.Where(`"user_id" = ? AND "chat_id" = ?`, userId, chatId).First(&models.Participant{}).Error; err != nil {
		return err
	}

	return nil
}
