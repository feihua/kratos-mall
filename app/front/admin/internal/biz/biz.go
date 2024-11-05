package biz

import (
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/oms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/pay"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/pms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sms"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/sys"
	"github.com/feihua/kratos-mall/app/front/admin/internal/biz/ums"
	"github.com/google/wire"
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
