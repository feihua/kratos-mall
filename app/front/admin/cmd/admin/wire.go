// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-mall/app/front/admin/internal/biz"
	"kratos-mall/app/front/admin/internal/conf"
	"kratos-mall/app/front/admin/internal/data"
	"kratos-mall/app/front/admin/internal/data/oms"
	"kratos-mall/app/front/admin/internal/data/pay"
	"kratos-mall/app/front/admin/internal/data/pms"
	"kratos-mall/app/front/admin/internal/data/sms"
	"kratos-mall/app/front/admin/internal/data/sys"
	"kratos-mall/app/front/admin/internal/data/ums"
	"kratos-mall/app/front/admin/internal/server"
	"kratos-mall/app/front/admin/internal/service"
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
