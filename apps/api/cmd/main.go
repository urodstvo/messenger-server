package main

import (
	"github.com/urodstvo/messenger-server/apps/api/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(app.App).Run()
}
