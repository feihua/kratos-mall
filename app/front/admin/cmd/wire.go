//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz"
	"github.com/feihua/kratos-mall/app/front/admin/internal/conf"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data/oms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data/pay"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data/pms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data/sms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data/sys"
	"github.com/feihua/kratos-mall/app/front/admin/internal/data/ums"
	"github.com/feihua/kratos-mall/app/front/admin/internal/server"
	"github.com/feihua/kratos-mall/app/front/admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		oms.ProviderSet,
		pms.ProviderSet,
		pay.ProviderSet,
		sms.ProviderSet,
		sys.ProviderSet,
		ums.ProviderSet,
		newApp,
	),
	)
}
