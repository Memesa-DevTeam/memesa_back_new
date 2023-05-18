package app

import (
	"go.uber.org/fx"
	"memesa_go_backend/pkg/config"
	"memesa_go_backend/pkg/server"
)

func MemesaServices() fx.Option {
	return fx.Options(config.Provide(), server.Provide(), Provide())
}
