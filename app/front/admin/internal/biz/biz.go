package biz

import (
	"github.com/google/wire"
	"kratos-mall/app/front/admin/internal/biz/oms"
	"kratos-mall/app/front/admin/internal/biz/pay"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/biz/sms"
	"kratos-mall/app/front/admin/internal/biz/sys"
	"kratos-mall/app/front/admin/internal/biz/ums"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewGreeterUsecase,
	oms.ProviderSet,
	pms.ProviderSet,
	pay.ProviderSet,
	sms.ProviderSet,
	sys.ProviderSet,
	ums.ProviderSet,
)
