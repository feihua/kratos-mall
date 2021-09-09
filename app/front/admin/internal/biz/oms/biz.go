package oms

import "github.com/google/wire"

// ProviderSet is oms providers.
var ProviderSet = wire.NewSet(
	NewCartItemUseCase,
	NewCompanyAddressUseCase,
	NewOperateHistoryUseCase,
	NewOrderUseCase,
	NewReturnApplyUseCase,
	NewReturnReasonUseCase,
	NewSettingUseCase,
)
