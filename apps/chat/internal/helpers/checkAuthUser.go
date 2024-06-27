package helpers

import (
	"errors"

	"github.com/olahol/melody"

	"github.com/urodstvo/messenger-server/libs/jwt"
)

func CheckAuthentificateUser(session *melody.Session, secret string) error {
	apiKey := session.Request.Header.Get("Authorization")
	if apiKey == "" {
		session.Close()
		return errors.New("no api key")
	}

	// check api key
	userId, err := jwt.Parse(apiKey, secret)
	if err != nil {
		session.Close()
		return err
	}

	session.Set("userId", userId)

	return nil
}
