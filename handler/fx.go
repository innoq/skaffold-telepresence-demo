package handler

import "go.uber.org/fx"

var Module = fx.Options(fx.Invoke(newHandler))
