package cmd

import (
	"go.uber.org/fx"
	"hello-go/infrafx"
)

func Execute() {
	app := fx.New(infrafx.Module)

	if err := app.Err(); err == nil {
		app.Run()
	} else {
		panic(err)
	}
}
