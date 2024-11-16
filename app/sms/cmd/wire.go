//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/feihua/kratos-mall/app/sms/internal/biz"
	"github.com/feihua/kratos-mall/app/sms/internal/conf"
	"github.com/feihua/kratos-mall/app/sms/internal/data"
	"github.com/feihua/kratos-mall/app/sms/internal/server"
	"github.com/feihua/kratos-mall/app/sms/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}