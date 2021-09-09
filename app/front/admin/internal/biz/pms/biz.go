package pms

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewBrandUseCase,
	NewCategoryUseCase,
	NewCommentUseCase,
	NewCommentReplayUseCase,
	NewFeightTemplateUseCase,
	NewMemberPriceUseCase,
	NewOperateHistoryUseCase,
	NewProductUseCase,
	NewSkuStockUseCase,
	NewVertifyRecordUseCase,
)
