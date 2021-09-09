package oms

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewCartItemRepo,
	NewCompanyAddressRepo,
	NewOperateHistoryRepo,
	NewOrderRepo,
	NewReturnApplyRepo,
	NewReturnReasonRepo,
	NewSettingRepo,
)
