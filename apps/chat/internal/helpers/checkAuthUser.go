package helpers

import (
	"errors"

	"github.com/olahol/melody"

	proto "github.com/urodstvo/messenger-server/libs/grpc/proto/__generated__"
)

//TODO

func CheckAuthentificateUser(session *melody.Session, authGRPC proto.AuthServiceClient) error {
	apiKey := session.Request.Header.Get("Authorization")
	if apiKey == "" {
		session.Close()
		return errors.New("no api key")
	}

	// check api key
	userId, err := authGRPC.CheckApiKey(apiKey)
	if err != nil {
		session.Close()
		return err
	}

	session.Set("userId", userId)

	return nil
}
