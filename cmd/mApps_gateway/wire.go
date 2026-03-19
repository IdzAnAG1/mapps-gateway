//go:build wireinject

package main

import (
	"mapps_gateway/internal/conf"
	"mapps_gateway/internal/data"
	"mapps_gateway/internal/server"
	"mapps_gateway/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp инициализирует приложение через Dependency Injection.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
