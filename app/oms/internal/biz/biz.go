package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewGreeterUsecase,
	NewCartItemUseCase,
	NewCompanyAddressUseCase,
	NewOperateHistoryUseCase,
	NewOrderUseCase,
	NewReturnApplyUseCase,
	NewReturnReasonUseCase,
	NewSettingUseCase,
)
