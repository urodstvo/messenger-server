package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/urodstvo/messenger-server/libs/config"
	"github.com/urodstvo/messenger-server/libs/grpc/constants"
	"go.uber.org/fx"
)

type Email struct {
	from    string
	address string
	auth    smtp.Auth
}

type Opts struct {
	fx.In

	config config.Config
}

func New(opts Opts) Email {
	return Email{
		from:    opts.config.EmailFrom,
		address: fmt.Sprintf("%s:%s", opts.config.EmailHost, opts.config.EmailPort),
		auth:    smtp.PlainAuth(string(constants.AUTH_SERVER_PORT), opts.config.EmailFrom, opts.config.EmailPassword, opts.config.EmailHost),
	}
}

func (e *Email) SendOTP(to []string, otp string) error {
	t, _ := template.ParseFiles("otp.html")
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", "One-Time-Password", mimeHeaders)))

	t.Execute(&body, struct {
		OTP string
	}{
		OTP: otp,
	})

	if err := smtp.SendMail(e.address, e.auth, e.from, to, body.Bytes()); err != nil {
		return err
	}

	return nil
}
